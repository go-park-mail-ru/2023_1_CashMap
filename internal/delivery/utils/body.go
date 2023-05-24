package utils

import (
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
)

func GetBody[T any](ctx *gin.Context) (*T, error) {
	var request = struct {
		Data *T `json:"body"`
	}{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		return nil, apperror.NewServerError(apperror.BadRequest, err)
	}

	if request.Data == nil {
		return nil, apperror.NewBadRequest()
	}
	return request.Data, nil
}
