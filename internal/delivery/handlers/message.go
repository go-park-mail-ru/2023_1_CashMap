package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageHandler struct {
	usecase.MessageUsecase
}

func NewMessageHandler(service usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{service}
}

// Send godoc
//
//	@Summary		Send message
//	@Description	Add message to db and send it to listeners
//	@Tags			Message
//	@Param			request	body	doc.SendMessageResponse	false	"Message info"
//	@Success		200
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/im/send [post]
func (handler *MessageHandler) Send(ctx *gin.Context) {
	var request = struct {
		Data *dto.NewMessage `json:"body"`
	}{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	if request.Data == nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}
	email, ok := ctx.Get("email")
	if !ok {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	e, ok := email.(string)
	if !ok {
		_ = ctx.Error(apperror.BadRequest)
		return
	}
	msg, err := handler.MessageUsecase.Send(e, request.Data)
	if err != nil {
		_ = ctx.Error(err)
	}

	ctx.Set("message", msg)
}

// NewChat godoc
//
//	@Summary		Create chat
//	@Description	Create new chat with user
//	@Tags			Message
//	@Param			request	body		doc.ChatCreateRequest	false	"Chat info"
//	@Success		200		{object}	doc.ChatCreateResponse
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/im/chat/create [post]
func (handler *MessageHandler) NewChat(ctx *gin.Context) {
	var request = struct {
		dto.CreateChatDTO `json:"body"`
	}{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	chat, err := handler.MessageUsecase.CreateChat(email.(string), &request.CreateChatDTO)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"chat": chat,
		},
	})

}

// GetChats godoc
//
//	@Summary		Get chats
//	@Description	Get chats list
//	@Tags			Message
//	@Param			batch_size	query	uint	false	"Batch size"
//	@Param			offset		query	uint	false	"offset"
//	@Success		200			{array}	dto.GetChatsDTO
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/im/chats [get]
func (handler *MessageHandler) GetChats(ctx *gin.Context) {
	getChatsDTO := dto.GetChatsDTO{}
	err := ctx.ShouldBind(&getChatsDTO)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	chats, err := handler.MessageUsecase.GetChatsList(email.(string), &getChatsDTO)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"chats": chats,
		},
	})
}

// GetMessagesByChatID godoc
//
//	@Summary		Get messages
//	@Description	Get messages batch by chatID sorted by date
//	@Tags			Message
//	@Param			chat_id			query		uint	false	"Chat id"
//	@Param			batch_size		query		uint	false	"Batch size"
//	@Param			last_post_date	query		uint	false	"Last post date"
//	@Success		200				{object}	doc.MessagesListResponse
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/im/messages [get]
func (handler *MessageHandler) GetMessagesByChatID(ctx *gin.Context) {
	getMsgDTO := dto.GetMessagesDTO{}
	err := ctx.ShouldBind(&getMsgDTO)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	messages, err := handler.MessageUsecase.GetMessagesByChatID(email.(string), &getMsgDTO)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"messages": messages,
		},
	})
}

// HasDialog godoc
//
//	@Summary		Check if dialog exists
//	@Description	User can check if chat with that user_link exists
//	@Tags			Message
//	@Param			user_link	query		string	false	"User link"
//	@Success		200			{object}	doc.HasDialogResponse
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/im/chat/check [get]
func (handler *MessageHandler) HasDialog(ctx *gin.Context) {
	hasDialogDTO := dto.HasDialogDTO{}
	err := ctx.ShouldBind(&hasDialogDTO)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	hasDialog, err := handler.MessageUsecase.HasDialog(email.(string), &hasDialogDTO)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"has_dialog": hasDialog,
		},
	})
}
