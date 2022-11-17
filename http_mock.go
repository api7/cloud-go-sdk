// Code generated by MockGen. DO NOT EDIT.
// Source: ./http.go

// Package cloud is a generated GoMock package.
package cloud

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockhttpClient is a mock of httpClient interface.
type MockhttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockhttpClientMockRecorder
}

// MockhttpClientMockRecorder is the mock recorder for MockhttpClient.
type MockhttpClientMockRecorder struct {
	mock *MockhttpClient
}

// NewMockhttpClient creates a new mock instance.
func NewMockhttpClient(ctrl *gomock.Controller) *MockhttpClient {
	mock := &MockhttpClient{ctrl: ctrl}
	mock.recorder = &MockhttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockhttpClient) EXPECT() *MockhttpClientMockRecorder {
	return m.recorder
}

// sendDeleteRequest mocks base method.
func (m *MockhttpClient) sendDeleteRequest(ctx context.Context, path, query string, payloadDecodeFunc payloadDecodeFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "sendDeleteRequest", ctx, path, query, payloadDecodeFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// sendDeleteRequest indicates an expected call of sendDeleteRequest.
func (mr *MockhttpClientMockRecorder) sendDeleteRequest(ctx, path, query, payloadDecodeFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sendDeleteRequest", reflect.TypeOf((*MockhttpClient)(nil).sendDeleteRequest), ctx, path, query, payloadDecodeFunc)
}

// sendGetRequest mocks base method.
func (m *MockhttpClient) sendGetRequest(ctx context.Context, path, query string, payloadDecodeFunc payloadDecodeFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "sendGetRequest", ctx, path, query, payloadDecodeFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// sendGetRequest indicates an expected call of sendGetRequest.
func (mr *MockhttpClientMockRecorder) sendGetRequest(ctx, path, query, payloadDecodeFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sendGetRequest", reflect.TypeOf((*MockhttpClient)(nil).sendGetRequest), ctx, path, query, payloadDecodeFunc)
}

// sendPatchRequest mocks base method.
func (m *MockhttpClient) sendPatchRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "sendPatchRequest", ctx, path, query, body, payloadDecodeFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// sendPatchRequest indicates an expected call of sendPatchRequest.
func (mr *MockhttpClientMockRecorder) sendPatchRequest(ctx, path, query, body, payloadDecodeFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sendPatchRequest", reflect.TypeOf((*MockhttpClient)(nil).sendPatchRequest), ctx, path, query, body, payloadDecodeFunc)
}

// sendPostRequest mocks base method.
func (m *MockhttpClient) sendPostRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "sendPostRequest", ctx, path, query, body, payloadDecodeFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// sendPostRequest indicates an expected call of sendPostRequest.
func (mr *MockhttpClientMockRecorder) sendPostRequest(ctx, path, query, body, payloadDecodeFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sendPostRequest", reflect.TypeOf((*MockhttpClient)(nil).sendPostRequest), ctx, path, query, body, payloadDecodeFunc)
}

// sendPutRequest mocks base method.
func (m *MockhttpClient) sendPutRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "sendPutRequest", ctx, path, query, body, payloadDecodeFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// sendPutRequest indicates an expected call of sendPutRequest.
func (mr *MockhttpClientMockRecorder) sendPutRequest(ctx, path, query, body, payloadDecodeFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sendPutRequest", reflect.TypeOf((*MockhttpClient)(nil).sendPutRequest), ctx, path, query, body, payloadDecodeFunc)
}

// sendRequest mocks base method.
func (m *MockhttpClient) sendRequest(req *http.Request, payloadDecodeFunc payloadDecodeFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "sendRequest", req, payloadDecodeFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// sendRequest indicates an expected call of sendRequest.
func (mr *MockhttpClientMockRecorder) sendRequest(req, payloadDecodeFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sendRequest", reflect.TypeOf((*MockhttpClient)(nil).sendRequest), req, payloadDecodeFunc)
}
