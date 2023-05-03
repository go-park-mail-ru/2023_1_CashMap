package delivery

import (
	"depeche/internal/static/entities"
	"depeche/internal/static/service"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const (
	MAX_IMG_SIZE = 50 * 1048576  // 50 MB
	MAX_DOC_SIZE = 200 * 1048576 // 200 MB
)

type FileHandler struct {
	FileUsecase service.FileUsecase
}

func NewFileHandler(fileUsecase service.FileUsecase) *FileHandler {
	return &FileHandler{
		FileUsecase: fileUsecase,
	}
}

// LoadFile godoc
//
//	@Summary		Load file on server
//	@Description	Users can upload many files using multipart/form-data
//	@Tags			static
//	@Produce		json
//	@Success		200	{object}	entities.UserFile
//	@Failure		400
//	@Failure		413
//	@Failure		500
//	@Router			/api/static/upload [post]
func (fileHandler *FileHandler) LoadFile(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse multipart-form")))
		return
	}

	var inputFilesReadStreams []io.ReadCloser

	var inputFiles []*entities.UserFile
	if len(form.File["attachments"]) == 0 {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("attachments are required")))
		return
	}

	for _, file := range form.File["attachments"] {
		var userFile = new(entities.UserFile)
		userFile.Name = file.Filename
		contentType := file.Header.Get("Content-Type")
		switch {
		case contentType == "":
			_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("content type is required")))
			return
		case strings.Contains(contentType, "image/"):
			if file.Size > MAX_IMG_SIZE {
				errorMsg := "max image size is " + strconv.FormatFloat(MAX_IMG_SIZE/math.Pow(2, 20), 'f', 1, 64) + "МБ"
				_ = ctx.Error(apperror.NewServerError(apperror.TooLargePayload, errors.New(errorMsg)))
				return
			}
			userFile.Type = entities.IMAGE
		default:
			splitFilename := strings.Split(userFile.Name, ".")
			extension := splitFilename[len(splitFilename)-1]
			if extension == "exe" {
				errorMsg := "exe files are not alowed"
				_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New(errorMsg)))
				return
			}

			if file.Size > MAX_DOC_SIZE {
				errorMsg := "max document size is " + strconv.FormatFloat(MAX_DOC_SIZE/math.Pow(2, 20), 'f', 1, 64) + "МБ"
				_ = ctx.Error(apperror.NewServerError(apperror.TooLargePayload, errors.New(errorMsg)))
				return
			}
			userFile.Type = entities.DOCUMENT
		}

		if err != nil {
			_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
			return
		}

		inputFiles = append(inputFiles, userFile)

		inputFileStream, err := file.Open()
		if err != nil {
			fmt.Println(6)
			_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
			return
		}
		inputFilesReadStreams = append(inputFilesReadStreams, inputFileStream.(io.ReadCloser))
	}

	outputFiles, err := fileHandler.FileUsecase.CreateFile(inputFiles, inputFilesReadStreams)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"form": outputFiles,
		},
		"status": http.StatusOK,
	})
}

// TODO: Сделать поддержку чтения сразу нескольких файлов

// GetFile godoc
//
//	@Summary		Read file from server
//	@Description	Users can read file on server
//	@Tags			static
//	@Produce		octet-stream
//	@Param			name	query	string	true	"File name"
//	@Param			type	query	string	true	"File type ("doc" or "img")"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/api/static/download [get]
func (fileHandler *FileHandler) GetFile(ctx *gin.Context) {
	var file entities.UserFile
	err := ctx.ShouldBind(&file)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	err = fileHandler.FileUsecase.ReadFile(&file, ctx.Writer)
	if err != nil {
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

// DeleteFile godoc
//
//	@Summary		Delete file from server
//	@Description	Users can delete file on server
//	@Tags			static
//	@Produce		json
//	@Param			name	query	string	true	"File name"
//	@Param			type	query	string	true	"File type ("doc" or "img")"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/api/static/delete [delete]
func (fileHandler *FileHandler) DeleteFile(ctx *gin.Context) {
	var file entities.UserFile
	err := ctx.ShouldBind(&file)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	err = fileHandler.FileUsecase.DeleteFile(&file)
	if err != nil {
		_ = ctx.Error(apperror.InternalServerError)
		return
	}
}
