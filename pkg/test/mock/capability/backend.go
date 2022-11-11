// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/apis/capability/v1/capability_grpc.pb.go

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/rancher/opni/pkg/apis/capability/v1"
	v10 "github.com/rancher/opni/pkg/apis/core/v1"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockBackendClient is a mock of BackendClient interface.
type MockBackendClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackendClientMockRecorder
}

// MockBackendClientMockRecorder is the mock recorder for MockBackendClient.
type MockBackendClientMockRecorder struct {
	mock *MockBackendClient
}

// NewMockBackendClient creates a new mock instance.
func NewMockBackendClient(ctrl *gomock.Controller) *MockBackendClient {
	mock := &MockBackendClient{ctrl: ctrl}
	mock.recorder = &MockBackendClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendClient) EXPECT() *MockBackendClientMockRecorder {
	return m.recorder
}

// CanInstall mocks base method.
func (m *MockBackendClient) CanInstall(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CanInstall", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanInstall indicates an expected call of CanInstall.
func (mr *MockBackendClientMockRecorder) CanInstall(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanInstall", reflect.TypeOf((*MockBackendClient)(nil).CanInstall), varargs...)
}

// CancelUninstall mocks base method.
func (m *MockBackendClient) CancelUninstall(ctx context.Context, in *v10.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelUninstall", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelUninstall indicates an expected call of CancelUninstall.
func (mr *MockBackendClientMockRecorder) CancelUninstall(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelUninstall", reflect.TypeOf((*MockBackendClient)(nil).CancelUninstall), varargs...)
}

// Info mocks base method.
func (m *MockBackendClient) Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Details, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(*v1.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockBackendClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBackendClient)(nil).Info), varargs...)
}

// Install mocks base method.
func (m *MockBackendClient) Install(ctx context.Context, in *v1.InstallRequest, opts ...grpc.CallOption) (*v1.InstallResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Install", varargs...)
	ret0, _ := ret[0].(*v1.InstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockBackendClientMockRecorder) Install(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockBackendClient)(nil).Install), varargs...)
}

// InstallerTemplate mocks base method.
func (m *MockBackendClient) InstallerTemplate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.InstallerTemplateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InstallerTemplate", varargs...)
	ret0, _ := ret[0].(*v1.InstallerTemplateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallerTemplate indicates an expected call of InstallerTemplate.
func (mr *MockBackendClientMockRecorder) InstallerTemplate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallerTemplate", reflect.TypeOf((*MockBackendClient)(nil).InstallerTemplate), varargs...)
}

// Status mocks base method.
func (m *MockBackendClient) Status(ctx context.Context, in *v1.StatusRequest, opts ...grpc.CallOption) (*v1.NodeCapabilityStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*v1.NodeCapabilityStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockBackendClientMockRecorder) Status(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockBackendClient)(nil).Status), varargs...)
}

// Uninstall mocks base method.
func (m *MockBackendClient) Uninstall(ctx context.Context, in *v1.UninstallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Uninstall", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockBackendClientMockRecorder) Uninstall(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockBackendClient)(nil).Uninstall), varargs...)
}

// UninstallStatus mocks base method.
func (m *MockBackendClient) UninstallStatus(ctx context.Context, in *v10.Reference, opts ...grpc.CallOption) (*v10.TaskStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UninstallStatus", varargs...)
	ret0, _ := ret[0].(*v10.TaskStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UninstallStatus indicates an expected call of UninstallStatus.
func (mr *MockBackendClientMockRecorder) UninstallStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallStatus", reflect.TypeOf((*MockBackendClient)(nil).UninstallStatus), varargs...)
}

// MockBackendServer is a mock of BackendServer interface.
type MockBackendServer struct {
	ctrl     *gomock.Controller
	recorder *MockBackendServerMockRecorder
}

// MockBackendServerMockRecorder is the mock recorder for MockBackendServer.
type MockBackendServerMockRecorder struct {
	mock *MockBackendServer
}

// NewMockBackendServer creates a new mock instance.
func NewMockBackendServer(ctrl *gomock.Controller) *MockBackendServer {
	mock := &MockBackendServer{ctrl: ctrl}
	mock.recorder = &MockBackendServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendServer) EXPECT() *MockBackendServerMockRecorder {
	return m.recorder
}

// CanInstall mocks base method.
func (m *MockBackendServer) CanInstall(arg0 context.Context, arg1 *emptypb.Empty) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanInstall", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanInstall indicates an expected call of CanInstall.
func (mr *MockBackendServerMockRecorder) CanInstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanInstall", reflect.TypeOf((*MockBackendServer)(nil).CanInstall), arg0, arg1)
}

