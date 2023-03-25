package delivery

import (
	"depeche/internal/static/service"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

const (
	IMG_STATIC_PATH = "../files/img"
	DOC_STATIC_PATH = "../files/doc"
)

type FileHandler struct {
	FileUsecase service.FileUsecase
}

func NewFileHandler(fileUsecase service.FileUsecase) *FileHandler {
	return &FileHandler{
		FileUsecase: fileUsecase,
	}
}

func (fileHandler *FileHandler) LoadFile(ctx *gin.Context) {
	files, err := ctx.MultipartForm()
	if err != nil {
		ctx.Error(apperror.BadRequest)
		return
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	for _, file := range files.File["img"] {
		err = ctx.SaveUploadedFile(file, IMG_STATIC_PATH)
		if err != nil {
			ctx.Error(apperror.InternalServerError)
			return
		}
	}
}
