package doc

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

// ChatCreateRequest entity info
//
//	@Description	All post information
type ChatCreateRequest struct {
	Body dto.CreateChatDTO `json:"body"`
}

// ChatCreateResponse entity info
//
//	@Description	All post information
type ChatCreateResponse struct {
	Body Chat `json:"body"`
}

// Chat entity info
//
//	@Description	All post information
type Chat struct {
	Chat entities.Chat `json:"chat"`
}

// SendMessageRequest entity info
//
//	@Description	All post information
type SendMessageRequest struct {
	Body dto.NewMessageDTO `json:"body"`
}

// SendMessageResponse entity info
//
//	@Description	All post information
type SendMessageResponse struct {
	Body dto.NewMessageDTO `json:"body"`
}

// MessagesListResponse entity info
//
//	@Description	All post information
type MessagesListResponse struct {
	Body Messages `json:"body"`
}

// Messages entity info
//
//	@Description	All post information
type Messages struct {
	Messages []entities.Message `json:"messages"`
}

// HasDialogResponse entity info
//
//	@Description	All post information
type HasDialogResponse struct {
	Body Exists `json:"body"`
}

type Exists struct {
	HasDialog bool `json:"has_dialog"`
}
