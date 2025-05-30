// Code generated by MockGen. DO NOT EDIT.
// Source: redis_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "github.com/ayoadeoye1/go-clean-archi/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSessRepository is a mock of SessRepository interface
type MockSessRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessRepositoryMockRecorder
}

// MockSessRepositoryMockRecorder is the mock recorder for MockSessRepository
type MockSessRepositoryMockRecorder struct {
	mock *MockSessRepository
}

// NewMockSessRepository creates a new mock instance
func NewMockSessRepository(ctrl *gomock.Controller) *MockSessRepository {
	mock := &MockSessRepository{ctrl: ctrl}
	mock.recorder = &MockSessRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessRepository) EXPECT() *MockSessRepositoryMockRecorder {
	return m.recorder
}

// CreateSession mocks base method
func (m *MockSessRepository) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, session, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockSessRepositoryMockRecorder) CreateSession(ctx, session, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessRepository)(nil).CreateSession), ctx, session, expire)
}

// GetSessionByID mocks base method
func (m *MockSessRepository) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByID", ctx, sessionID)
	ret0, _ := ret[0].(*models.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByID indicates an expected call of GetSessionByID
func (mr *MockSessRepositoryMockRecorder) GetSessionByID(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByID", reflect.TypeOf((*MockSessRepository)(nil).GetSessionByID), ctx, sessionID)
}

// DeleteByID mocks base method
func (m *MockSessRepository) DeleteByID(ctx context.Context, sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockSessRepositoryMockRecorder) DeleteByID(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockSessRepository)(nil).DeleteByID), ctx, sessionID)
}