// CancelUninstall mocks base method.
func (m *MockBackendServer) CancelUninstall(arg0 context.Context, arg1 *v10.Reference) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelUninstall", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelUninstall indicates an expected call of CancelUninstall.
func (mr *MockBackendServerMockRecorder) CancelUninstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelUninstall", reflect.TypeOf((*MockBackendServer)(nil).CancelUninstall), arg0, arg1)
}

// Info mocks base method.
func (m *MockBackendServer) Info(arg0 context.Context, arg1 *emptypb.Empty) (*v1.Details, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*v1.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockBackendServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBackendServer)(nil).Info), arg0, arg1)
}

// Install mocks base method.
func (m *MockBackendServer) Install(arg0 context.Context, arg1 *v1.InstallRequest) (*v1.InstallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1)
	ret0, _ := ret[0].(*v1.InstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockBackendServerMockRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockBackendServer)(nil).Install), arg0, arg1)
}

// InstallerTemplate mocks base method.
func (m *MockBackendServer) InstallerTemplate(arg0 context.Context, arg1 *emptypb.Empty) (*v1.InstallerTemplateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallerTemplate", arg0, arg1)
	ret0, _ := ret[0].(*v1.InstallerTemplateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallerTemplate indicates an expected call of InstallerTemplate.
func (mr *MockBackendServerMockRecorder) InstallerTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallerTemplate", reflect.TypeOf((*MockBackendServer)(nil).InstallerTemplate), arg0, arg1)
}

// Status mocks base method.
func (m *MockBackendServer) Status(arg0 context.Context, arg1 *v1.StatusRequest) (*v1.NodeCapabilityStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*v1.NodeCapabilityStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockBackendServerMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockBackendServer)(nil).Status), arg0, arg1)
}

// Uninstall mocks base method.
func (m *MockBackendServer) Uninstall(arg0 context.Context, arg1 *v1.UninstallRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockBackendServerMockRecorder) Uninstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockBackendServer)(nil).Uninstall), arg0, arg1)
}

// UninstallStatus mocks base method.
func (m *MockBackendServer) UninstallStatus(arg0 context.Context, arg1 *v10.Reference) (*v10.TaskStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallStatus", arg0, arg1)
	ret0, _ := ret[0].(*v10.TaskStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UninstallStatus indicates an expected call of UninstallStatus.
func (mr *MockBackendServerMockRecorder) UninstallStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallStatus", reflect.TypeOf((*MockBackendServer)(nil).UninstallStatus), arg0, arg1)
}

// mustEmbedUnimplementedBackendServer mocks base method.
func (m *MockBackendServer) mustEmbedUnimplementedBackendServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBackendServer")
}

// mustEmbedUnimplementedBackendServer indicates an expected call of mustEmbedUnimplementedBackendServer.
func (mr *MockBackendServerMockRecorder) mustEmbedUnimplementedBackendServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBackendServer", reflect.TypeOf((*MockBackendServer)(nil).mustEmbedUnimplementedBackendServer))
}

// MockUnsafeBackendServer is a mock of UnsafeBackendServer interface.
type MockUnsafeBackendServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeBackendServerMockRecorder
}

