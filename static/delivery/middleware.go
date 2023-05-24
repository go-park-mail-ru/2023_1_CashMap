package delivery

import (
	"depeche/pkg/apperror"
	"depeche/static/service"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AuthMiddleware(authUsecase service.AuthUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := ctx.Cookie("session_id")
		if err != nil {
			err = apperror.NewServerError(apperror.NoAuth, errors.New("auth cookie were not provided"))
			log.Error(err)

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  apperror.Errors[apperror.NoAuth].Code,
				"message": apperror.Errors[apperror.NoAuth].Message,
			})

			return
		}

		err = authUsecase.CheckSession(session)
		if err != nil {
			log.Error(err)

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  apperror.Errors[apperror.NoAuth].Code,
				"message": apperror.Errors[apperror.NoAuth].Message,
			})

			return
		}
	}
}
