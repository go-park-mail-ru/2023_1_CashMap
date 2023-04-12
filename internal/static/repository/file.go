package repository

import (
	"bufio"
	"context"
	"depeche/internal/static/entities"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	IMG_STATIC_PATH = "/internal/static/files/img/"
	DOC_STATIC_PATH = "/internal/static/files/doc/"
)

var BASE_PATH, _ = os.Getwd()

type FileRepository interface {
	ReadFile(file *entities.UserFile, outputStream io.Writer) error
	WriteFile(file *entities.UserFile, fileDescriptor io.ReadCloser, cancelCtx context.Context, finishCtx context.Context, errorChan chan error)
	DeleteFile(file *entities.UserFile) error
}

type FileStorage struct{}

func NewFileRepository() FileRepository {
	return &FileStorage{}
}

func (fs *FileStorage) ReadFile(file *entities.UserFile, outputStreamWriter io.Writer) error {
	var inputStream *os.File
	// TODO: Адекватно это обработать внутри дефера
	defer inputStream.Close()

	var err error

	switch file.Type {
	case entities.IMAGE:
		inputStream, err = os.Open(BASE_PATH + IMG_STATIC_PATH + file.Name)
	case entities.DOCUMENT:
		inputStream, err = os.Open(BASE_PATH + DOC_STATIC_PATH + file.Name)
	}

	if err != nil {
		return err
	}

	fileReader := bufio.NewReader(inputStream)

	_, err = io.Copy(outputStreamWriter, fileReader)
	if err != nil {
		// TODO: ВЫЛЕТАКЕТ invalid ARGUMENT
		fmt.Println(err)
		return err
	}

	return nil
}

func (fs *FileStorage) WriteFile(file *entities.UserFile, inputFileDescriptor io.ReadCloser, cancelCtx context.Context, finishCtx context.Context, errorChan chan error) {
	var outputFileDescriptor *os.File
	// TODO: Адекватно это обработать внутри дефера
	defer outputFileDescriptor.Close()
	defer inputFileDescriptor.Close()

	var err error

	var filePath string
	switch file.Type {
	case entities.IMAGE:
		filePath = BASE_PATH + IMG_STATIC_PATH + file.Name
	case entities.DOCUMENT:
		filePath = BASE_PATH + DOC_STATIC_PATH + file.Name
	}
	outputFileDescriptor, err = os.Create(filePath)

	if err != nil {
		errorChan <- err
		return
	}

	readWriter := bufio.NewReadWriter(bufio.NewReader(inputFileDescriptor), bufio.NewWriter(outputFileDescriptor))

	errorChan <- nil
	select {
	case <-cancelCtx.Done():
		err = os.Remove(filePath)
		if err != nil {
			errorChan <- err
		}
	case <-finishCtx.Done():
		// TODO: Удалять файл, если тут вылетела ошибка
		_, err := io.Copy(readWriter.Writer, readWriter.Reader)
		if err != nil {
			return
		}

		err = readWriter.Flush()
		if err != nil {
			errorChan <- err
		}
	}
}

func (fs *FileStorage) DeleteFile(file *entities.UserFile) error {
	var filepath string
	switch file.Type {
	case entities.DOCUMENT:
		filepath = BASE_PATH + DOC_STATIC_PATH + file.Name
	case entities.IMAGE:
		filepath = BASE_PATH + IMG_STATIC_PATH + file.Name
	default:
		return errors.New("unknown file type")
	}

	err := os.Remove(filepath)
	if err != nil {
		return err
	}

	return nil
}
