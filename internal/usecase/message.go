package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type Message interface {
	Send(message *dto.NewMessage) (*entities.Message, error)
	GetMembersByChatId(chatId uint) ([]*entities.User, error)
}
