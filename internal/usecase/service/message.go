package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
)

type Message struct {
	repo     repository.Message
	userRepo repository.UserRepository
}

func NewMessageService(repo repository.Message) usecase.Message {
	return &Message{
		repo: repo,
	}
}

func (m *Message) Send(message *dto.NewMessage) (*entities.Message, error) {
	user, err := m.userRepo.GetUserByLink(message.Link)
	if err != nil {
		return nil, err
	}
	message.UserId = user.ID
	return m.repo.SaveMsg(message)
}

func (m *Message) GetMembersByChatId(chatId uint) ([]*entities.User, error) {
	return m.repo.GetMembersByChatId(chatId)
}
