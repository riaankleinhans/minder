// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go
//
// Generated by this command:
//
//	mockgen -package mock_github -destination=./mock/service.go -source=./service.go
//

// Package mock_github is a generated GoMock package.
package mock_github

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	db "github.com/stacklok/minder/internal/db"
	v1 "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryService is a mock of RepositoryService interface.
type MockRepositoryService struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryServiceMockRecorder
}

// MockRepositoryServiceMockRecorder is the mock recorder for MockRepositoryService.
type MockRepositoryServiceMockRecorder struct {
	mock *MockRepositoryService
}

// NewMockRepositoryService creates a new mock instance.
func NewMockRepositoryService(ctrl *gomock.Controller) *MockRepositoryService {
	mock := &MockRepositoryService{ctrl: ctrl}
	mock.recorder = &MockRepositoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryService) EXPECT() *MockRepositoryServiceMockRecorder {
	return m.recorder
}

// CreateRepository mocks base method.
func (m *MockRepositoryService) CreateRepository(ctx context.Context, provider *db.Provider, projectID uuid.UUID, repoName, repoOwner string) (*v1.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRepository", ctx, provider, projectID, repoName, repoOwner)
	ret0, _ := ret[0].(*v1.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRepository indicates an expected call of CreateRepository.
func (mr *MockRepositoryServiceMockRecorder) CreateRepository(ctx, provider, projectID, repoName, repoOwner any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRepository", reflect.TypeOf((*MockRepositoryService)(nil).CreateRepository), ctx, provider, projectID, repoName, repoOwner)
}

// DeleteByID mocks base method.
func (m *MockRepositoryService) DeleteByID(ctx context.Context, repoID, projectID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, repoID, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockRepositoryServiceMockRecorder) DeleteByID(ctx, repoID, projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockRepositoryService)(nil).DeleteByID), ctx, repoID, projectID)
}

// DeleteByName mocks base method.
func (m *MockRepositoryService) DeleteByName(ctx context.Context, repoOwner, repoName string, projectID uuid.UUID, providerName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByName", ctx, repoOwner, repoName, projectID, providerName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByName indicates an expected call of DeleteByName.
func (mr *MockRepositoryServiceMockRecorder) DeleteByName(ctx, repoOwner, repoName, projectID, providerName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByName", reflect.TypeOf((*MockRepositoryService)(nil).DeleteByName), ctx, repoOwner, repoName, projectID, providerName)
}

// GetRepositoryById mocks base method.
func (m *MockRepositoryService) GetRepositoryById(ctx context.Context, repositoryID, projectID uuid.UUID) (db.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryById", ctx, repositoryID, projectID)
	ret0, _ := ret[0].(db.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryById indicates an expected call of GetRepositoryById.
func (mr *MockRepositoryServiceMockRecorder) GetRepositoryById(ctx, repositoryID, projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryById", reflect.TypeOf((*MockRepositoryService)(nil).GetRepositoryById), ctx, repositoryID, projectID)
}

// GetRepositoryByName mocks base method.
func (m *MockRepositoryService) GetRepositoryByName(ctx context.Context, repoOwner, repoName string, projectID uuid.UUID, providerName string) (db.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByName", ctx, repoOwner, repoName, projectID, providerName)
	ret0, _ := ret[0].(db.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByName indicates an expected call of GetRepositoryByName.
func (mr *MockRepositoryServiceMockRecorder) GetRepositoryByName(ctx, repoOwner, repoName, projectID, providerName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryByName", reflect.TypeOf((*MockRepositoryService)(nil).GetRepositoryByName), ctx, repoOwner, repoName, projectID, providerName)
}
