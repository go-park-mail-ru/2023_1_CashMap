package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type Message interface {
	SaveMsg(message *dto.NewMessage) (*entities.Message, error)
	GetMembersByChatId(chatId uint) ([]*entities.User, error)
}
