package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type MessageUsecase interface {
	GetMessagesByChatID(senderEmail string, dto *dto.GetMessagesDTO) ([]*entities.Message, error)
	GetChatsList(senderEmail string, dto *dto.GetChatsDTO) ([]*entities.Chat, error)
	CreateChat(senderEmail string, dto *dto.CreateChatDTO) (*entities.Chat, error)
	HasDialog(senderEmail string, dto *dto.HasDialogDTO) (bool, error)
	Send(email string, message *dto.NewMessage) (*entities.Message, error)
	GetMembersByChatId(chatId uint) ([]*entities.User, error)
}
