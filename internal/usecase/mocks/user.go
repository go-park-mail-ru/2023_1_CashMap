// Code generated by MockGen. DO NOT EDIT.
// Source: depeche/internal/usecase (interfaces: User)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	dto "depeche/internal/delivery/dto"
	entities "depeche/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// EditProfile mocks base method.
func (m *MockUser) EditProfile(arg0 string, arg1 *dto.EditProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditProfile indicates an expected call of EditProfile.
func (mr *MockUserMockRecorder) EditProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfile", reflect.TypeOf((*MockUser)(nil).EditProfile), arg0, arg1)
}

// GetAllUsers mocks base method.
func (m *MockUser) GetAllUsers(arg0 string, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserMockRecorder) GetAllUsers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUser)(nil).GetAllUsers), arg0, arg1, arg2)
}

// GetFriendsByEmail mocks base method.
func (m *MockUser) GetFriendsByEmail(arg0 string, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriendsByEmail", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriendsByEmail indicates an expected call of GetFriendsByEmail.
func (mr *MockUserMockRecorder) GetFriendsByEmail(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendsByEmail", reflect.TypeOf((*MockUser)(nil).GetFriendsByEmail), arg0, arg1, arg2)
}

// GetFriendsByLink mocks base method.
func (m *MockUser) GetFriendsByLink(arg0, arg1 string, arg2, arg3 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriendsByLink", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriendsByLink indicates an expected call of GetFriendsByLink.
func (mr *MockUserMockRecorder) GetFriendsByLink(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendsByLink", reflect.TypeOf((*MockUser)(nil).GetFriendsByLink), arg0, arg1, arg2, arg3)
}

// GetPendingRequestsByEmail mocks base method.
func (m *MockUser) GetPendingRequestsByEmail(arg0 string, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingRequestsByEmail", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingRequestsByEmail indicates an expected call of GetPendingRequestsByEmail.
func (mr *MockUserMockRecorder) GetPendingRequestsByEmail(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingRequestsByEmail", reflect.TypeOf((*MockUser)(nil).GetPendingRequestsByEmail), arg0, arg1, arg2)
}

// GetProfileByEmail mocks base method.
func (m *MockUser) GetProfileByEmail(arg0 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileByEmail", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileByEmail indicates an expected call of GetProfileByEmail.
func (mr *MockUserMockRecorder) GetProfileByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileByEmail", reflect.TypeOf((*MockUser)(nil).GetProfileByEmail), arg0)
}

// GetProfileByLink mocks base method.
func (m *MockUser) GetProfileByLink(arg0, arg1 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileByLink", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileByLink indicates an expected call of GetProfileByLink.
func (mr *MockUserMockRecorder) GetProfileByLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileByLink", reflect.TypeOf((*MockUser)(nil).GetProfileByLink), arg0, arg1)
}

// GetSubscribersByEmail mocks base method.
func (m *MockUser) GetSubscribersByEmail(arg0 string, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribersByEmail", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribersByEmail indicates an expected call of GetSubscribersByEmail.
func (mr *MockUserMockRecorder) GetSubscribersByEmail(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribersByEmail", reflect.TypeOf((*MockUser)(nil).GetSubscribersByEmail), arg0, arg1, arg2)
}

// GetSubscribersByLink mocks base method.
func (m *MockUser) GetSubscribersByLink(arg0, arg1 string, arg2, arg3 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribersByLink", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribersByLink indicates an expected call of GetSubscribersByLink.
func (mr *MockUserMockRecorder) GetSubscribersByLink(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribersByLink", reflect.TypeOf((*MockUser)(nil).GetSubscribersByLink), arg0, arg1, arg2, arg3)
}

// GetSubscribesByEmail mocks base method.
func (m *MockUser) GetSubscribesByEmail(arg0 string, arg1, arg2 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribesByEmail", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribesByEmail indicates an expected call of GetSubscribesByEmail.
func (mr *MockUserMockRecorder) GetSubscribesByEmail(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribesByEmail", reflect.TypeOf((*MockUser)(nil).GetSubscribesByEmail), arg0, arg1, arg2)
}

// GetSubscribesByLink mocks base method.
func (m *MockUser) GetSubscribesByLink(arg0, arg1 string, arg2, arg3 int) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribesByLink", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribesByLink indicates an expected call of GetSubscribesByLink.
func (mr *MockUserMockRecorder) GetSubscribesByLink(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribesByLink", reflect.TypeOf((*MockUser)(nil).GetSubscribesByLink), arg0, arg1, arg2, arg3)
}

// Reject mocks base method.
func (m *MockUser) Reject(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reject indicates an expected call of Reject.
func (mr *MockUserMockRecorder) Reject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockUser)(nil).Reject), arg0, arg1)
}

// SignIn mocks base method.
func (m *MockUser) SignIn(arg0 *dto.SignIn) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserMockRecorder) SignIn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUser)(nil).SignIn), arg0)
}

// SignUp mocks base method.
func (m *MockUser) SignUp(arg0 *dto.SignUp) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserMockRecorder) SignUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUser)(nil).SignUp), arg0)
}

// Subscribe mocks base method.
func (m *MockUser) Subscribe(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockUserMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockUser)(nil).Subscribe), arg0, arg1)
}

// Unsubscribe mocks base method.
func (m *MockUser) Unsubscribe(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockUserMockRecorder) Unsubscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockUser)(nil).Unsubscribe), arg0, arg1)
}