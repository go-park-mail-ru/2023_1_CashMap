package postgres

import (
	"context"
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MessageStorage struct {
	db *sqlx.DB
}

func NewMessageRepository(DB *sqlx.DB) repository.MessageRepository {
	return &MessageStorage{DB}
}

func (storage *MessageStorage) SaveMsg(message *dto.NewMessage) (*entities.Message, error) {
	msg := &entities.Message{}
	err := storage.db.QueryRowx(CreateMessage,
		message.UserId, message.ChatId,
		message.ContentType,
		message.Text,
		utils.CurrentTimeString(),
		message.ReplyTo).Scan(&msg.Id)

	if err != nil {
		fmt.Println(err)
		return nil, apperror.BadRequest
	}

	err = storage.db.QueryRowx(MessageById, msg.Id).StructScan(msg)
	if err != nil {
		fmt.Println(err)
		return nil, apperror.BadRequest
	}
	msg.Link = message.Link
	return msg, nil
}

func (storage *MessageStorage) GetMembersByChatId(chatId uint) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := storage.db.Queryx(GetMembersByChatId, chatId)
	if err != nil {
		return nil, apperror.BadRequest
	}
	for rows.Next() {
		user := &entities.User{}
		err := rows.StructScan(user)
		if err != nil {
			return nil, apperror.InternalServerError
		}
		users = append(users, user)
	}
	return users, nil
}

func (storage *MessageStorage) SelectMessagesByChatID(senderEmail string, dto *dto.GetMessagesDTO) ([]*entities.Message, error) {
	hasAccess, err := storage.isChatMember(senderEmail, *dto.ChatID)
	if err != nil {
		return nil, err
	}
	if !hasAccess {
		return nil, errors.New("access to messages isn't allowed")
	}

	rows, err := storage.db.Queryx("SELECT msg.id, msg.chat_id, text_content, author.link as link, msg.creation_date, msg.change_date, msg.reply_to, msg.is_deleted "+
		"FROM Message AS msg JOIN UserProfile AS author ON msg.user_id = author.id "+
		"WHERE msg.chat_id = (SELECT id FROM Chat WHERE id = $1) AND msg.creation_date > $2 AND msg.is_deleted = false ORDER BY msg.creation_date DESC LIMIT $3",
		dto.ChatID,
		dto.LastPostDate,
		dto.BatchSize)

	if err != nil {
		return nil, err
	}

	messages, err := getSliceFromRows[entities.Message](rows, *dto.BatchSize)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (storage *MessageStorage) SelectChats(senderEmail string, dto *dto.GetChatsDTO) ([]*entities.Chat, error) {
	rows, err := storage.db.Queryx("SELECT chat.id as chat_id FROM ChatMember as member"+
		" JOIN Chat ON chat_id = chat.id"+
		" WHERE member.user_id = (SELECT id FROM UserProfile WHERE email = $1 LIMIT $2 OFFSET $3)",
		senderEmail,
		dto.BatchSize,
		dto.Offset)
	if err != nil {
		return nil, err
	}

	chats, err := getSliceFromRows[entities.Chat](rows, *dto.BatchSize)
	if err != nil {
		return nil, err
	}

	for ind, chat := range chats {
		var userLinks []string
		err := storage.db.Select(&userLinks, "SELECT link FROM ChatMember JOIN UserProfile ON id = user_id WHERE chat_id = $1", chat.ChatID)
		if err != nil {
			return nil, err
		}

		chats[ind].UserLinks = userLinks
	}

	return chats, nil
}

func (storage *MessageStorage) CreateChat(senderEmail string, dto *dto.CreateChatDTO) (uint, error) {
	var chatID int
	switch len(dto.UserLinks) {
	case 0:
		return 0, errors.New("empty list of users")
	case 1:
		exists, err := storage.isPersonalChatExists(senderEmail, dto.UserLinks[0])
		if err != nil {
			return 0, err
		}
		if exists {
			return 0, errors.New("chat already exists")
		}

		tx, err := storage.db.Beginx()
		if err != nil {
			return 0, err
		}

		query := "INSERT INTO Chat (id, members_number) VALUES (DEFAULT, $1) RETURNING id"
		err = tx.GetContext(context.Background(), &chatID, query, len(dto.UserLinks)+1)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return 0, err
			}
			return 0, err
		}

		err = storage.addChatMembers(senderEmail, dto.UserLinks, uint(chatID), tx)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return 0, err
			}
			return 0, err
		}

		err = tx.Commit()
		if err != nil {
			return 0, err
		}

	default:
		return 0, errors.New("group chat are unsupported")
		// TODO: реализовать групповые чаты
	}

	return uint(chatID), nil
}

func (storage *MessageStorage) HasDialog(senderEmail string, dto *dto.HasDialogDTO) (bool, error) {
	exists, err := storage.isPersonalChatExists(senderEmail, *dto.UserLink)
	if err != nil {
		return false, nil
	}

	if exists {
		return true, nil
	}

	return false, nil
}

func (storage *MessageStorage) isPersonalChatExists(email string, userLink string) (bool, error) {
	var exists bool
	err := storage.db.Get(&exists, "WITH CommonChats AS (SELECT DISTINCT first.chat_id as chat_id FROM ChatMember first JOIN ChatMember second ON first.chat_id = second.chat_id "+
		"WHERE first.user_id = (SELECT id FROM UserProfile WHERE link = $1) AND second.user_id = (SELECT id FROM UserProfile WHERE email = $2)) "+
		"SELECT true as exists FROM Chat as chat JOIN CommonChats as common ON common.chat_id = chat.id WHERE chat.members_number = 2",
		userLink,
		email)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return true, err
	}

	return true, nil
}

func (storage *MessageStorage) isChatMember(email string, chatID uint) (bool, error) {
	var hasAccess bool
	err := storage.db.Get(&hasAccess, "SELECT true FROM ChatMember member JOIN Chat as chat on chat.id = member.chat_id "+
		"WHERE chat_id = $1 AND member.user_id = (SELECT id FROM UserProfile WHERE email = $2)",
		chatID,
		email)
	if err != nil {
		return false, err
	}

	if !hasAccess {
		return false, nil
	}

	return true, nil
}

func (storage *MessageStorage) addChatMembers(senderEmail string, userLinks []string, chatID uint, tx *sqlx.Tx) error {
	var senderLink string
	err := tx.Get(&senderLink, "SELECT link FROM UserProfile WHERE email = $1", senderEmail)
	if err != nil {
		return err
	}
	userLinks = append(userLinks, senderLink)
	for i := 0; i < len(userLinks); i++ {
		var userID uint
		err = tx.Get(&userID, "SELECT id FROM UserProfile WHERE link = $1", userLinks[i])
		if err != nil {
			return err
		}

		_, err := tx.NamedExec("INSERT INTO ChatMember (chat_id, user_id, role) VALUES (:chat_id, :user_id, DEFAULT)",
			map[string]interface{}{
				"chat_id": chatID,
				"user_id": userID,
			})
		if err != nil {
			return err
		}

	}
	if err != nil {
		return err
	}
	return nil
}