// MockUnsafeBackendServerMockRecorder is the mock recorder for MockUnsafeBackendServer.
type MockUnsafeBackendServerMockRecorder struct {
	mock *MockUnsafeBackendServer
}

// NewMockUnsafeBackendServer creates a new mock instance.
func NewMockUnsafeBackendServer(ctrl *gomock.Controller) *MockUnsafeBackendServer {
	mock := &MockUnsafeBackendServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeBackendServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeBackendServer) EXPECT() *MockUnsafeBackendServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedBackendServer mocks base method.
func (m *MockUnsafeBackendServer) mustEmbedUnimplementedBackendServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBackendServer")
}

// mustEmbedUnimplementedBackendServer indicates an expected call of mustEmbedUnimplementedBackendServer.
func (mr *MockUnsafeBackendServerMockRecorder) mustEmbedUnimplementedBackendServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBackendServer", reflect.TypeOf((*MockUnsafeBackendServer)(nil).mustEmbedUnimplementedBackendServer))
}

// MockNodeClient is a mock of NodeClient interface.
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient.
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance.
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// SyncNow mocks base method.
func (m *MockNodeClient) SyncNow(ctx context.Context, in *v1.Filter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncNow", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncNow indicates an expected call of SyncNow.
func (mr *MockNodeClientMockRecorder) SyncNow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncNow", reflect.TypeOf((*MockNodeClient)(nil).SyncNow), varargs...)
}

// MockNodeServer is a mock of NodeServer interface.
type MockNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeServerMockRecorder
}

// MockNodeServerMockRecorder is the mock recorder for MockNodeServer.
type MockNodeServerMockRecorder struct {
	mock *MockNodeServer
}

// NewMockNodeServer creates a new mock instance.
func NewMockNodeServer(ctrl *gomock.Controller) *MockNodeServer {
	mock := &MockNodeServer{ctrl: ctrl}
	mock.recorder = &MockNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeServer) EXPECT() *MockNodeServerMockRecorder {
	return m.recorder
}

// SyncNow mocks base method.
func (m *MockNodeServer) SyncNow(arg0 context.Context, arg1 *v1.Filter) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncNow", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncNow indicates an expected call of SyncNow.
func (mr *MockNodeServerMockRecorder) SyncNow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncNow", reflect.TypeOf((*MockNodeServer)(nil).SyncNow), arg0, arg1)
}

// mustEmbedUnimplementedNodeServer mocks base method.
func (m *MockNodeServer) mustEmbedUnimplementedNodeServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedNodeServer")
}

// mustEmbedUnimplementedNodeServer indicates an expected call of mustEmbedUnimplementedNodeServer.
func (mr *MockNodeServerMockRecorder) mustEmbedUnimplementedNodeServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedNodeServer", reflect.TypeOf((*MockNodeServer)(nil).mustEmbedUnimplementedNodeServer))
}

// MockUnsafeNodeServer is a mock of UnsafeNodeServer interface.
type MockUnsafeNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeNodeServerMockRecorder
}

// MockUnsafeNodeServerMockRecorder is the mock recorder for MockUnsafeNodeServer.
type MockUnsafeNodeServerMockRecorder struct {
	mock *MockUnsafeNodeServer
}

// NewMockUnsafeNodeServer creates a new mock instance.
func NewMockUnsafeNodeServer(ctrl *gomock.Controller) *MockUnsafeNodeServer {
	mock := &MockUnsafeNodeServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeNodeServer) EXPECT() *MockUnsafeNodeServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedNodeServer mocks base method.
func (m *MockUnsafeNodeServer) mustEmbedUnimplementedNodeServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedNodeServer")
}

// mustEmbedUnimplementedNodeServer indicates an expected call of mustEmbedUnimplementedNodeServer.
func (mr *MockUnsafeNodeServerMockRecorder) mustEmbedUnimplementedNodeServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedNodeServer", reflect.TypeOf((*MockUnsafeNodeServer)(nil).mustEmbedUnimplementedNodeServer))
}

