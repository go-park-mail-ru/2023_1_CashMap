package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type MessageUsecase interface {
	GetMessagesByChatID(senderEmail string, dto *dto.GetMessagesDTO) ([]*entities.Message, bool, error)
	GetChatsList(senderEmail string, dto *dto.GetChatsDTO) ([]*entities.Chat, error)
	CreateChat(senderEmail string, dto *dto.CreateChatDTO) (*entities.Chat, error)
	HasDialog(senderEmail string, dto *dto.HasDialogDTO) (*int, error)
	Send(email string, message *dto.NewMessageDTO) (*entities.Message, error)
	GetMembersByChatId(chatId uint) ([]*entities.User, error)
	GetUnreadChatsCount(email string) (int, error)
	SetLastRead(email string, chatID int, time string) error
}
