// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: token_manager/v1/token_manager.proto

package token_managerv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TokenManagerServiceClient is the client API for TokenManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenManagerServiceClient interface {
	// 检测
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// 查询广告计划学习期状态
	GetAdStatExtraInfo(ctx context.Context, in *GetAdStatExtraInfoRequest, opts ...grpc.CallOption) (*GetAdStatExtraInfoResponse, error)
	// 获取广告主
	FindAdvertiser(ctx context.Context, in *FindAdvertiserRequest, opts ...grpc.CallOption) (*FindAdvertiserResponse, error)
	// 根据广告主获取用户
	FindUser(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserResponse, error)
	// 获取行业分类
	FindIndustryList(ctx context.Context, in *FindIndustryListRequest, opts ...grpc.CallOption) (*FindIndustryListResponse, error)
	// 获取日流水
	QianchuanFinanceDetailGet(ctx context.Context, in *QianchuanFinanceDetailGetRequest, opts ...grpc.CallOption) (*QianchuanFinanceDetailGetResponse, error)
	// 获取素材id
	FindMaterial(ctx context.Context, in *FindMaterialRequest, opts ...grpc.CallOption) (*FindMaterialResponse, error)
}

type tokenManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenManagerServiceClient(cc grpc.ClientConnInterface) TokenManagerServiceClient {
	return &tokenManagerServiceClient{cc}
}

