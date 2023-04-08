package postgres

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/pkg/apperror"
	"github.com/jmoiron/sqlx"
	"time"
)

type MessageRepo struct {
	DB *sqlx.DB
}

func NewMessageRepo(DB *sqlx.DB) repository.Message {
	return &MessageRepo{DB: DB}
}

func (m *MessageRepo) SaveMsg(message *dto.NewMessage) (*entities.Message, error) {
	msg := &entities.Message{}
	err := m.DB.QueryRowx(CreateMessage,
		message.UserId, message.ChatId,
		message.ContentType,
		time.Now().String(),
		message.Text, message.ReplyTo).StructScan(msg)

	if err != nil {
		return nil, apperror.BadRequest
	}
	msg.Link = message.Link
	return msg, nil
}

func (m *MessageRepo) GetMembersByChatId(chatId uint) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := m.DB.Queryx(GetMembersByChatId, chatId)
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
