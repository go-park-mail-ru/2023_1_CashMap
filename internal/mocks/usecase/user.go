package usecase

import (
	"depeche/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) SignIn(user *entities.User) (*entities.User, error) {
	returned := m.Called(user)

	var r0 *entities.User
	if returned.Get(0) != nil {
		r0 = returned.Get(0).(*entities.User)
	}

	var r1 error
	if returned.Get(1) != nil {
		r1 = returned.Get(1).(error)
	}

	return r0, r1
}

func (m *MockUserService) SignUp(user *entities.User) (*entities.User, error) {
	returned := m.Called(user)

	var r0 *entities.User
	if returned.Get(0) != nil {
		r0 = returned.Get(0).(*entities.User)
	}

	var r1 error
	if returned.Get(1) != nil {
		r1 = returned.Get(1).(error)
	}

	return r0, r1
}