func (c *tokenManagerServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/token_manager.v1.TokenManagerService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerServiceClient) GetAdStatExtraInfo(ctx context.Context, in *GetAdStatExtraInfoRequest, opts ...grpc.CallOption) (*GetAdStatExtraInfoResponse, error) {
	out := new(GetAdStatExtraInfoResponse)
	err := c.cc.Invoke(ctx, "/token_manager.v1.TokenManagerService/GetAdStatExtraInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerServiceClient) FindAdvertiser(ctx context.Context, in *FindAdvertiserRequest, opts ...grpc.CallOption) (*FindAdvertiserResponse, error) {
	out := new(FindAdvertiserResponse)
	err := c.cc.Invoke(ctx, "/token_manager.v1.TokenManagerService/FindAdvertiser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerServiceClient) FindUser(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserResponse, error) {
	out := new(FindUserResponse)
	err := c.cc.Invoke(ctx, "/token_manager.v1.TokenManagerService/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerServiceClient) FindIndustryList(ctx context.Context, in *FindIndustryListRequest, opts ...grpc.CallOption) (*FindIndustryListResponse, error) {
	out := new(FindIndustryListResponse)
	err := c.cc.Invoke(ctx, "/token_manager.v1.TokenManagerService/FindIndustryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerServiceClient) QianchuanFinanceDetailGet(ctx context.Context, in *QianchuanFinanceDetailGetRequest, opts ...grpc.CallOption) (*QianchuanFinanceDetailGetResponse, error) {
	out := new(QianchuanFinanceDetailGetResponse)
	err := c.cc.Invoke(ctx, "/token_manager.v1.TokenManagerService/QianchuanFinanceDetailGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerServiceClient) FindMaterial(ctx context.Context, in *FindMaterialRequest, opts ...grpc.CallOption) (*FindMaterialResponse, error) {
	out := new(FindMaterialResponse)
	err := c.cc.Invoke(ctx, "/token_manager.v1.TokenManagerService/FindMaterial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenManagerServiceServer is the server API for TokenManagerService service.
// All implementations must embed UnimplementedTokenManagerServiceServer
// for forward compatibility
type TokenManagerServiceServer interface {
	// 检测
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// 查询广告计划学习期状态
	GetAdStatExtraInfo(context.Context, *GetAdStatExtraInfoRequest) (*GetAdStatExtraInfoResponse, error)
	// 获取广告主
	FindAdvertiser(context.Context, *FindAdvertiserRequest) (*FindAdvertiserResponse, error)
	// 根据广告主获取用户
	FindUser(context.Context, *FindUserRequest) (*FindUserResponse, error)
	// 获取行业分类
	FindIndustryList(context.Context, *FindIndustryListRequest) (*FindIndustryListResponse, error)
	// 获取日流水
	QianchuanFinanceDetailGet(context.Context, *QianchuanFinanceDetailGetRequest) (*QianchuanFinanceDetailGetResponse, error)
	// 获取素材id
	FindMaterial(context.Context, *FindMaterialRequest) (*FindMaterialResponse, error)
	mustEmbedUnimplementedTokenManagerServiceServer()
}

// UnimplementedTokenManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenManagerServiceServer struct {
}

func (UnimplementedTokenManagerServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTokenManagerServiceServer) GetAdStatExtraInfo(context.Context, *GetAdStatExtraInfoRequest) (*GetAdStatExtraInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdStatExtraInfo not implemented")
}
func (UnimplementedTokenManagerServiceServer) FindAdvertiser(context.Context, *FindAdvertiserRequest) (*FindAdvertiserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAdvertiser not implemented")
}
func (UnimplementedTokenManagerServiceServer) FindUser(context.Context, *FindUserRequest) (*FindUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (UnimplementedTokenManagerServiceServer) FindIndustryList(context.Context, *FindIndustryListRequest) (*FindIndustryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindIndustryList not implemented")
}
func (UnimplementedTokenManagerServiceServer) QianchuanFinanceDetailGet(context.Context, *QianchuanFinanceDetailGetRequest) (*QianchuanFinanceDetailGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QianchuanFinanceDetailGet not implemented")
}
func (UnimplementedTokenManagerServiceServer) FindMaterial(context.Context, *FindMaterialRequest) (*FindMaterialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMaterial not implemented")
}
func (UnimplementedTokenManagerServiceServer) mustEmbedUnimplementedTokenManagerServiceServer() {}

// UnsafeTokenManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenManagerServiceServer will
// result in compilation errors.
type UnsafeTokenManagerServiceServer interface {
	mustEmbedUnimplementedTokenManagerServiceServer()
}

func RegisterTokenManagerServiceServer(s grpc.ServiceRegistrar, srv TokenManagerServiceServer) {
	s.RegisterService(&TokenManagerService_ServiceDesc, srv)
}

func _TokenManagerService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token_manager.v1.TokenManagerService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManagerService_GetAdStatExtraInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdStatExtraInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServiceServer).GetAdStatExtraInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token_manager.v1.TokenManagerService/GetAdStatExtraInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServiceServer).GetAdStatExtraInfo(ctx, req.(*GetAdStatExtraInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManagerService_FindAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServiceServer).FindAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token_manager.v1.TokenManagerService/FindAdvertiser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServiceServer).FindAdvertiser(ctx, req.(*FindAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManagerService_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServiceServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token_manager.v1.TokenManagerService/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServiceServer).FindUser(ctx, req.(*FindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManagerService_FindIndustryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindIndustryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServiceServer).FindIndustryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token_manager.v1.TokenManagerService/FindIndustryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServiceServer).FindIndustryList(ctx, req.(*FindIndustryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManagerService_QianchuanFinanceDetailGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QianchuanFinanceDetailGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServiceServer).QianchuanFinanceDetailGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token_manager.v1.TokenManagerService/QianchuanFinanceDetailGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServiceServer).QianchuanFinanceDetailGet(ctx, req.(*QianchuanFinanceDetailGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManagerService_FindMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMaterialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServiceServer).FindMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token_manager.v1.TokenManagerService/FindMaterial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServiceServer).FindMaterial(ctx, req.(*FindMaterialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenManagerService_ServiceDesc is the grpc.ServiceDesc for TokenManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "token_manager.v1.TokenManagerService",
	HandlerType: (*TokenManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _TokenManagerService_Ping_Handler,
		},
		{
			MethodName: "GetAdStatExtraInfo",
			Handler:    _TokenManagerService_GetAdStatExtraInfo_Handler,
		},
		{
			MethodName: "FindAdvertiser",
			Handler:    _TokenManagerService_FindAdvertiser_Handler,
		},
		{
			MethodName: "FindUser",
			Handler:    _TokenManagerService_FindUser_Handler,
		},
		{
			MethodName: "FindIndustryList",
			Handler:    _TokenManagerService_FindIndustryList_Handler,
		},
		{
			MethodName: "QianchuanFinanceDetailGet",
			Handler:    _TokenManagerService_QianchuanFinanceDetailGet_Handler,
		},
		{
			MethodName: "FindMaterial",
			Handler:    _TokenManagerService_FindMaterial_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token_manager/v1/token_manager.proto",
}
