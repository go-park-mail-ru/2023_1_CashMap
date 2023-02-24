package repository

type Repository interface {
	GetPasswordHash(login string) (string, error)
}
