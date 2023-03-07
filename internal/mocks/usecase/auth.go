package usecase

import (
	"depeche/internal/entities"
	"depeche/internal/session"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Authenticate(user *entities.User) (string, error) {
	returned := m.Called(user)

	var r0 string
	if returned.Get(0) != nil {
		r0 = returned.Get(0).(string)
	}

	var r1 error
	if returned.Get(1) != nil {
		r1 = returned.Get(1).(error)
	}

	return r0, r1
}

func (m *MockAuthService) LogOut(token string) error {
	returned := m.Called(token)

	var r0 error
	if returned.Get(1) != nil {
		r0 = returned.Get(1).(error)
	}

	return r0
}

func (m *MockAuthService) CheckSession(token string) (*session.Session, error) {
	returned := m.Called(token)

	var r0 *session.Session
	if returned.Get(0) != nil {
		r0 = returned.Get(0).(*session.Session)
	}

	var r1 error
	if returned.Get(1) != nil {
		r1 = returned.Get(1).(error)
	}

	return r0, r1
}
