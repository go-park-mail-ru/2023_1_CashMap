package usecase

import "depeche/internal/auth/repository"

type UseCase interface {
}

type AuthUseCase struct {
	repository repository.Repository
}
