package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type MessageRepository interface {
	SelectMessagesByChatID(senderEmail string, dto *dto.GetMessagesDTO) ([]*entities.Message, error)
	SelectChats(senderEmail string, dto *dto.GetChatsDTO) ([]*entities.Chat, error)
	CreateChat(senderEmail string, dto *dto.CreateChatDTO) (uint, error)
	HasDialog(senderEmail string, dto *dto.HasDialogDTO) (bool, error)
	SaveMsg(message *dto.NewMessage) (*entities.Message, error)
	GetMembersByChatId(chatId uint) ([]*entities.User, error)
	GetUsersInfoByChatID(chatID uint) ([]*entities.UserInfo, error)
	GetUserInfoByMessageId(messageID uint) (*entities.UserInfo, error)
}
