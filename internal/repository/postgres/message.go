package postgres

import (
	"context"
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	utildb "depeche/internal/repository/utils"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"errors"
	"github.com/jmoiron/sqlx"
)

type MessageStorage struct {
	db *sqlx.DB
}

func NewMessageRepository(DB *sqlx.DB) repository.MessageRepository {
	return &MessageStorage{DB}
}

func (storage *MessageStorage) SaveMsg(message *dto.NewMessageDTO) (*entities.Message, error) {
	msg := &entities.Message{}
	err := storage.db.QueryRowx(CreateMessage,
		message.UserId, message.ChatId,
		message.ContentType,
		message.Text,
		utils.CurrentTimeString(),
		message.ReplyTo).Scan(&msg.Id)

	if err != nil {
		return nil, apperror.NewServerError(apperror.BadRequest, err)
	}

	err = storage.db.QueryRowx(MessageById, msg.Id).StructScan(msg)
	if err != nil {

		return nil, apperror.NewServerError(apperror.BadRequest, err)
	}
	return msg, nil
}

func (storage *MessageStorage) GetMembersByChatId(chatId uint) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := storage.db.Queryx(GetMembersByChatId, chatId)
	defer utildb.CloseRows(rows)
	if err != nil {
		return nil, apperror.NewServerError(apperror.BadRequest, err)
	}
	for rows.Next() {
		user := &entities.User{}
		err := rows.StructScan(user)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
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

	rows, err := storage.db.Queryx(
		MessageByChatIdQuery,
		dto.ChatID,
		dto.LastMessageDate,
		dto.BatchSize)
	defer utildb.CloseRows(rows)
	if err != nil {
		return nil, err
	}

	messages, err := getSliceFromRows[entities.Message](rows, *dto.BatchSize)
	if err != nil {
		return nil, err
	}

	for _, message := range messages {
		info, err := storage.GetUserInfoByMessageId(*message.Id)
		if err != nil {
			return nil, err
		}

		message.SenderInfo = info
	}

	return messages, nil
}

func (storage *MessageStorage) SelectChats(senderEmail string, dto *dto.GetChatsDTO) ([]*entities.Chat, error) {
	rows, err := storage.db.Queryx(
		ChatsQuery,
		senderEmail,
		dto.BatchSize,
		dto.Offset)
	defer utildb.CloseRows(rows)
	if err != nil {
		return nil, err
	}

	chats, err := getSliceFromRows[entities.Chat](rows, *dto.BatchSize)
	if err != nil {
		return nil, err
	}

	for _, chat := range chats {
		usersInfo, err := storage.GetUsersInfoByChatID(chat.ChatID)
		if err != nil {
			return nil, err
		}
		chat.Users = usersInfo
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
		if exists != nil {
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

func (storage *MessageStorage) HasDialog(senderEmail string, dto *dto.HasDialogDTO) (*int, error) {
	chatId, err := storage.isPersonalChatExists(senderEmail, *dto.UserLink)
	if err != nil {
		return nil, nil
	}

	if chatId != nil {
		return chatId, nil
	}

	return nil, nil
}

func (storage *MessageStorage) GetUsersInfoByChatID(chatID uint) ([]*entities.UserInfo, error) {
	var chat []*entities.UserInfo
	err := storage.db.Select(&chat, UserInfoByChatIdQuery, chatID)
	if err != nil {
		return nil, err
	}

	return chat, err
}

func (storage *MessageStorage) GetUserInfoByMessageId(messageID uint) (*entities.UserInfo, error) {
	senderInfo := &entities.UserInfo{}
	err := storage.db.Get(senderInfo, UserInfoByMessageIdQuery,
		messageID)
	if err != nil {
		return nil, err
	}

	return senderInfo, nil
}

func (storage *MessageStorage) isPersonalChatExists(email string, userLink string) (*int, error) {
	var chatId int
	err := storage.db.Get(
		&chatId,
		IsChatExistsQuery,
		userLink,
		email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &chatId, nil
}

func (storage *MessageStorage) isChatMember(email string, chatID uint) (bool, error) {
	var hasAccess bool
	err := storage.db.Get(
		&hasAccess,
		IsChatMemberQuery,
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
