package service

import (
	"depeche/internal/static/entities"
)

type FileUsecase interface {
	CreateFile(filename string) (*entities.UserFile, error)
	ReadFile(file *entities.UserFile) error
	UpdateFile()
	DeleteFile()
}

type FileService struct {
}

func NewFileUsecase() FileUsecase {
	return &FileService{}
}

func (fileService *FileService) CreateFile(filename string) (*entities.UserFile, error) {
	return nil, nil
}

func (fileService *FileService) ReadFile(file *entities.UserFile) error {
	return nil
}

func (fileService *FileService) UpdateFile() {

}

func (fileService *FileService) DeleteFile() {

}
