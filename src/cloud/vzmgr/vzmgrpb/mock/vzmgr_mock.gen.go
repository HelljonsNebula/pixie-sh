// Code generated by MockGen. DO NOT EDIT.
// Source: service.pb.go

// Package mock_vzmgrpb is a generated GoMock package.
package mock_vzmgrpb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	cloudpb "pixielabs.ai/pixielabs/src/cloud/cloudpb"
	vzmgrpb "pixielabs.ai/pixielabs/src/cloud/vzmgr/vzmgrpb"
	proto "pixielabs.ai/pixielabs/src/common/uuid/proto"
	reflect "reflect"
)

// MockVZMgrServiceClient is a mock of VZMgrServiceClient interface
type MockVZMgrServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockVZMgrServiceClientMockRecorder
}

// MockVZMgrServiceClientMockRecorder is the mock recorder for MockVZMgrServiceClient
type MockVZMgrServiceClientMockRecorder struct {
	mock *MockVZMgrServiceClient
}

// NewMockVZMgrServiceClient creates a new mock instance
func NewMockVZMgrServiceClient(ctrl *gomock.Controller) *MockVZMgrServiceClient {
	mock := &MockVZMgrServiceClient{ctrl: ctrl}
	mock.recorder = &MockVZMgrServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVZMgrServiceClient) EXPECT() *MockVZMgrServiceClientMockRecorder {
	return m.recorder
}

