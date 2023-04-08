package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service usecase.Message
}

func NewMessageHandler(service usecase.Message) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
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

	msg, err := mh.service.Send(request.Data)
	if err != nil {
		_ = ctx.Error(err)
	}
	
	ctx.Set("message", msg)
}
