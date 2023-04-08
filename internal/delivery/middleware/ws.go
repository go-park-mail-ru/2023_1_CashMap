package middleware

import (
	"depeche/internal/delivery/wsPool"
	"depeche/internal/entities"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type WsMiddleware struct {
	pool       *wsPool.ConnectionPool
	msgService usecase.MessageUsecase
}

func NewWsMiddleware(pool *wsPool.ConnectionPool, msgService usecase.MessageUsecase) *WsMiddleware {
	return &WsMiddleware{
		pool:       pool,
		msgService: msgService,
	}
}

func (wm *WsMiddleware) SendMsg(ctx *gin.Context) {
	if len(ctx.Errors) > 0 {
		return
	}

	emailRaw, ok := ctx.Get("email")
	email, ok := emailRaw.(string)
	if !ok {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	msgRaw, ok := ctx.Get("message")
	if !ok {
		_ = ctx.Error(apperror.BadRequest)
	}
	msg, ok := msgRaw.(*entities.Message)
	if !ok {
		_ = ctx.Error(apperror.BadRequest)
	}

	raw, err := json.Marshal(msg)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
	}

	users, err := wm.msgService.GetMembersByChatId(*msg.ChatId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	go func(users []*entities.User, sender string, msg []byte) {
		for _, user := range users {
			err = wm.pool.SendMsg(user.Email, msg)
			if err != nil {
				// TODO log
			}
		}
	}(users, email, raw)

}
