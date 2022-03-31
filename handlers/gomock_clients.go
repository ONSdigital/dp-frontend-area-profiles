// Code generated by MockGen. DO NOT EDIT.
// Source: handlers/clients.go

// Package handlers is a generated GoMock package.
package handlers

import (
	context "context"
	io "io"
	reflect "reflect"

	areas "github.com/ONSdigital/dp-api-clients-go/v2/areas"
	healthcheck "github.com/ONSdigital/dp-healthcheck/healthcheck"
	model "github.com/ONSdigital/dp-renderer/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAreaApiClient is a mock of AreaApiClient interface.
type MockAreaApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockAreaApiClientMockRecorder
}

// MockAreaApiClientMockRecorder is the mock recorder for MockAreaApiClient.
type MockAreaApiClientMockRecorder struct {
	mock *MockAreaApiClient
}

// NewMockAreaApiClient creates a new mock instance.
func NewMockAreaApiClient(ctrl *gomock.Controller) *MockAreaApiClient {
	mock := &MockAreaApiClient{ctrl: ctrl}
	mock.recorder = &MockAreaApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAreaApiClient) EXPECT() *MockAreaApiClientMockRecorder {
	return m.recorder
}

// Checker mocks base method.
func (m *MockAreaApiClient) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checker", ctx, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checker indicates an expected call of Checker.
func (mr *MockAreaApiClientMockRecorder) Checker(ctx, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checker", reflect.TypeOf((*MockAreaApiClient)(nil).Checker), ctx, check)
}

// GetArea mocks base method.
func (m *MockAreaApiClient) GetArea(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang string) (areas.AreaDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArea", ctx, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang)
	ret0, _ := ret[0].(areas.AreaDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArea indicates an expected call of GetArea.
func (mr *MockAreaApiClientMockRecorder) GetArea(ctx, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArea", reflect.TypeOf((*MockAreaApiClient)(nil).GetArea), ctx, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang)
}

// GetRelations mocks base method.
func (m *MockAreaApiClient) GetRelations(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang string) ([]areas.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelations", ctx, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang)
	ret0, _ := ret[0].([]areas.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelations indicates an expected call of GetRelations.
func (mr *MockAreaApiClientMockRecorder) GetRelations(ctx, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelations", reflect.TypeOf((*MockAreaApiClient)(nil).GetRelations), ctx, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang)
}

// MockClientError is a mock of ClientError interface.
type MockClientError struct {
	ctrl     *gomock.Controller
	recorder *MockClientErrorMockRecorder
}

// MockClientErrorMockRecorder is the mock recorder for MockClientError.
type MockClientErrorMockRecorder struct {
	mock *MockClientError
}

// NewMockClientError creates a new mock instance.
func NewMockClientError(ctrl *gomock.Controller) *MockClientError {
	mock := &MockClientError{ctrl: ctrl}
	mock.recorder = &MockClientErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientError) EXPECT() *MockClientErrorMockRecorder {
	return m.recorder
}

// Code mocks base method.
func (m *MockClientError) Code() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Code")
	ret0, _ := ret[0].(int)
	return ret0
}

// Code indicates an expected call of Code.
func (mr *MockClientErrorMockRecorder) Code() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Code", reflect.TypeOf((*MockClientError)(nil).Code))
}

// Error mocks base method.
func (m *MockClientError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockClientErrorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockClientError)(nil).Error))
}

// MockRenderClient is a mock of RenderClient interface.
type MockRenderClient struct {
	ctrl     *gomock.Controller
	recorder *MockRenderClientMockRecorder
}

// MockRenderClientMockRecorder is the mock recorder for MockRenderClient.
type MockRenderClientMockRecorder struct {
	mock *MockRenderClient
}

// NewMockRenderClient creates a new mock instance.
func NewMockRenderClient(ctrl *gomock.Controller) *MockRenderClient {
	mock := &MockRenderClient{ctrl: ctrl}
	mock.recorder = &MockRenderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRenderClient) EXPECT() *MockRenderClientMockRecorder {
	return m.recorder
}

// BuildPage mocks base method.
func (m *MockRenderClient) BuildPage(w io.Writer, pageModel interface{}, templateName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BuildPage", w, pageModel, templateName)
}

// BuildPage indicates an expected call of BuildPage.
func (mr *MockRenderClientMockRecorder) BuildPage(w, pageModel, templateName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildPage", reflect.TypeOf((*MockRenderClient)(nil).BuildPage), w, pageModel, templateName)
}

// NewBasePageModel mocks base method.
func (m *MockRenderClient) NewBasePageModel() model.Page {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBasePageModel")
	ret0, _ := ret[0].(model.Page)
	return ret0
}

// NewBasePageModel indicates an expected call of NewBasePageModel.
func (mr *MockRenderClientMockRecorder) NewBasePageModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBasePageModel", reflect.TypeOf((*MockRenderClient)(nil).NewBasePageModel))
}

// MockRendererClient is a mock of RendererClient interface.
type MockRendererClient struct {
	ctrl     *gomock.Controller
	recorder *MockRendererClientMockRecorder
}

// MockRendererClientMockRecorder is the mock recorder for MockRendererClient.
type MockRendererClientMockRecorder struct {
	mock *MockRendererClient
}

// NewMockRendererClient creates a new mock instance.
func NewMockRendererClient(ctrl *gomock.Controller) *MockRendererClient {
	mock := &MockRendererClient{ctrl: ctrl}
	mock.recorder = &MockRendererClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRendererClient) EXPECT() *MockRendererClientMockRecorder {
	return m.recorder
}

// Checker mocks base method.
func (m *MockRendererClient) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checker", ctx, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checker indicates an expected call of Checker.
func (mr *MockRendererClientMockRecorder) Checker(ctx, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checker", reflect.TypeOf((*MockRendererClient)(nil).Checker), ctx, check)
}

// Do mocks base method.
func (m *MockRendererClient) Do(arg0 string, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockRendererClientMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockRendererClient)(nil).Do), arg0, arg1)
}
