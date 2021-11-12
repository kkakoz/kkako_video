// Code generated by MockGen. DO NOT EDIT.
// Source: api/user/v1/user_grpc.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserServiceClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUser", varargs...)
	ret0, _ := ret[0].(*AddUserRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserServiceClientMockRecorder) AddUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserServiceClient)(nil).AddUser), varargs...)
}

// GetUser mocks base method.
func (m *MockUserServiceClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*GetUserRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceClient)(nil).GetUser), varargs...)
}

// UserFollow mocks base method.
func (m *MockUserServiceClient) UserFollow(ctx context.Context, in *UserFollowReq, opts ...grpc.CallOption) (*UserFollowRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserFollow", varargs...)
	ret0, _ := ret[0].(*UserFollowRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserFollow indicates an expected call of UserFollow.
func (mr *MockUserServiceClientMockRecorder) UserFollow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserFollow", reflect.TypeOf((*MockUserServiceClient)(nil).UserFollow), varargs...)
}

// UserLike mocks base method.
func (m *MockUserServiceClient) UserLike(ctx context.Context, in *UserLikeReq, opts ...grpc.CallOption) (*UserLikeRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserLike", varargs...)
	ret0, _ := ret[0].(*UserLikeRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLike indicates an expected call of UserLike.
func (mr *MockUserServiceClientMockRecorder) UserLike(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLike", reflect.TypeOf((*MockUserServiceClient)(nil).UserLike), varargs...)
}

// UserNews mocks base method.
func (m *MockUserServiceClient) UserNews(ctx context.Context, in *UserNewsReq, opts ...grpc.CallOption) (*UserNewsRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserNews", varargs...)
	ret0, _ := ret[0].(*UserNewsRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserNews indicates an expected call of UserNews.
func (mr *MockUserServiceClientMockRecorder) UserNews(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserNews", reflect.TypeOf((*MockUserServiceClient)(nil).UserNews), varargs...)
}

// MockUserServiceServer is a mock of UserServiceServer interface.
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer.
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance.
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserServiceServer) AddUser(arg0 context.Context, arg1 *AddUserReq) (*AddUserRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(*AddUserRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserServiceServerMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserServiceServer)(nil).AddUser), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockUserServiceServer) GetUser(arg0 context.Context, arg1 *GetUserReq) (*GetUserRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*GetUserRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceServer)(nil).GetUser), arg0, arg1)
}

// UserFollow mocks base method.
func (m *MockUserServiceServer) UserFollow(arg0 context.Context, arg1 *UserFollowReq) (*UserFollowRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserFollow", arg0, arg1)
	ret0, _ := ret[0].(*UserFollowRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserFollow indicates an expected call of UserFollow.
func (mr *MockUserServiceServerMockRecorder) UserFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserFollow", reflect.TypeOf((*MockUserServiceServer)(nil).UserFollow), arg0, arg1)
}

// UserLike mocks base method.
func (m *MockUserServiceServer) UserLike(arg0 context.Context, arg1 *UserLikeReq) (*UserLikeRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLike", arg0, arg1)
	ret0, _ := ret[0].(*UserLikeRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLike indicates an expected call of UserLike.
func (mr *MockUserServiceServerMockRecorder) UserLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLike", reflect.TypeOf((*MockUserServiceServer)(nil).UserLike), arg0, arg1)
}

// UserNews mocks base method.
func (m *MockUserServiceServer) UserNews(arg0 context.Context, arg1 *UserNewsReq) (*UserNewsRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserNews", arg0, arg1)
	ret0, _ := ret[0].(*UserNewsRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserNews indicates an expected call of UserNews.
func (mr *MockUserServiceServerMockRecorder) UserNews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserNews", reflect.TypeOf((*MockUserServiceServer)(nil).UserNews), arg0, arg1)
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUnsafeUserServiceServer is a mock of UnsafeUserServiceServer interface.
type MockUnsafeUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserServiceServerMockRecorder
}

// MockUnsafeUserServiceServerMockRecorder is the mock recorder for MockUnsafeUserServiceServer.
type MockUnsafeUserServiceServerMockRecorder struct {
	mock *MockUnsafeUserServiceServer
}

// NewMockUnsafeUserServiceServer creates a new mock instance.
func NewMockUnsafeUserServiceServer(ctrl *gomock.Controller) *MockUnsafeUserServiceServer {
	mock := &MockUnsafeUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserServiceServer) EXPECT() *MockUnsafeUserServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUnsafeUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUnsafeUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}