// Code generated by MockGen. DO NOT EDIT.
// Source: depeche/internal/usecase (interfaces: MessageUsecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	dto "depeche/internal/delivery/dto"
	entities "depeche/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMessageUsecase is a mock of MessageUsecase interface.
type MockMessageUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockMessageUsecaseMockRecorder
}

// MockMessageUsecaseMockRecorder is the mock recorder for MockMessageUsecase.
type MockMessageUsecaseMockRecorder struct {
	mock *MockMessageUsecase
}

// NewMockMessageUsecase creates a new mock instance.
func NewMockMessageUsecase(ctrl *gomock.Controller) *MockMessageUsecase {
	mock := &MockMessageUsecase{ctrl: ctrl}
	mock.recorder = &MockMessageUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageUsecase) EXPECT() *MockMessageUsecaseMockRecorder {
	return m.recorder
}

// CreateChat mocks base method.
func (m *MockMessageUsecase) CreateChat(arg0 string, arg1 *dto.CreateChatDTO) (*entities.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChat", arg0, arg1)
	ret0, _ := ret[0].(*entities.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChat indicates an expected call of CreateChat.
func (mr *MockMessageUsecaseMockRecorder) CreateChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockMessageUsecase)(nil).CreateChat), arg0, arg1)
}

// GetChatsList mocks base method.
func (m *MockMessageUsecase) GetChatsList(arg0 string, arg1 *dto.GetChatsDTO) ([]*entities.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatsList", arg0, arg1)
	ret0, _ := ret[0].([]*entities.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatsList indicates an expected call of GetChatsList.
func (mr *MockMessageUsecaseMockRecorder) GetChatsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatsList", reflect.TypeOf((*MockMessageUsecase)(nil).GetChatsList), arg0, arg1)
}

// GetMembersByChatId mocks base method.
func (m *MockMessageUsecase) GetMembersByChatId(arg0 uint) ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembersByChatId", arg0)
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembersByChatId indicates an expected call of GetMembersByChatId.
func (mr *MockMessageUsecaseMockRecorder) GetMembersByChatId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembersByChatId", reflect.TypeOf((*MockMessageUsecase)(nil).GetMembersByChatId), arg0)
}

// GetMessagesByChatID mocks base method.
func (m *MockMessageUsecase) GetMessagesByChatID(arg0 string, arg1 *dto.GetMessagesDTO) ([]*entities.Message, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessagesByChatID", arg0, arg1)
	ret0, _ := ret[0].([]*entities.Message)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMessagesByChatID indicates an expected call of GetMessagesByChatID.
func (mr *MockMessageUsecaseMockRecorder) GetMessagesByChatID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessagesByChatID", reflect.TypeOf((*MockMessageUsecase)(nil).GetMessagesByChatID), arg0, arg1)
}

// GetUnreadChatsCount mocks base method.
func (m *MockMessageUsecase) GetUnreadChatsCount(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnreadChatsCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnreadChatsCount indicates an expected call of GetUnreadChatsCount.
func (mr *MockMessageUsecaseMockRecorder) GetUnreadChatsCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnreadChatsCount", reflect.TypeOf((*MockMessageUsecase)(nil).GetUnreadChatsCount), arg0)
}

// HasDialog mocks base method.
func (m *MockMessageUsecase) HasDialog(arg0 string, arg1 *dto.HasDialogDTO) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasDialog", arg0, arg1)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasDialog indicates an expected call of HasDialog.
func (mr *MockMessageUsecaseMockRecorder) HasDialog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasDialog", reflect.TypeOf((*MockMessageUsecase)(nil).HasDialog), arg0, arg1)
}

// Send mocks base method.
func (m *MockMessageUsecase) Send(arg0 string, arg1 *dto.NewMessageDTO) (*entities.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*entities.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockMessageUsecaseMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMessageUsecase)(nil).Send), arg0, arg1)
}

// SetLastRead mocks base method.
func (m *MockMessageUsecase) SetLastRead(arg0 string, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLastRead", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLastRead indicates an expected call of SetLastRead.
func (mr *MockMessageUsecaseMockRecorder) SetLastRead(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastRead", reflect.TypeOf((*MockMessageUsecase)(nil).SetLastRead), arg0, arg1, arg2)
}
