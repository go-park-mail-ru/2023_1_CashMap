package response

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type SendRequest struct {
	Body *dto.NewMessageDTO `json:"body"`
}

type NewChatRequest struct {
	Body *dto.CreateChatDTO `json:"body"`
}

type NewChatResponse struct {
	Body NewChatBody `json:"body"`
}

type NewChatBody struct {
	Chat *entities.Chat `json:"chat"`
}

type GetChatsResponse struct {
	Body GetChatsBody `json:"body"`
}

type GetChatsBody struct {
	Chats []*entities.Chat `json:"chats"`
}

type GetMessagesByChatIDResponse struct {
	Body GetMessagesByChatIDBody `json:"body"`
}

type GetMessagesByChatIDBody struct {
	Messages []*entities.Message `json:"messages"`
	HasNext  bool                `json:"has_next"`
}

type HasDialogResponse struct {
	Body HasDialogBody `json:"body"`
}

type HasDialogBody struct {
	ChatId    int  `json:"chat_id"`
	HasDialog bool `json:"has_dialog"`
}
