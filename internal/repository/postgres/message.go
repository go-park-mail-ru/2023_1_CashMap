package postgres

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/jmoiron/sqlx"
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
		message.Text,
		utils.CurrentTimeString(),
		message.ReplyTo).Scan(&msg.Id)

	if err != nil {
		fmt.Println(err)
		return nil, apperror.BadRequest
	}

	err = m.DB.QueryRowx(MessageById, msg.Id).StructScan(msg)
	if err != nil {
		fmt.Println(err)
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