// MockNodeManagerClient is a mock of NodeManagerClient interface.
type MockNodeManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerClientMockRecorder
}

// MockNodeManagerClientMockRecorder is the mock recorder for MockNodeManagerClient.
type MockNodeManagerClientMockRecorder struct {
	mock *MockNodeManagerClient
}

// NewMockNodeManagerClient creates a new mock instance.
func NewMockNodeManagerClient(ctrl *gomock.Controller) *MockNodeManagerClient {
	mock := &MockNodeManagerClient{ctrl: ctrl}
	mock.recorder = &MockNodeManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeManagerClient) EXPECT() *MockNodeManagerClientMockRecorder {
	return m.recorder
}

// RequestSync mocks base method.
func (m *MockNodeManagerClient) RequestSync(ctx context.Context, in *v1.SyncRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestSync", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestSync indicates an expected call of RequestSync.
func (mr *MockNodeManagerClientMockRecorder) RequestSync(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestSync", reflect.TypeOf((*MockNodeManagerClient)(nil).RequestSync), varargs...)
}

// MockNodeManagerServer is a mock of NodeManagerServer interface.
type MockNodeManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerServerMockRecorder
}

// MockNodeManagerServerMockRecorder is the mock recorder for MockNodeManagerServer.
type MockNodeManagerServerMockRecorder struct {
	mock *MockNodeManagerServer
}

// NewMockNodeManagerServer creates a new mock instance.
func NewMockNodeManagerServer(ctrl *gomock.Controller) *MockNodeManagerServer {
	mock := &MockNodeManagerServer{ctrl: ctrl}
	mock.recorder = &MockNodeManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeManagerServer) EXPECT() *MockNodeManagerServerMockRecorder {
	return m.recorder
}

// RequestSync mocks base method.
func (m *MockNodeManagerServer) RequestSync(arg0 context.Context, arg1 *v1.SyncRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestSync", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestSync indicates an expected call of RequestSync.
func (mr *MockNodeManagerServerMockRecorder) RequestSync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestSync", reflect.TypeOf((*MockNodeManagerServer)(nil).RequestSync), arg0, arg1)
}

// mustEmbedUnimplementedNodeManagerServer mocks base method.
func (m *MockNodeManagerServer) mustEmbedUnimplementedNodeManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedNodeManagerServer")
}

// mustEmbedUnimplementedNodeManagerServer indicates an expected call of mustEmbedUnimplementedNodeManagerServer.
func (mr *MockNodeManagerServerMockRecorder) mustEmbedUnimplementedNodeManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedNodeManagerServer", reflect.TypeOf((*MockNodeManagerServer)(nil).mustEmbedUnimplementedNodeManagerServer))
}

// MockUnsafeNodeManagerServer is a mock of UnsafeNodeManagerServer interface.
type MockUnsafeNodeManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeNodeManagerServerMockRecorder
}

// MockUnsafeNodeManagerServerMockRecorder is the mock recorder for MockUnsafeNodeManagerServer.
type MockUnsafeNodeManagerServerMockRecorder struct {
	mock *MockUnsafeNodeManagerServer
}

// NewMockUnsafeNodeManagerServer creates a new mock instance.
func NewMockUnsafeNodeManagerServer(ctrl *gomock.Controller) *MockUnsafeNodeManagerServer {
	mock := &MockUnsafeNodeManagerServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeNodeManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeNodeManagerServer) EXPECT() *MockUnsafeNodeManagerServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedNodeManagerServer mocks base method.
func (m *MockUnsafeNodeManagerServer) mustEmbedUnimplementedNodeManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedNodeManagerServer")
}

// mustEmbedUnimplementedNodeManagerServer indicates an expected call of mustEmbedUnimplementedNodeManagerServer.
func (mr *MockUnsafeNodeManagerServerMockRecorder) mustEmbedUnimplementedNodeManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedNodeManagerServer", reflect.TypeOf((*MockUnsafeNodeManagerServer)(nil).mustEmbedUnimplementedNodeManagerServer))
}
