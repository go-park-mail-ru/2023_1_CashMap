package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"errors"
	"github.com/asaskevich/govalidator"
)

type MessageService struct {
	repository.MessageRepository
	repository.UserRepository
}

func NewMessageService(repo repository.MessageRepository, userRepo repository.UserRepository) usecase.MessageUsecase {
	return &MessageService{repo, userRepo}
}

func (service *MessageService) Send(email string, message *dto.NewMessageDTO) (*entities.Message, error) {
	user, err := service.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	message.UserId = user.ID

	message = utils.Escaping(message)
	msg, err := service.MessageRepository.SaveMsg(message)
	if err != nil {
		return nil, err
	}
	info, err := service.MessageRepository.GetUserInfoByMessageId(*msg.Id)
	if err != nil {
		return nil, err
	}
	msg.SenderInfo = info

	if message.Attachments != nil {
		if len(message.Attachments) > 10 {
			return nil, apperror.NewServerError(apperror.TooMuchAttachments, nil)
		}
		err := service.MessageRepository.AddMessageAttachments(*msg.Id, msg.Attachments)
		if err != nil {
			return msg, err
		}
		msg.Attachments = message.Attachments
	}

	return msg, nil
}

func (service *MessageService) GetMembersByChatId(chatId uint) ([]*entities.User, error) {
	return service.MessageRepository.GetMembersByChatId(chatId)
}

func (service *MessageService) GetMessagesByChatID(senderEmail string, dto *dto.GetMessagesDTO) ([]*entities.Message, error) {
	isValid, err := govalidator.ValidateStruct(dto)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("invalid struct")
	}

	if dto.LastMessageDate == nil {
		dto.LastMessageDate = new(string)
		*dto.LastMessageDate = "0"
	}

	messages, err := service.MessageRepository.SelectMessagesByChatID(senderEmail, dto)
	if err != nil {
		return nil, err
	}
	for _, message := range messages {
		attachments, err := service.MessageRepository.GetMessageAttachments(*message.Id)
		if err != nil {
			return nil, err
		}
		message.Attachments = attachments
	}
	return messages, nil
}

func (service *MessageService) GetChatsList(senderEmail string, dto *dto.GetChatsDTO) ([]*entities.Chat, error) {
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

func (service *MessageService) CreateChat(senderEmail string, dto *dto.CreateChatDTO) (*entities.Chat, error) {
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
	usersInfo, err := service.MessageRepository.GetUsersInfoByChatID(chatID)
	if err != nil {
		return nil, err
	}
	return &entities.Chat{
		ChatID: chatID,
		Users:  usersInfo,
	}, err
}

func (service *MessageService) HasDialog(senderEmail string, dto *dto.HasDialogDTO) (*int, error) {
	isValid, err := govalidator.ValidateStruct(dto)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("invalid struct")
	}
	return service.MessageRepository.HasDialog(senderEmail, dto)
}
