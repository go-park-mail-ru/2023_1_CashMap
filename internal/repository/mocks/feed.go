// Code generated by MockGen. DO NOT EDIT.
// Source: feed.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	dto "depeche/internal/delivery/dto"
	entities "depeche/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFeedRepository is a mock of FeedRepository interface.
type MockFeedRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFeedRepositoryMockRecorder
}

// MockFeedRepositoryMockRecorder is the mock recorder for MockFeedRepository.
type MockFeedRepositoryMockRecorder struct {
	mock *MockFeedRepository
}

// NewMockFeedRepository creates a new mock instance.
func NewMockFeedRepository(ctrl *gomock.Controller) *MockFeedRepository {
	mock := &MockFeedRepository{ctrl: ctrl}
	mock.recorder = &MockFeedRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeedRepository) EXPECT() *MockFeedRepositoryMockRecorder {
	return m.recorder
}

// GetFriendsPosts mocks base method.
func (m *MockFeedRepository) GetFriendsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriendsPosts", email, feedDTO)
	ret0, _ := ret[0].([]*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriendsPosts indicates an expected call of GetFriendsPosts.
func (mr *MockFeedRepositoryMockRecorder) GetFriendsPosts(email, feedDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendsPosts", reflect.TypeOf((*MockFeedRepository)(nil).GetFriendsPosts), email, feedDTO)
}

// GetGroupsPosts mocks base method.
func (m *MockFeedRepository) GetGroupsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupsPosts", email, feedDTO)
	ret0, _ := ret[0].([]*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupsPosts indicates an expected call of GetGroupsPosts.
func (mr *MockFeedRepositoryMockRecorder) GetGroupsPosts(email, feedDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupsPosts", reflect.TypeOf((*MockFeedRepository)(nil).GetGroupsPosts), email, feedDTO)
}