// CreateVizierCluster mocks base method
func (m *MockVZMgrServiceClient) CreateVizierCluster(ctx context.Context, in *vzmgrpb.CreateVizierClusterRequest, opts ...grpc.CallOption) (*proto.UUID, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateVizierCluster", varargs...)
	ret0, _ := ret[0].(*proto.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVizierCluster indicates an expected call of CreateVizierCluster
func (mr *MockVZMgrServiceClientMockRecorder) CreateVizierCluster(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVizierCluster", reflect.TypeOf((*MockVZMgrServiceClient)(nil).CreateVizierCluster), varargs...)
}

// GetViziersByOrg mocks base method
func (m *MockVZMgrServiceClient) GetViziersByOrg(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*vzmgrpb.GetViziersByOrgResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetViziersByOrg", varargs...)
	ret0, _ := ret[0].(*vzmgrpb.GetViziersByOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViziersByOrg indicates an expected call of GetViziersByOrg
func (mr *MockVZMgrServiceClientMockRecorder) GetViziersByOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViziersByOrg", reflect.TypeOf((*MockVZMgrServiceClient)(nil).GetViziersByOrg), varargs...)
}

// GetVizierInfo mocks base method
func (m *MockVZMgrServiceClient) GetVizierInfo(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*cloudpb.VizierInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVizierInfo", varargs...)
	ret0, _ := ret[0].(*cloudpb.VizierInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVizierInfo indicates an expected call of GetVizierInfo
func (mr *MockVZMgrServiceClientMockRecorder) GetVizierInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVizierInfo", reflect.TypeOf((*MockVZMgrServiceClient)(nil).GetVizierInfo), varargs...)
}

// GetVizierConnectionInfo mocks base method
func (m *MockVZMgrServiceClient) GetVizierConnectionInfo(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*cloudpb.VizierConnectionInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVizierConnectionInfo", varargs...)
	ret0, _ := ret[0].(*cloudpb.VizierConnectionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVizierConnectionInfo indicates an expected call of GetVizierConnectionInfo
func (mr *MockVZMgrServiceClientMockRecorder) GetVizierConnectionInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVizierConnectionInfo", reflect.TypeOf((*MockVZMgrServiceClient)(nil).GetVizierConnectionInfo), varargs...)
}

// VizierConnected mocks base method
func (m *MockVZMgrServiceClient) VizierConnected(ctx context.Context, in *cloudpb.RegisterVizierRequest, opts ...grpc.CallOption) (*cloudpb.RegisterVizierAck, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VizierConnected", varargs...)
	ret0, _ := ret[0].(*cloudpb.RegisterVizierAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VizierConnected indicates an expected call of VizierConnected
func (mr *MockVZMgrServiceClientMockRecorder) VizierConnected(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VizierConnected", reflect.TypeOf((*MockVZMgrServiceClient)(nil).VizierConnected), varargs...)
}

// HandleVizierHeartbeat mocks base method
func (m *MockVZMgrServiceClient) HandleVizierHeartbeat(ctx context.Context, in *cloudpb.VizierHeartbeat, opts ...grpc.CallOption) (*cloudpb.VizierHeartbeatAck, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HandleVizierHeartbeat", varargs...)
	ret0, _ := ret[0].(*cloudpb.VizierHeartbeatAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleVizierHeartbeat indicates an expected call of HandleVizierHeartbeat
func (mr *MockVZMgrServiceClientMockRecorder) HandleVizierHeartbeat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleVizierHeartbeat", reflect.TypeOf((*MockVZMgrServiceClient)(nil).HandleVizierHeartbeat), varargs...)
}

// MockVZMgrServiceServer is a mock of VZMgrServiceServer interface
type MockVZMgrServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockVZMgrServiceServerMockRecorder
}

// MockVZMgrServiceServerMockRecorder is the mock recorder for MockVZMgrServiceServer
type MockVZMgrServiceServerMockRecorder struct {
	mock *MockVZMgrServiceServer
}

// NewMockVZMgrServiceServer creates a new mock instance
func NewMockVZMgrServiceServer(ctrl *gomock.Controller) *MockVZMgrServiceServer {
	mock := &MockVZMgrServiceServer{ctrl: ctrl}
	mock.recorder = &MockVZMgrServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVZMgrServiceServer) EXPECT() *MockVZMgrServiceServerMockRecorder {
	return m.recorder
}

// CreateVizierCluster mocks base method
func (m *MockVZMgrServiceServer) CreateVizierCluster(arg0 context.Context, arg1 *vzmgrpb.CreateVizierClusterRequest) (*proto.UUID, error) {
	ret := m.ctrl.Call(m, "CreateVizierCluster", arg0, arg1)
	ret0, _ := ret[0].(*proto.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVizierCluster indicates an expected call of CreateVizierCluster
func (mr *MockVZMgrServiceServerMockRecorder) CreateVizierCluster(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVizierCluster", reflect.TypeOf((*MockVZMgrServiceServer)(nil).CreateVizierCluster), arg0, arg1)
}

// GetViziersByOrg mocks base method
func (m *MockVZMgrServiceServer) GetViziersByOrg(arg0 context.Context, arg1 *proto.UUID) (*vzmgrpb.GetViziersByOrgResponse, error) {
	ret := m.ctrl.Call(m, "GetViziersByOrg", arg0, arg1)
	ret0, _ := ret[0].(*vzmgrpb.GetViziersByOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViziersByOrg indicates an expected call of GetViziersByOrg
func (mr *MockVZMgrServiceServerMockRecorder) GetViziersByOrg(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViziersByOrg", reflect.TypeOf((*MockVZMgrServiceServer)(nil).GetViziersByOrg), arg0, arg1)
}

// GetVizierInfo mocks base method
func (m *MockVZMgrServiceServer) GetVizierInfo(arg0 context.Context, arg1 *proto.UUID) (*cloudpb.VizierInfo, error) {
	ret := m.ctrl.Call(m, "GetVizierInfo", arg0, arg1)
	ret0, _ := ret[0].(*cloudpb.VizierInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVizierInfo indicates an expected call of GetVizierInfo
func (mr *MockVZMgrServiceServerMockRecorder) GetVizierInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVizierInfo", reflect.TypeOf((*MockVZMgrServiceServer)(nil).GetVizierInfo), arg0, arg1)
}

// GetVizierConnectionInfo mocks base method
func (m *MockVZMgrServiceServer) GetVizierConnectionInfo(arg0 context.Context, arg1 *proto.UUID) (*cloudpb.VizierConnectionInfo, error) {
	ret := m.ctrl.Call(m, "GetVizierConnectionInfo", arg0, arg1)
	ret0, _ := ret[0].(*cloudpb.VizierConnectionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVizierConnectionInfo indicates an expected call of GetVizierConnectionInfo
func (mr *MockVZMgrServiceServerMockRecorder) GetVizierConnectionInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVizierConnectionInfo", reflect.TypeOf((*MockVZMgrServiceServer)(nil).GetVizierConnectionInfo), arg0, arg1)
}

// VizierConnected mocks base method
func (m *MockVZMgrServiceServer) VizierConnected(arg0 context.Context, arg1 *cloudpb.RegisterVizierRequest) (*cloudpb.RegisterVizierAck, error) {
	ret := m.ctrl.Call(m, "VizierConnected", arg0, arg1)
	ret0, _ := ret[0].(*cloudpb.RegisterVizierAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VizierConnected indicates an expected call of VizierConnected
func (mr *MockVZMgrServiceServerMockRecorder) VizierConnected(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VizierConnected", reflect.TypeOf((*MockVZMgrServiceServer)(nil).VizierConnected), arg0, arg1)
}

// HandleVizierHeartbeat mocks base method
func (m *MockVZMgrServiceServer) HandleVizierHeartbeat(arg0 context.Context, arg1 *cloudpb.VizierHeartbeat) (*cloudpb.VizierHeartbeatAck, error) {
	ret := m.ctrl.Call(m, "HandleVizierHeartbeat", arg0, arg1)
	ret0, _ := ret[0].(*cloudpb.VizierHeartbeatAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleVizierHeartbeat indicates an expected call of HandleVizierHeartbeat
func (mr *MockVZMgrServiceServerMockRecorder) HandleVizierHeartbeat(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleVizierHeartbeat", reflect.TypeOf((*MockVZMgrServiceServer)(nil).HandleVizierHeartbeat), arg0, arg1)
}
