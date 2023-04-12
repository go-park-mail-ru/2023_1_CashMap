package service

import (
	"context"
	"depeche/internal/static/entities"
	"depeche/internal/static/repository"
	"errors"
	"github.com/google/uuid"
	"io"
	"strings"
	"sync"
)

type FileUsecase interface {
	CreateFile(userFile []*entities.UserFile, fileDescriptors []io.ReadCloser) ([]*entities.UserFile, error)
	ReadFile(file *entities.UserFile, outputStreamWriter io.Writer) error
	DeleteFile(file *entities.UserFile) error
}

type FileService struct {
	repository.FileRepository
}

func NewFileUsecase() FileUsecase {
	return &FileService{
		repository.NewFileRepository(),
	}
}

func (fileService *FileService) CreateFile(userFiles []*entities.UserFile, inputStreams []io.ReadCloser) ([]*entities.UserFile, error) {
	wg := sync.WaitGroup{}
	cancelCtx, cancelWriting := context.WithCancel(context.Background())
	finishCtx, finishWriting := context.WithCancel(context.Background())
	defer cancelWriting()
	defer finishWriting()

	errorChan := make(chan error, len(userFiles))
	defer close(errorChan)

	for i := 0; i < len(userFiles); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			inputStream := inputStreams[i]
			userFile := userFiles[i]
			uid, err := uuid.NewUUID()
			if err != nil {
				errorChan <- err
				return
			}
			splitFilename := strings.Split(userFile.Name, ".")
			if len(splitFilename) == 1 {
				errorChan <- errors.New("invalid file type (empty extension)")
				return
			}

			extension := splitFilename[len(splitFilename)-1]
			uniqueFilename := uid.String() + "." + extension
			userFile.Name = uniqueFilename

			fileService.FileRepository.WriteFile(userFile, inputStream, cancelCtx, finishCtx, errorChan)
		}(i)
	}

	var err error = nil
	for i := 0; i < len(userFiles); i++ {
		err = <-errorChan
		if err != nil {
			cancelWriting()
			break
		}
	}

	if err == nil {
		finishWriting()
	}

	wg.Wait()

	return userFiles, err
}

func (fileService *FileService) ReadFile(file *entities.UserFile, outputStreamWriter io.Writer) error {
	return fileService.FileRepository.ReadFile(file, outputStreamWriter)
}

func (fileService *FileService) DeleteFile(file *entities.UserFile) error {
	return fileService.FileRepository.DeleteFile(file)
}
