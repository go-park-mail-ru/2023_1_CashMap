package middleware

import (
	"depeche/authorization_ms/authEntities"
	authService "depeche/authorization_ms/service"
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
)

type CSRFMiddleware struct {
	service authService.CSRFUsecase
}

func NewCSRFMiddleware(csrfService authService.CSRFUsecase) *CSRFMiddleware {
	return &CSRFMiddleware{csrfService}
}

func (csrf *CSRFMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.FullPath() == "/api/ws" {
			return
		}
		csrfToken := ctx.Request.Header.Get("X-Csrf-Token")
		if csrfToken == "" {
			err := apperror.Forbidden
			RejectInMiddleware(ctx, err)
			return
		}

		email, exists := ctx.Get("email")
		if !exists {
			err := apperror.Forbidden
			RejectInMiddleware(ctx, err)
			return
		}

		csrfData := authEntities.CSRF{
			Token: csrfToken,
			Email: email.(string),
		}

		_, err := csrf.service.ValidateCSRFToken(&csrfData)
		if err != nil {
			err := apperror.Forbidden
			RejectInMiddleware(ctx, err)
			return
		}
	}
}

func RejectInMiddleware(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(apperror.Errors[err].Code, gin.H{
		"status":  apperror.Errors[err].Code,
		"message": apperror.Errors[err].Message,
	})
}
