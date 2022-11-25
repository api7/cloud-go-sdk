// Code generated by MockGen. DO NOT EDIT.
// Source: ./types.go

package cloud

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateAPI mocks base method.
func (m *MockInterface) CreateAPI(ctx context.Context, api *API, opts *ResourceCreateOptions) (*API, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAPI", ctx, api, opts)
	ret0, _ := ret[0].(*API)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAPI indicates an expected call of CreateAPI.
func (mr *MockInterfaceMockRecorder) CreateAPI(ctx, api, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAPI", reflect.TypeOf((*MockInterface)(nil).CreateAPI), ctx, api, opts)
}

// CreateAccessToken mocks base method.
func (m *MockInterface) CreateAccessToken(ctx context.Context, token *AccessToken) (*AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", ctx, token)
	ret0, _ := ret[0].(*AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockInterfaceMockRecorder) CreateAccessToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockInterface)(nil).CreateAccessToken), ctx, token)
}

// CreateApplication mocks base method.
func (m *MockInterface) CreateApplication(ctx context.Context, app *Application, opts *ResourceCreateOptions) (*Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", ctx, app, opts)
	ret0, _ := ret[0].(*Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockInterfaceMockRecorder) CreateApplication(ctx, app, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockInterface)(nil).CreateApplication), ctx, app, opts)
}

// DeleteAPI mocks base method.
func (m *MockInterface) DeleteAPI(ctx context.Context, apiID ID, opts *ResourceDeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAPI", ctx, apiID, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAPI indicates an expected call of DeleteAPI.
func (mr *MockInterfaceMockRecorder) DeleteAPI(ctx, apiID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAPI", reflect.TypeOf((*MockInterface)(nil).DeleteAPI), ctx, apiID, opts)
}

// DeleteAccessToken mocks base method.
func (m *MockInterface) DeleteAccessToken(ctx context.Context, token *AccessToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessToken", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessToken indicates an expected call of DeleteAccessToken.
func (mr *MockInterfaceMockRecorder) DeleteAccessToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessToken", reflect.TypeOf((*MockInterface)(nil).DeleteAccessToken), ctx, token)
}

// DeleteApplication mocks base method.
func (m *MockInterface) DeleteApplication(ctx context.Context, appID ID, opts *ResourceDeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", ctx, appID, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockInterfaceMockRecorder) DeleteApplication(ctx, appID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockInterface)(nil).DeleteApplication), ctx, appID, opts)
}

// GetAPI mocks base method.
func (m *MockInterface) GetAPI(ctx context.Context, apiID ID, opts *ResourceGetOptions) (*API, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPI", ctx, apiID, opts)
	ret0, _ := ret[0].(*API)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPI indicates an expected call of GetAPI.
func (mr *MockInterfaceMockRecorder) GetAPI(ctx, apiID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPI", reflect.TypeOf((*MockInterface)(nil).GetAPI), ctx, apiID, opts)
}

// GetApplication mocks base method.
func (m *MockInterface) GetApplication(ctx context.Context, appID ID, opts *ResourceGetOptions) (*Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", ctx, appID, opts)
	ret0, _ := ret[0].(*Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockInterfaceMockRecorder) GetApplication(ctx, appID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockInterface)(nil).GetApplication), ctx, appID, opts)
}

// ListAPIs mocks base method.
func (m *MockInterface) ListAPIs(ctx context.Context, opts *ResourceListOptions) (APIListIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAPIs", ctx, opts)
	ret0, _ := ret[0].(APIListIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAPIs indicates an expected call of ListAPIs.
func (mr *MockInterfaceMockRecorder) ListAPIs(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAPIs", reflect.TypeOf((*MockInterface)(nil).ListAPIs), ctx, opts)
}

// ListApplications mocks base method.
func (m *MockInterface) ListApplications(ctx context.Context, opts *ResourceListOptions) (ApplicationListIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", ctx, opts)
	ret0, _ := ret[0].(ApplicationListIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockInterfaceMockRecorder) ListApplications(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockInterface)(nil).ListApplications), ctx, opts)
}

