// Code generated by MockGen. DO NOT EDIT.
// Source: depeche/internal/repository (interfaces: UserRepository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	dto "depeche/internal/delivery/dto"
	entities "depeche/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CheckLinkExists mocks base method.
func (m *MockUserRepository) CheckLinkExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckLinkExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckLinkExists indicates an expected call of CheckLinkExists.
func (mr *MockUserRepositoryMockRecorder) CheckLinkExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckLinkExists", reflect.TypeOf((*MockUserRepository)(nil).CheckLinkExists), arg0)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(arg0 *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), arg0)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), arg0)
}

// GetFriends mocks base method.
func (m *MockUserRepository) GetFriends(arg0 *entities.User, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriends", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriends indicates an expected call of GetFriends.
func (mr *MockUserRepositoryMockRecorder) GetFriends(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriends", reflect.TypeOf((*MockUserRepository)(nil).GetFriends), arg0, arg1, arg2)
}

// GetPendingFriendRequests mocks base method.
func (m *MockUserRepository) GetPendingFriendRequests(arg0 *entities.User, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingFriendRequests", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingFriendRequests indicates an expected call of GetPendingFriendRequests.
func (mr *MockUserRepositoryMockRecorder) GetPendingFriendRequests(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingFriendRequests", reflect.TypeOf((*MockUserRepository)(nil).GetPendingFriendRequests), arg0, arg1, arg2)
}

// GetSubscribers mocks base method.
func (m *MockUserRepository) GetSubscribers(arg0 *entities.User, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers.
func (mr *MockUserRepositoryMockRecorder) GetSubscribers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockUserRepository)(nil).GetSubscribers), arg0, arg1, arg2)
}

// GetSubscribes mocks base method.
func (m *MockUserRepository) GetSubscribes(arg0 *entities.User, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribes", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribes indicates an expected call of GetSubscribes.
func (mr *MockUserRepositoryMockRecorder) GetSubscribes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribes", reflect.TypeOf((*MockUserRepository)(nil).GetSubscribes), arg0, arg1, arg2)
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(arg0 string, arg1 ...interface{}) (*entities.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), varargs...)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepository) GetUserByEmail(arg0 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), arg0)
}

// GetUserById mocks base method.
func (m *MockUserRepository) GetUserById(arg0 uint) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserRepositoryMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserRepository)(nil).GetUserById), arg0)
}

// GetUserByLink mocks base method.
func (m *MockUserRepository) GetUserByLink(arg0 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLink", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByLink indicates an expected call of GetUserByLink.
func (mr *MockUserRepositoryMockRecorder) GetUserByLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLink", reflect.TypeOf((*MockUserRepository)(nil).GetUserByLink), arg0)
}

// GetUsers mocks base method.
func (m *MockUserRepository) GetUsers(arg0 string, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserRepositoryMockRecorder) GetUsers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserRepository)(nil).GetUsers), arg0, arg1, arg2)
}

// HasPendingRequest mocks base method.
func (m *MockUserRepository) HasPendingRequest(arg0, arg1 *entities.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPendingRequest", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasPendingRequest indicates an expected call of HasPendingRequest.
func (mr *MockUserRepositoryMockRecorder) HasPendingRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPendingRequest", reflect.TypeOf((*MockUserRepository)(nil).HasPendingRequest), arg0, arg1)
}

// IsFriend mocks base method.
func (m *MockUserRepository) IsFriend(arg0, arg1 *entities.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFriend", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFriend indicates an expected call of IsFriend.
func (mr *MockUserRepositoryMockRecorder) IsFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFriend", reflect.TypeOf((*MockUserRepository)(nil).IsFriend), arg0, arg1)
}

// IsSubscriber mocks base method.
func (m *MockUserRepository) IsSubscriber(arg0, arg1 *entities.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSubscriber", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSubscriber indicates an expected call of IsSubscriber.
func (mr *MockUserRepositoryMockRecorder) IsSubscriber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSubscriber", reflect.TypeOf((*MockUserRepository)(nil).IsSubscriber), arg0, arg1)
}

// RejectFriendRequest mocks base method.
func (m *MockUserRepository) RejectFriendRequest(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RejectFriendRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RejectFriendRequest indicates an expected call of RejectFriendRequest.
func (mr *MockUserRepositoryMockRecorder) RejectFriendRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectFriendRequest", reflect.TypeOf((*MockUserRepository)(nil).RejectFriendRequest), arg0, arg1)
}

// Subscribe mocks base method.
func (m *MockUserRepository) Subscribe(arg0, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockUserRepositoryMockRecorder) Subscribe(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockUserRepository)(nil).Subscribe), arg0, arg1, arg2)
}

// Unsubscribe mocks base method.
func (m *MockUserRepository) Unsubscribe(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockUserRepositoryMockRecorder) Unsubscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockUserRepository)(nil).Unsubscribe), arg0, arg1)
}

// UpdateAvatar mocks base method.
func (m *MockUserRepository) UpdateAvatar(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockUserRepositoryMockRecorder) UpdateAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockUserRepository)(nil).UpdateAvatar), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(arg0 string, arg1 *dto.EditProfile) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), arg0, arg1)
}