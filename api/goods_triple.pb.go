// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.20.1
// source: goods.proto

package api

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// GoodsServiceClient is the client API for GoodsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsServiceClient interface {
	GetGoodsById(ctx context.Context, in *GoodsId, opts ...grpc_go.CallOption) (*Goods, common.ErrorWithAttachment)
	GetGoodsByIds(ctx context.Context, in *GoodsIds, opts ...grpc_go.CallOption) (*Goodss, common.ErrorWithAttachment)
}

type goodsServiceClient struct {
	cc *triple.TripleConn
}

type GoodsServiceClientImpl struct {
	GetGoodsById  func(ctx context.Context, in *GoodsId) (*Goods, error)
	GetGoodsByIds func(ctx context.Context, in *GoodsIds) (*Goodss, error)
}

func (c *GoodsServiceClientImpl) GetDubboStub(cc *triple.TripleConn) GoodsServiceClient {
	return NewGoodsServiceClient(cc)
}

func (c *GoodsServiceClientImpl) XXX_InterfaceName() string {
	return "api.GoodsService"
}

func NewGoodsServiceClient(cc *triple.TripleConn) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) GetGoodsById(ctx context.Context, in *GoodsId, opts ...grpc_go.CallOption) (*Goods, common.ErrorWithAttachment) {
	out := new(Goods)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetGoodsById", in, out)
}

func (c *goodsServiceClient) GetGoodsByIds(ctx context.Context, in *GoodsIds, opts ...grpc_go.CallOption) (*Goodss, common.ErrorWithAttachment) {
	out := new(Goodss)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetGoodsByIds", in, out)
}

// GoodsServiceServer is the server API for GoodsService service.
// All implementations must embed UnimplementedGoodsServiceServer
// for forward compatibility
type GoodsServiceServer interface {
	GetGoodsById(context.Context, *GoodsId) (*Goods, error)
	GetGoodsByIds(context.Context, *GoodsIds) (*Goodss, error)
	mustEmbedUnimplementedGoodsServiceServer()
}

// UnimplementedGoodsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServiceServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedGoodsServiceServer) GetGoodsById(context.Context, *GoodsId) (*Goods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsById not implemented")
}
func (UnimplementedGoodsServiceServer) GetGoodsByIds(context.Context, *GoodsIds) (*Goodss, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsByIds not implemented")
}
func (s *UnimplementedGoodsServiceServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedGoodsServiceServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedGoodsServiceServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &GoodsService_ServiceDesc
}
func (s *UnimplementedGoodsServiceServer) XXX_InterfaceName() string {
	return "api.GoodsService"
}

func (UnimplementedGoodsServiceServer) mustEmbedUnimplementedGoodsServiceServer() {}

// UnsafeGoodsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServiceServer will
// result in compilation errors.
type UnsafeGoodsServiceServer interface {
	mustEmbedUnimplementedGoodsServiceServer()
}

func RegisterGoodsServiceServer(s grpc_go.ServiceRegistrar, srv GoodsServiceServer) {
	s.RegisterService(&GoodsService_ServiceDesc, srv)
}

func _GoodsService_GetGoodsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetGoodsById", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GetGoodsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetGoodsByIds", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsService_ServiceDesc is the grpc_go.ServiceDesc for GoodsService service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsService_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "api.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "GetGoodsById",
			Handler:    _GoodsService_GetGoodsById_Handler,
		},
		{
			MethodName: "GetGoodsByIds",
			Handler:    _GoodsService_GetGoodsByIds_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "goods.proto",
}
