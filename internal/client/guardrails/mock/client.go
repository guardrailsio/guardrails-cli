// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/guardrailsio/guardrails-cli/internal/client/guardrails (interfaces: GuardRailsClient)

// Package mockguardrailsclient is a generated GoMock package.
package mockguardrailsclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
)

// MockGuardRailsClient is a mock of GuardRailsClient interface.
type MockGuardRailsClient struct {
	ctrl     *gomock.Controller
	recorder *MockGuardRailsClientMockRecorder
}

// MockGuardRailsClientMockRecorder is the mock recorder for MockGuardRailsClient.
type MockGuardRailsClientMockRecorder struct {
	mock *MockGuardRailsClient
}

// NewMockGuardRailsClient creates a new mock instance.
func NewMockGuardRailsClient(ctrl *gomock.Controller) *MockGuardRailsClient {
	mock := &MockGuardRailsClient{ctrl: ctrl}
	mock.recorder = &MockGuardRailsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGuardRailsClient) EXPECT() *MockGuardRailsClientMockRecorder {
	return m.recorder
}

// CreateUploadURL mocks base method.
func (m *MockGuardRailsClient) CreateUploadURL(arg0 context.Context, arg1 *guardrailsclient.CreateUploadURLReq) (*guardrailsclient.CreateUploadURLResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUploadURL", arg0, arg1)
	ret0, _ := ret[0].(*guardrailsclient.CreateUploadURLResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUploadURL indicates an expected call of CreateUploadURL.
func (mr *MockGuardRailsClientMockRecorder) CreateUploadURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUploadURL", reflect.TypeOf((*MockGuardRailsClient)(nil).CreateUploadURL), arg0, arg1)
}

// GetScanData mocks base method.
func (m *MockGuardRailsClient) GetScanData(arg0 context.Context, arg1 *guardrailsclient.GetScanDataReq) (*guardrailsclient.GetScanDataResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScanData", arg0, arg1)
	ret0, _ := ret[0].(*guardrailsclient.GetScanDataResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScanData indicates an expected call of GetScanData.
func (mr *MockGuardRailsClientMockRecorder) GetScanData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScanData", reflect.TypeOf((*MockGuardRailsClient)(nil).GetScanData), arg0, arg1)
}

// TriggerScan mocks base method.
func (m *MockGuardRailsClient) TriggerScan(arg0 context.Context, arg1 *guardrailsclient.TriggerScanReq) (*guardrailsclient.TriggerScanResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerScan", arg0, arg1)
	ret0, _ := ret[0].(*guardrailsclient.TriggerScanResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerScan indicates an expected call of TriggerScan.
func (mr *MockGuardRailsClientMockRecorder) TriggerScan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerScan", reflect.TypeOf((*MockGuardRailsClient)(nil).TriggerScan), arg0, arg1)
}

// UploadProject mocks base method.
func (m *MockGuardRailsClient) UploadProject(arg0 context.Context, arg1 *guardrailsclient.UploadProjectReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadProject indicates an expected call of UploadProject.
func (mr *MockGuardRailsClientMockRecorder) UploadProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadProject", reflect.TypeOf((*MockGuardRailsClient)(nil).UploadProject), arg0, arg1)
}
