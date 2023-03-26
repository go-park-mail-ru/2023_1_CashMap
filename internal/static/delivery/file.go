package delivery

import (
	"depeche/internal/static/service"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

var BASE_PATH, _ = os.Getwd()

const (
	IMG_STATIC_PATH = "/internal/static/files/img/"
	DOC_STATIC_PATH = "/internal/static/files/doc/"
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

	// TOOO: ОТФИЛЬТРОВАТЬ КАЧЕСТВЕННО КОНТЕНТ ТУПЕ
	// TODO: ЗАХШИРОВАТЬ ИМЯ ФАЙЛА
	for _, file := range files.File["img"] {
		filename := file.Filename
		if contentType := file.Header.Get("Content-Type"); strings.Contains(contentType, "image") {
			err = ctx.SaveUploadedFile(file, BASE_PATH+IMG_STATIC_PATH+filename)
		} else if contentType == "" {
			ctx.Error(apperror.BadRequest)
			return
		} else {
			err = ctx.SaveUploadedFile(file, BASE_PATH+DOC_STATIC_PATH+filename)
		}

		if err != nil {
			fmt.Println(err)
			ctx.Error(apperror.InternalServerError)
			return
		}
	}
}

func (fileHandler *FileHandler) GetFile(ctx *gin.Context) {
	ctx.Error(apperror.Forbidden)
	return
}
