// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: payment.proto

package booking

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

const (
	PaymentService_Create_FullMethodName           = "/booking.PaymentService/Create"
	PaymentService_GetById_FullMethodName          = "/booking.PaymentService/GetById"
	PaymentService_GetAll_FullMethodName           = "/booking.PaymentService/GetAll"
	PaymentService_Update_FullMethodName           = "/booking.PaymentService/Update"
	PaymentService_Delete_FullMethodName           = "/booking.PaymentService/Delete"
	PaymentService_GetBookingId_FullMethodName     = "/booking.PaymentService/GetBookingId"
	PaymentService_GetBookingAmount_FullMethodName = "/booking.PaymentService/GetBookingAmount"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	Create(ctx context.Context, in *PaymentRes, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*PaymentGetByIdRes, error)
	GetAll(ctx context.Context, in *PaymentGetAllReq, opts ...grpc.CallOption) (*PaymentGetAllRes, error)
	Update(ctx context.Context, in *PaymentUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetBookingId(ctx context.Context, in *ById, opts ...grpc.CallOption) (*ById, error)
	GetBookingAmount(ctx context.Context, in *ById, opts ...grpc.CallOption) (*GetAmountRes, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) Create(ctx context.Context, in *PaymentRes, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, PaymentService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*PaymentGetByIdRes, error) {
	out := new(PaymentGetByIdRes)
	err := c.cc.Invoke(ctx, PaymentService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetAll(ctx context.Context, in *PaymentGetAllReq, opts ...grpc.CallOption) (*PaymentGetAllRes, error) {
	out := new(PaymentGetAllRes)
	err := c.cc.Invoke(ctx, PaymentService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) Update(ctx context.Context, in *PaymentUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, PaymentService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, PaymentService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetBookingId(ctx context.Context, in *ById, opts ...grpc.CallOption) (*ById, error) {
	out := new(ById)
	err := c.cc.Invoke(ctx, PaymentService_GetBookingId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetBookingAmount(ctx context.Context, in *ById, opts ...grpc.CallOption) (*GetAmountRes, error) {
	out := new(GetAmountRes)
	err := c.cc.Invoke(ctx, PaymentService_GetBookingAmount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	Create(context.Context, *PaymentRes) (*Void, error)
	GetById(context.Context, *ById) (*PaymentGetByIdRes, error)
	GetAll(context.Context, *PaymentGetAllReq) (*PaymentGetAllRes, error)
	Update(context.Context, *PaymentUpdateReq) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	GetBookingId(context.Context, *ById) (*ById, error)
	GetBookingAmount(context.Context, *ById) (*GetAmountRes, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) Create(context.Context, *PaymentRes) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPaymentServiceServer) GetById(context.Context, *ById) (*PaymentGetByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedPaymentServiceServer) GetAll(context.Context, *PaymentGetAllReq) (*PaymentGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPaymentServiceServer) Update(context.Context, *PaymentUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPaymentServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPaymentServiceServer) GetBookingId(context.Context, *ById) (*ById, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingId not implemented")
}
func (UnimplementedPaymentServiceServer) GetBookingAmount(context.Context, *ById) (*GetAmountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingAmount not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Create(ctx, req.(*PaymentRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetAll(ctx, req.(*PaymentGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Update(ctx, req.(*PaymentUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetBookingId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetBookingId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetBookingId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetBookingId(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetBookingAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetBookingAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetBookingAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetBookingAmount(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PaymentService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _PaymentService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PaymentService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PaymentService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PaymentService_Delete_Handler,
		},
		{
			MethodName: "GetBookingId",
			Handler:    _PaymentService_GetBookingId_Handler,
		},
		{
			MethodName: "GetBookingAmount",
			Handler:    _PaymentService_GetBookingAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}