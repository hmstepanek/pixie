// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_idmanager is a generated GoMock package.
package mock_idmanager

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	idmanager "px.dev/pixie/src/cloud/profile/controller/idmanager"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CreateInviteLink mocks base method.
func (m *MockManager) CreateInviteLink(ctx context.Context, req *idmanager.CreateInviteLinkRequest) (*idmanager.CreateInviteLinkResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInviteLink", ctx, req)
	ret0, _ := ret[0].(*idmanager.CreateInviteLinkResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInviteLink indicates an expected call of CreateInviteLink.
func (mr *MockManagerMockRecorder) CreateInviteLink(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInviteLink", reflect.TypeOf((*MockManager)(nil).CreateInviteLink), ctx, req)
}
