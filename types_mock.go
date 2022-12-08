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

// CreateCanaryRelease mocks base method.
func (m *MockInterface) CreateCanaryRelease(ctx context.Context, cr *CanaryRelease, opts *ResourceCreateOptions) (*CanaryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCanaryRelease", ctx, cr, opts)
	ret0, _ := ret[0].(*CanaryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCanaryRelease indicates an expected call of CreateCanaryRelease.
func (mr *MockInterfaceMockRecorder) CreateCanaryRelease(ctx, cr, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCanaryRelease", reflect.TypeOf((*MockInterface)(nil).CreateCanaryRelease), ctx, cr, opts)
}

// CreateConsumer mocks base method.
func (m *MockInterface) CreateConsumer(ctx context.Context, consumer *Consumer, opts *ResourceCreateOptions) (*Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConsumer", ctx, consumer, opts)
	ret0, _ := ret[0].(*Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConsumer indicates an expected call of CreateConsumer.
func (mr *MockInterfaceMockRecorder) CreateConsumer(ctx, consumer, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConsumer", reflect.TypeOf((*MockInterface)(nil).CreateConsumer), ctx, consumer, opts)
}

// CreateLogCollection mocks base method.
func (m *MockInterface) CreateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceCreateOptions) (*LogCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLogCollection", ctx, lc, opts)
	ret0, _ := ret[0].(*LogCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLogCollection indicates an expected call of CreateLogCollection.
func (mr *MockInterfaceMockRecorder) CreateLogCollection(ctx, lc, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogCollection", reflect.TypeOf((*MockInterface)(nil).CreateLogCollection), ctx, lc, opts)
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

// DeleteCanaryRelease mocks base method.
func (m *MockInterface) DeleteCanaryRelease(ctx context.Context, crID ID, opts *ResourceDeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCanaryRelease", ctx, crID, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCanaryRelease indicates an expected call of DeleteCanaryRelease.
func (mr *MockInterfaceMockRecorder) DeleteCanaryRelease(ctx, crID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCanaryRelease", reflect.TypeOf((*MockInterface)(nil).DeleteCanaryRelease), ctx, crID, opts)
}

// DeleteConsumer mocks base method.
func (m *MockInterface) DeleteConsumer(ctx context.Context, consumerID ID, opts *ResourceDeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConsumer", ctx, consumerID, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConsumer indicates an expected call of DeleteConsumer.
func (mr *MockInterfaceMockRecorder) DeleteConsumer(ctx, consumerID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConsumer", reflect.TypeOf((*MockInterface)(nil).DeleteConsumer), ctx, consumerID, opts)
}

// DeleteLogCollection mocks base method.
func (m *MockInterface) DeleteLogCollection(ctx context.Context, lcID ID, opts *ResourceDeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLogCollection", ctx, lcID, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLogCollection indicates an expected call of DeleteLogCollection.
func (mr *MockInterfaceMockRecorder) DeleteLogCollection(ctx, lcID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLogCollection", reflect.TypeOf((*MockInterface)(nil).DeleteLogCollection), ctx, lcID, opts)
}

// FinishCanaryRelease mocks base method.
func (m *MockInterface) FinishCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishCanaryRelease", ctx, crID, opts)
	ret0, _ := ret[0].(*CanaryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinishCanaryRelease indicates an expected call of FinishCanaryRelease.
func (mr *MockInterfaceMockRecorder) FinishCanaryRelease(ctx, crID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishCanaryRelease", reflect.TypeOf((*MockInterface)(nil).FinishCanaryRelease), ctx, crID, opts)
}

// GenerateGatewaySideCertificate mocks base method.
func (m *MockInterface) GenerateGatewaySideCertificate(ctx context.Context, opts *ResourceCreateOptions) (*TLSBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateGatewaySideCertificate", ctx, opts)
	ret0, _ := ret[0].(*TLSBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateGatewaySideCertificate indicates an expected call of GenerateGatewaySideCertificate.
func (mr *MockInterfaceMockRecorder) GenerateGatewaySideCertificate(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateGatewaySideCertificate", reflect.TypeOf((*MockInterface)(nil).GenerateGatewaySideCertificate), ctx, opts)
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

// GetCanaryRelease mocks base method.
func (m *MockInterface) GetCanaryRelease(ctx context.Context, crID ID, opts *ResourceGetOptions) (*CanaryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCanaryRelease", ctx, crID, opts)
	ret0, _ := ret[0].(*CanaryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCanaryRelease indicates an expected call of GetCanaryRelease.
func (mr *MockInterfaceMockRecorder) GetCanaryRelease(ctx, crID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanaryRelease", reflect.TypeOf((*MockInterface)(nil).GetCanaryRelease), ctx, crID, opts)
}

// GetConsumer mocks base method.
func (m *MockInterface) GetConsumer(ctx context.Context, consumerID ID, opts *ResourceGetOptions) (*Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumer", ctx, consumerID, opts)
	ret0, _ := ret[0].(*Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumer indicates an expected call of GetConsumer.
func (mr *MockInterfaceMockRecorder) GetConsumer(ctx, consumerID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumer", reflect.TypeOf((*MockInterface)(nil).GetConsumer), ctx, consumerID, opts)
}

// GetLogCollection mocks base method.
func (m *MockInterface) GetLogCollection(ctx context.Context, lcID ID, opts *ResourceGetOptions) (*LogCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogCollection", ctx, lcID, opts)
	ret0, _ := ret[0].(*LogCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogCollection indicates an expected call of GetLogCollection.
func (mr *MockInterfaceMockRecorder) GetLogCollection(ctx, lcID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogCollection", reflect.TypeOf((*MockInterface)(nil).GetLogCollection), ctx, lcID, opts)
}

// GetOrganization mocks base method.
func (m *MockInterface) GetOrganization(ctx context.Context, orgID ID, opts *ResourceGetOptions) (*Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganization", ctx, orgID, opts)
	ret0, _ := ret[0].(*Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganization indicates an expected call of GetOrganization.
func (mr *MockInterfaceMockRecorder) GetOrganization(ctx, orgID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganization", reflect.TypeOf((*MockInterface)(nil).GetOrganization), ctx, orgID, opts)
}

// InviteMember mocks base method.
func (m *MockInterface) InviteMember(ctx context.Context, email string, role *Role, opts *ResourceCreateOptions) (*Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InviteMember", ctx, email, role, opts)
	ret0, _ := ret[0].(*Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InviteMember indicates an expected call of InviteMember.
func (mr *MockInterfaceMockRecorder) InviteMember(ctx, email, role, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteMember", reflect.TypeOf((*MockInterface)(nil).InviteMember), ctx, email, role, opts)
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

// ListAllGatewayInstances mocks base method.
func (m *MockInterface) ListAllGatewayInstances(ctx context.Context, cpID ID, opts *ResourceListOptions) ([]GatewayInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllGatewayInstances", ctx, cpID, opts)
	ret0, _ := ret[0].([]GatewayInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllGatewayInstances indicates an expected call of ListAllGatewayInstances.
func (mr *MockInterfaceMockRecorder) ListAllGatewayInstances(ctx, cpID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllGatewayInstances", reflect.TypeOf((*MockInterface)(nil).ListAllGatewayInstances), ctx, cpID, opts)
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

// ListCanaryReleases mocks base method.
func (m *MockInterface) ListCanaryReleases(ctx context.Context, opts *ResourceListOptions) (CanaryReleaseListIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCanaryReleases", ctx, opts)
	ret0, _ := ret[0].(CanaryReleaseListIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCanaryReleases indicates an expected call of ListCanaryReleases.
func (mr *MockInterfaceMockRecorder) ListCanaryReleases(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCanaryReleases", reflect.TypeOf((*MockInterface)(nil).ListCanaryReleases), ctx, opts)
}

// ListConsumers mocks base method.
func (m *MockInterface) ListConsumers(ctx context.Context, opts *ResourceListOptions) (ConsumerListIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConsumers", ctx, opts)
	ret0, _ := ret[0].(ConsumerListIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConsumers indicates an expected call of ListConsumers.
func (mr *MockInterfaceMockRecorder) ListConsumers(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConsumers", reflect.TypeOf((*MockInterface)(nil).ListConsumers), ctx, opts)
}

// ListControlPlanes mocks base method.
func (m *MockInterface) ListControlPlanes(ctx context.Context, opts *ResourceListOptions) (ControlPlaneListIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListControlPlanes", ctx, opts)
	ret0, _ := ret[0].(ControlPlaneListIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListControlPlanes indicates an expected call of ListControlPlanes.
func (mr *MockInterfaceMockRecorder) ListControlPlanes(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListControlPlanes", reflect.TypeOf((*MockInterface)(nil).ListControlPlanes), ctx, opts)
}

// ListLogCollections mocks base method.
func (m *MockInterface) ListLogCollections(ctx context.Context, opts *ResourceListOptions) (LogCollectionIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLogCollections", ctx, opts)
	ret0, _ := ret[0].(LogCollectionIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLogCollections indicates an expected call of ListLogCollections.
func (mr *MockInterfaceMockRecorder) ListLogCollections(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLogCollections", reflect.TypeOf((*MockInterface)(nil).ListLogCollections), ctx, opts)
}

// ListMembers mocks base method.
func (m *MockInterface) ListMembers(ctx context.Context, opts *ResourceListOptions) (MemberListIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMembers", ctx, opts)
	ret0, _ := ret[0].(MemberListIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockInterfaceMockRecorder) ListMembers(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockInterface)(nil).ListMembers), ctx, opts)
}

// ListRegions mocks base method.
func (m *MockInterface) ListRegions(ctx context.Context, opts *ResourceListOptions) (RegionListIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRegions", ctx, opts)
	ret0, _ := ret[0].(RegionListIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegions indicates an expected call of ListRegions.
func (mr *MockInterfaceMockRecorder) ListRegions(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegions", reflect.TypeOf((*MockInterface)(nil).ListRegions), ctx, opts)
}

// ListRoles mocks base method.
func (m *MockInterface) ListRoles(ctx context.Context, opts *ResourceListOptions) (RoleListIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoles", ctx, opts)
	ret0, _ := ret[0].(RoleListIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoles indicates an expected call of ListRoles.
func (mr *MockInterfaceMockRecorder) ListRoles(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockInterface)(nil).ListRoles), ctx, opts)
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

// PauseCanaryRelease mocks base method.
func (m *MockInterface) PauseCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseCanaryRelease", ctx, crID, opts)
	ret0, _ := ret[0].(*CanaryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PauseCanaryRelease indicates an expected call of PauseCanaryRelease.
func (mr *MockInterfaceMockRecorder) PauseCanaryRelease(ctx, crID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseCanaryRelease", reflect.TypeOf((*MockInterface)(nil).PauseCanaryRelease), ctx, crID, opts)
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

// ReInviteMember mocks base method.
func (m *MockInterface) ReInviteMember(ctx context.Context, memberID ID, opts *ResourceUpdateOptions) (*Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReInviteMember", ctx, memberID, opts)
	ret0, _ := ret[0].(*Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReInviteMember indicates an expected call of ReInviteMember.
func (mr *MockInterfaceMockRecorder) ReInviteMember(ctx, memberID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReInviteMember", reflect.TypeOf((*MockInterface)(nil).ReInviteMember), ctx, memberID, opts)
}

// StartCanaryRelease mocks base method.
func (m *MockInterface) StartCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCanaryRelease", ctx, crID, opts)
	ret0, _ := ret[0].(*CanaryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartCanaryRelease indicates an expected call of StartCanaryRelease.
func (mr *MockInterfaceMockRecorder) StartCanaryRelease(ctx, crID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCanaryRelease", reflect.TypeOf((*MockInterface)(nil).StartCanaryRelease), ctx, crID, opts)
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

// UpdateCanaryRelease mocks base method.
func (m *MockInterface) UpdateCanaryRelease(ctx context.Context, cr *CanaryRelease, opts *ResourceUpdateOptions) (*CanaryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCanaryRelease", ctx, cr, opts)
	ret0, _ := ret[0].(*CanaryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCanaryRelease indicates an expected call of UpdateCanaryRelease.
func (mr *MockInterfaceMockRecorder) UpdateCanaryRelease(ctx, cr, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCanaryRelease", reflect.TypeOf((*MockInterface)(nil).UpdateCanaryRelease), ctx, cr, opts)
}

// UpdateConsumer mocks base method.
func (m *MockInterface) UpdateConsumer(ctx context.Context, consumer *Consumer, opts *ResourceUpdateOptions) (*Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConsumer", ctx, consumer, opts)
	ret0, _ := ret[0].(*Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConsumer indicates an expected call of UpdateConsumer.
func (mr *MockInterfaceMockRecorder) UpdateConsumer(ctx, consumer, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConsumer", reflect.TypeOf((*MockInterface)(nil).UpdateConsumer), ctx, consumer, opts)
}

// UpdateLogCollection mocks base method.
func (m *MockInterface) UpdateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceUpdateOptions) (*LogCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLogCollection", ctx, lc, opts)
	ret0, _ := ret[0].(*LogCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLogCollection indicates an expected call of UpdateLogCollection.
func (mr *MockInterfaceMockRecorder) UpdateLogCollection(ctx, lc, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLogCollection", reflect.TypeOf((*MockInterface)(nil).UpdateLogCollection), ctx, lc, opts)
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
