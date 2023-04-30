package utils

import (
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLimitOffset(ctx *gin.Context) (int, int, error) {
	limitQ := ctx.Query("limit")
	offsetQ := ctx.Query("offset")

	limit, err := strconv.Atoi(limitQ)
	if err != nil {
		return 0, 0, apperror.NewServerError(apperror.BadRequest, nil)
	}
	offset, err := strconv.Atoi(offsetQ)
	if err != nil {
		return 0, 0, apperror.NewServerError(apperror.BadRequest, nil)
	}
	return limit, offset, nil
}

func GetEmail(ctx *gin.Context) (string, error) {
	e, ok := ctx.Get("email")
	if !ok {
		return "", apperror.NewServerError(apperror.NoAuth, nil)
	}
	email, ok := e.(string)
	if !ok {
		return "", apperror.NewServerError(apperror.BadRequest, nil)
	}
	return email, nil
}
