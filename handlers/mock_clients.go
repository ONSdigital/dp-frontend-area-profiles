package handlers

import (
	"github.com/ONSdigital/dp-renderer/model"
	"github.com/golang/mock/gomock"
	"io"
	"reflect"
)

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