// Me mocks base method.
func (m *MockInterface) Me(ctx context.Context) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Me", ctx)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Me indicates an expected call of Me.
func (mr *MockInterfaceMockRecorder) Me(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Me", reflect.TypeOf((*MockInterface)(nil).Me), ctx)
}

// PublishAPI mocks base method.
func (m *MockInterface) PublishAPI(ctx context.Context, apiID ID, opts *ResourceUpdateOptions) (*API, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishAPI", ctx, apiID, opts)
	ret0, _ := ret[0].(*API)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishAPI indicates an expected call of PublishAPI.
func (mr *MockInterfaceMockRecorder) PublishAPI(ctx, apiID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAPI", reflect.TypeOf((*MockInterface)(nil).PublishAPI), ctx, apiID, opts)
}

// PublishApplication mocks base method.
func (m *MockInterface) PublishApplication(ctx context.Context, appID ID, opts *ResourceUpdateOptions) (*Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishApplication", ctx, appID, opts)
	ret0, _ := ret[0].(*Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishApplication indicates an expected call of PublishApplication.
func (mr *MockInterfaceMockRecorder) PublishApplication(ctx, appID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishApplication", reflect.TypeOf((*MockInterface)(nil).PublishApplication), ctx, appID, opts)
}

// TraceChan mocks base method.
func (m *MockInterface) TraceChan() <-chan *TraceSeries {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceChan")
	ret0, _ := ret[0].(<-chan *TraceSeries)
	return ret0
}

// TraceChan indicates an expected call of TraceChan.
func (mr *MockInterfaceMockRecorder) TraceChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceChan", reflect.TypeOf((*MockInterface)(nil).TraceChan))
}

// UnpublishAPI mocks base method.
func (m *MockInterface) UnpublishAPI(ctx context.Context, apiID ID, opts *ResourceUpdateOptions) (*API, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpublishAPI", ctx, apiID, opts)
	ret0, _ := ret[0].(*API)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpublishAPI indicates an expected call of UnpublishAPI.
func (mr *MockInterfaceMockRecorder) UnpublishAPI(ctx, apiID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishAPI", reflect.TypeOf((*MockInterface)(nil).UnpublishAPI), ctx, apiID, opts)
}

// UnpublishApplication mocks base method.
func (m *MockInterface) UnpublishApplication(ctx context.Context, appID ID, opts *ResourceUpdateOptions) (*Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpublishApplication", ctx, appID, opts)
	ret0, _ := ret[0].(*Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpublishApplication indicates an expected call of UnpublishApplication.
func (mr *MockInterfaceMockRecorder) UnpublishApplication(ctx, appID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishApplication", reflect.TypeOf((*MockInterface)(nil).UnpublishApplication), ctx, appID, opts)
}

// UpdateAPI mocks base method.
func (m *MockInterface) UpdateAPI(ctx context.Context, api *API, opts *ResourceUpdateOptions) (*API, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAPI", ctx, api, opts)
	ret0, _ := ret[0].(*API)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAPI indicates an expected call of UpdateAPI.
func (mr *MockInterfaceMockRecorder) UpdateAPI(ctx, api, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAPI", reflect.TypeOf((*MockInterface)(nil).UpdateAPI), ctx, api, opts)
}

// UpdateApplication mocks base method.
func (m *MockInterface) UpdateApplication(ctx context.Context, app *Application, opts *ResourceUpdateOptions) (*Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplication", ctx, app, opts)
	ret0, _ := ret[0].(*Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApplication indicates an expected call of UpdateApplication.
func (mr *MockInterfaceMockRecorder) UpdateApplication(ctx, app, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplication", reflect.TypeOf((*MockInterface)(nil).UpdateApplication), ctx, app, opts)
}

// sendSeries mocks base method.
func (m *MockInterface) sendSeries(series *TraceSeries) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "sendSeries", series)
}

// sendSeries indicates an expected call of sendSeries.
func (mr *MockInterfaceMockRecorder) sendSeries(series interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sendSeries", reflect.TypeOf((*MockInterface)(nil).sendSeries), series)
}
