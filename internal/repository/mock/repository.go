// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/guardrailsio/guardrails-cli/internal/repository (interfaces: Repository)

// Package mockrepository is a generated GoMock package.
package mockrepository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repository "github.com/guardrailsio/guardrails-cli/internal/repository"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetMetadataFromRemoteURL mocks base method.
func (m *MockRepository) GetMetadataFromRemoteURL() (*repository.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadataFromRemoteURL")
	ret0, _ := ret[0].(*repository.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadataFromRemoteURL indicates an expected call of GetMetadataFromRemoteURL.
func (mr *MockRepositoryMockRecorder) GetMetadataFromRemoteURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataFromRemoteURL", reflect.TypeOf((*MockRepository)(nil).GetMetadataFromRemoteURL))
}

// ListFiles mocks base method.
func (m *MockRepository) ListFiles() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockRepositoryMockRecorder) ListFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockRepository)(nil).ListFiles))
}
