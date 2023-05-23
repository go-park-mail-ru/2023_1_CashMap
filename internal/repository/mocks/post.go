// Code generated by MockGen. DO NOT EDIT.
// Source: depeche/internal/repository (interfaces: PostRepository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	dto "depeche/internal/delivery/dto"
	entities "depeche/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// AddPostAttachments mocks base method.
func (m *MockPostRepository) AddPostAttachments(arg0 uint, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPostAttachments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPostAttachments indicates an expected call of AddPostAttachments.
func (mr *MockPostRepositoryMockRecorder) AddPostAttachments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPostAttachments", reflect.TypeOf((*MockPostRepository)(nil).AddPostAttachments), arg0, arg1)
}

// CancelLike mocks base method.
func (m *MockPostRepository) CancelLike(arg0 string, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelLike", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelLike indicates an expected call of CancelLike.
func (mr *MockPostRepositoryMockRecorder) CancelLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelLike", reflect.TypeOf((*MockPostRepository)(nil).CancelLike), arg0, arg1)
}

// CheckReadAccess mocks base method.
func (m *MockPostRepository) CheckReadAccess(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReadAccess", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckReadAccess indicates an expected call of CheckReadAccess.
func (mr *MockPostRepositoryMockRecorder) CheckReadAccess(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReadAccess", reflect.TypeOf((*MockPostRepository)(nil).CheckReadAccess), arg0)
}

// CheckWriteAccess mocks base method.
func (m *MockPostRepository) CheckWriteAccess(arg0 string, arg1 *dto.PostCreate) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckWriteAccess", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckWriteAccess indicates an expected call of CheckWriteAccess.
func (mr *MockPostRepositoryMockRecorder) CheckWriteAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckWriteAccess", reflect.TypeOf((*MockPostRepository)(nil).CheckWriteAccess), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockPostRepository) CreatePost(arg0 string, arg1 *dto.PostCreate) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostRepositoryMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostRepository)(nil).CreatePost), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockPostRepository) DeletePost(arg0 string, arg1 *dto.PostDelete) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostRepositoryMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostRepository)(nil).DeletePost), arg0, arg1)
}

// GetLikesAmount mocks base method.
func (m *MockPostRepository) GetLikesAmount(arg0 string, arg1 uint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikesAmount", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikesAmount indicates an expected call of GetLikesAmount.
func (mr *MockPostRepositoryMockRecorder) GetLikesAmount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikesAmount", reflect.TypeOf((*MockPostRepository)(nil).GetLikesAmount), arg0, arg1)
}

// GetPostAttachments mocks base method.
func (m *MockPostRepository) GetPostAttachments(arg0 uint) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostAttachments", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostAttachments indicates an expected call of GetPostAttachments.
func (mr *MockPostRepositoryMockRecorder) GetPostAttachments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostAttachments", reflect.TypeOf((*MockPostRepository)(nil).GetPostAttachments), arg0)
}

// GetPostSenderInfo mocks base method.
func (m *MockPostRepository) GetPostSenderInfo(arg0 uint) (*entities.UserInfo, *entities.CommunityInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostSenderInfo", arg0)
	ret0, _ := ret[0].(*entities.UserInfo)
	ret1, _ := ret[1].(*entities.CommunityInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPostSenderInfo indicates an expected call of GetPostSenderInfo.
func (mr *MockPostRepositoryMockRecorder) GetPostSenderInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostSenderInfo", reflect.TypeOf((*MockPostRepository)(nil).GetPostSenderInfo), arg0)
}

// SelectPostById mocks base method.
func (m *MockPostRepository) SelectPostById(arg0 uint, arg1 string) (*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectPostById", arg0, arg1)
	ret0, _ := ret[0].(*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectPostById indicates an expected call of SelectPostById.
func (mr *MockPostRepositoryMockRecorder) SelectPostById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectPostById", reflect.TypeOf((*MockPostRepository)(nil).SelectPostById), arg0, arg1)
}

// SelectPostsByCommunityLink mocks base method.
func (m *MockPostRepository) SelectPostsByCommunityLink(arg0 *dto.PostsGetByLink, arg1 string) ([]*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectPostsByCommunityLink", arg0, arg1)
	ret0, _ := ret[0].([]*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectPostsByCommunityLink indicates an expected call of SelectPostsByCommunityLink.
func (mr *MockPostRepositoryMockRecorder) SelectPostsByCommunityLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectPostsByCommunityLink", reflect.TypeOf((*MockPostRepository)(nil).SelectPostsByCommunityLink), arg0, arg1)
}

// SelectPostsByUserLink mocks base method.
func (m *MockPostRepository) SelectPostsByUserLink(arg0 *dto.PostsGetByLink, arg1 string) ([]*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectPostsByUserLink", arg0, arg1)
	ret0, _ := ret[0].([]*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectPostsByUserLink indicates an expected call of SelectPostsByUserLink.
func (mr *MockPostRepositoryMockRecorder) SelectPostsByUserLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectPostsByUserLink", reflect.TypeOf((*MockPostRepository)(nil).SelectPostsByUserLink), arg0, arg1)
}

// SetLike mocks base method.
func (m *MockPostRepository) SetLike(arg0 string, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLike", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLike indicates an expected call of SetLike.
func (mr *MockPostRepositoryMockRecorder) SetLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLike", reflect.TypeOf((*MockPostRepository)(nil).SetLike), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockPostRepository) UpdatePost(arg0 string, arg1 *dto.PostUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostRepositoryMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostRepository)(nil).UpdatePost), arg0, arg1)
}

// UpdatePostAttachments mocks base method.
func (m *MockPostRepository) UpdatePostAttachments(arg0 uint, arg1 *dto.UpdateAttachments) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePostAttachments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePostAttachments indicates an expected call of UpdatePostAttachments.
func (mr *MockPostRepositoryMockRecorder) UpdatePostAttachments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePostAttachments", reflect.TypeOf((*MockPostRepository)(nil).UpdatePostAttachments), arg0, arg1)
}
