package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/entities/response"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
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
	inputDTO := new(response.SendRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	msg, err := handler.MessageUsecase.Send(email, inputDTO.Body)
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
	inputDTO := new(response.NewChatRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	chat, err := handler.MessageUsecase.CreateChat(email.(string), inputDTO.Body)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	_response := response.NewChatResponse{
		Body: response.NewChatBody{
			Chat: chat,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)

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

	_response := response.GetChatsResponse{
		Body: response.GetChatsBody{
			Chats: chats,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

// GetMessagesByChatID godoc
//
//	@Summary		Get messages
//	@Description	Get messages batch by chatID sorted by date
//	@Tags			Message
//	@Param			chat_id			query		uint	false	"Chat id"
//	@Param			batch_size		query		uint	false	"Batch size"
//	@Param			last_msg_date	query		uint	false	"Last message date"
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

	messages, hasNextMessages, err := handler.MessageUsecase.GetMessagesByChatID(email.(string), &getMsgDTO)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	_response := response.GetMessagesByChatIDResponse{
		Body: response.GetMessagesByChatIDBody{
			Messages: messages,
			HasNext:  hasNextMessages,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
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

	chatId, err := handler.MessageUsecase.HasDialog(email.(string), &hasDialogDTO)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	if chatId != nil {
		_response := response.HasDialogResponse{
			Body: response.HasDialogBody{
				ChatId:    *chatId,
				HasDialog: true,
			},
		}

		responseJSON, err := _response.MarshalJSON()
		if err != nil {
			_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
			return
		}

		ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
		return
	}

	_response := response.HasDialogResponse{
		Body: response.HasDialogBody{
			ChatId:    0,
			HasDialog: false,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)

}

func (handler *MessageHandler) GetUnreadChatsCount(ctx *gin.Context) {
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	count, err := handler.MessageUsecase.GetUnreadChatsCount(email)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetUnreadChatsCountResponse{
		Body: response.GetUnreadChatsCountBody{
			Count: count,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (handler *MessageHandler) SetLastRead(ctx *gin.Context) {
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	inputDTO := new(dto.SetLastReadTime)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	err = handler.MessageUsecase.SetLastRead(email, inputDTO.ChatID, inputDTO.Time)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}
}
