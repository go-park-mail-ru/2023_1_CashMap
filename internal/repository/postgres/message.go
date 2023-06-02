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
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
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
		message.ContentType, message.StickerID,
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
	if *msg.ContentType == "sticker" {
		sticker, err := storage.GetStickerById(*msg.StickerId)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		msg.Sticker = sticker
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

	utils.Reverse(messages)

	for _, message := range messages {
		info, err := storage.GetUserInfoByMessageId(*message.Id)
		if err != nil {
			return nil, err
		}

		message.SenderInfo = info

		if *message.ContentType == "sticker" {
			sticker, err := storage.GetStickerById(*message.StickerId)
			if err != nil {
				continue
			}
			message.Sticker = sticker
		}
	}

	return messages, nil
}

func (storage *MessageStorage) GetStickerById(stickerID uint) (*entities.Sticker, error) {
	sticker := &entities.Sticker{}
	err := storage.db.QueryRowx(GetStickerByID, stickerID).StructScan(sticker)
	if err == sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.StickerNotFound, nil)
	}
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return sticker, nil
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
		exists, err := storage.isChatExists(senderEmail, dto.UserLinks[0])
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
	chatId, err := storage.isChatExists(senderEmail, *dto.UserLink)
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

func (storage *MessageStorage) isChatExists(email string, userLink string) (*int, error) {
	var chatId int

	var link string
	err := storage.db.Get(&link, "SELECT link FROM userprofile where email = $1", email)
	if err != nil {
		return nil, err
	}

	if link == userLink {
		err = storage.db.Get(
			&chatId,
			IsPersonalChatExistsQuery,
			userLink)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		return &chatId, nil
	}

	err = storage.db.Get(
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

func (storage *MessageStorage) AddMessageAttachments(messageID uint, attachments []string) error {

	query := `insert into attachment (url) values `
	for i := 1; i < len(attachments)+1; i++ {
		query += fmt.Sprintf("($%d), ", i)
	}
	query, _ = strings.CutSuffix(query, ", ")
	query += " returning id"

	msgAttachments := make([]interface{}, len(attachments))
	for i, att := range attachments {
		msgAttachments[i] = att
	}

	rows, err := storage.db.Queryx(query, msgAttachments...)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	var attachmentIds []uint

	for rows.Next() {
		var id uint
		err := rows.Scan(&id)
		if err != nil {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
		attachmentIds = append(attachmentIds, id)
	}

	msgAttQuery := `insert into messageattachment (doc_id, message_id) values `
	for i := 2; i < len(attachmentIds)+2; i++ {
		msgAttQuery += fmt.Sprintf("($%d, $1), ", i)
	}
	msgAttQuery, _ = strings.CutSuffix(msgAttQuery, ", ")
	params := make([]interface{}, len(attachmentIds)+1)
	params[0] = messageID
	for i, id := range attachmentIds {
		params[i+1] = id
	}
	err = storage.db.QueryRowx(msgAttQuery, params...).Scan()
	if err != nil && err != sql.ErrNoRows {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}
	return nil
}

func (storage *MessageStorage) GetMessageAttachments(messageID uint) ([]string, error) {
	var attachments []string
	rows, err := storage.db.Queryx(GetMsgAttachments, messageID)
	if err != nil && err != sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var attach string
		err := rows.Scan(&attach)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		attachments = append(attachments, attach)
	}

	return attachments, nil
}

func (storage *MessageStorage) CheckRead(email string, chatID uint) (bool, error) {
	var read bool
	err := storage.db.QueryRowx(CheckChatRead, chatID, email).Scan(read)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return read, nil
}

func (storage *MessageStorage) GetUnreadChatsCount(email string) (int, error) {
	var count int
	err := storage.db.QueryRowx(GetUnreadChatCount, email).Scan(&count)
	if err != nil {
		return 0, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return count, nil
}

func (storage *MessageStorage) SetLastRead(email string, chatID int, time string) error {
	_, err := storage.db.Exec(SetLastRead, email, chatID, time)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}
	return nil
}

func (storage *MessageStorage) addChatMembers(senderEmail string, userLinks []string, chatID uint, tx *sqlx.Tx) error {
	var senderLink string
	err := tx.Get(&senderLink, "SELECT link FROM UserProfile WHERE email = $1", senderEmail)
	if err != nil {
		return err
	}

	if senderLink != userLinks[0] {
		userLinks = append(userLinks, senderLink)
	}
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
