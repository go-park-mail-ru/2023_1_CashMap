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

func (mh *MessageHandler) Send(ctx *gin.Context) {
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

	msg, err := mh.MessageUsecase.Send(request.Data)
	if err != nil {
		_ = ctx.Error(err)
	}

	ctx.Set("message", msg)
}

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
