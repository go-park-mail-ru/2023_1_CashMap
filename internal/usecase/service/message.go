package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"errors"
	"github.com/asaskevich/govalidator"
)

type MessageService struct {
	repository.MessageRepository
}

func NewMessageService(repository repository.MessageRepository) usecase.MessageUsecase {
	return &MessageService{repository}
}

func (service MessageService) GetMessagesByChatID(senderEmail string, dto *dto.GetMessagesDTO) ([]*entities.Message, error) {
	isValid, err := govalidator.ValidateStruct(dto)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("invalid struct")
	}

	return service.MessageRepository.SelectMessagesByChatID(senderEmail, dto)
}

func (service MessageService) GetChatsList(senderEmail string, dto *dto.GetChatsDTO) ([]*entities.Chat, error) {
	isValid, err := govalidator.ValidateStruct(dto)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("invalid struct")
	}

	if dto.Offset == nil {
		dto.Offset = new(uint)
		*dto.Offset = 0
	}

	return service.MessageRepository.SelectChats(senderEmail, dto)
}

func (service MessageService) CreateChat(senderEmail string, dto *dto.CreateChatDTO) (*entities.Chat, error) {
	isValid, err := govalidator.ValidateStruct(dto)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("invalid struct")
	}

	chatID, err := service.MessageRepository.CreateChat(senderEmail, dto)
	if err != nil {
		return nil, err
	}

	//TODO: доделать репу на запрос учатсников чата по id
	return &entities.Chat{
		ChatID: chatID,
	}, err
}

func (service MessageService) HasDialog(senderEmail string, dto *dto.HasDialogDTO) (bool, error) {
	isValid, err := govalidator.ValidateStruct(dto)
	if err != nil {
		return false, err
	}
	if !isValid {
		return false, errors.New("invalid struct")
	}

	return service.MessageRepository.HasDialog(senderEmail, dto)
}
