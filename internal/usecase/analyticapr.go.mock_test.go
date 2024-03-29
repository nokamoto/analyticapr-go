// Code generated by MockGen. DO NOT EDIT.
// Source: analyticapr.go
//
// Generated by this command:
//
//	mockgen -source=analyticapr.go -destination=analyticapr.go.mock_test.go -package=usecase
//

// Package usecase is a generated GoMock package.
package usecase

import (
	reflect "reflect"

	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
	gomock "go.uber.org/mock/gomock"
)

// Mockgh is a mock of gh interface.
type Mockgh struct {
	ctrl     *gomock.Controller
	recorder *MockghMockRecorder
}

// MockghMockRecorder is the mock recorder for Mockgh.
type MockghMockRecorder struct {
	mock *Mockgh
}

// NewMockgh creates a new mock instance.
func NewMockgh(ctrl *gomock.Controller) *Mockgh {
	mock := &Mockgh{ctrl: ctrl}
	mock.recorder = &MockghMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockgh) EXPECT() *MockghMockRecorder {
	return m.recorder
}

// ListPulls mocks base method.
func (m *Mockgh) ListPulls(arg0 *v1.Repository, arg1 string) ([]*v1.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPulls", arg0, arg1)
	ret0, _ := ret[0].([]*v1.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPulls indicates an expected call of ListPulls.
func (mr *MockghMockRecorder) ListPulls(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPulls", reflect.TypeOf((*Mockgh)(nil).ListPulls), arg0, arg1)
}
