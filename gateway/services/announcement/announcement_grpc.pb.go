// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: announcement.proto

package pb

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

// AnnouncementClient is the client API for Announcement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnnouncementClient interface {
	CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error)
	GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*GetCardResponse, error)
	CreateOffer(ctx context.Context, in *CreateOfferRequest, opts ...grpc.CallOption) (*CreateOfferResponse, error)
	GetOffer(ctx context.Context, in *GetOfferRequest, opts ...grpc.CallOption) (*GetOfferResponse, error)
	GetCardProfile(ctx context.Context, in *GetCardProfileRequest, opts ...grpc.CallOption) (*GetCardProfileResponse, error)
}

type announcementClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnouncementClient(cc grpc.ClientConnInterface) AnnouncementClient {
	return &announcementClient{cc}
}

func (c *announcementClient) CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error) {
	out := new(CreateCardResponse)
	err := c.cc.Invoke(ctx, "/Announcement/CreateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*GetCardResponse, error) {
	out := new(GetCardResponse)
	err := c.cc.Invoke(ctx, "/Announcement/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) CreateOffer(ctx context.Context, in *CreateOfferRequest, opts ...grpc.CallOption) (*CreateOfferResponse, error) {
	out := new(CreateOfferResponse)
	err := c.cc.Invoke(ctx, "/Announcement/CreateOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) GetOffer(ctx context.Context, in *GetOfferRequest, opts ...grpc.CallOption) (*GetOfferResponse, error) {
	out := new(GetOfferResponse)
	err := c.cc.Invoke(ctx, "/Announcement/GetOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) GetCardProfile(ctx context.Context, in *GetCardProfileRequest, opts ...grpc.CallOption) (*GetCardProfileResponse, error) {
	out := new(GetCardProfileResponse)
	err := c.cc.Invoke(ctx, "/Announcement/GetCardProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnouncementServer is the server API for Announcement service.
// All implementations must embed UnimplementedAnnouncementServer
// for forward compatibility
type AnnouncementServer interface {
	CreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error)
	GetCard(context.Context, *GetCardRequest) (*GetCardResponse, error)
	CreateOffer(context.Context, *CreateOfferRequest) (*CreateOfferResponse, error)
	GetOffer(context.Context, *GetOfferRequest) (*GetOfferResponse, error)
	GetCardProfile(context.Context, *GetCardProfileRequest) (*GetCardProfileResponse, error)
	mustEmbedUnimplementedAnnouncementServer()
}

// UnimplementedAnnouncementServer must be embedded to have forward compatible implementations.
type UnimplementedAnnouncementServer struct {
}

func (UnimplementedAnnouncementServer) CreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedAnnouncementServer) GetCard(context.Context, *GetCardRequest) (*GetCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (UnimplementedAnnouncementServer) CreateOffer(context.Context, *CreateOfferRequest) (*CreateOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOffer not implemented")
}
func (UnimplementedAnnouncementServer) GetOffer(context.Context, *GetOfferRequest) (*GetOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffer not implemented")
}
func (UnimplementedAnnouncementServer) GetCardProfile(context.Context, *GetCardProfileRequest) (*GetCardProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardProfile not implemented")
}
func (UnimplementedAnnouncementServer) mustEmbedUnimplementedAnnouncementServer() {}

// UnsafeAnnouncementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnnouncementServer will
// result in compilation errors.
type UnsafeAnnouncementServer interface {
	mustEmbedUnimplementedAnnouncementServer()
}

func RegisterAnnouncementServer(s grpc.ServiceRegistrar, srv AnnouncementServer) {
	s.RegisterService(&Announcement_ServiceDesc, srv)
}

func _Announcement_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/CreateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).CreateCard(ctx, req.(*CreateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).GetCard(ctx, req.(*GetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_CreateOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).CreateOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/CreateOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).CreateOffer(ctx, req.(*CreateOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_GetOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).GetOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/GetOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).GetOffer(ctx, req.(*GetOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_GetCardProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).GetCardProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/GetCardProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).GetCardProfile(ctx, req.(*GetCardProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Announcement_ServiceDesc is the grpc.ServiceDesc for Announcement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Announcement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Announcement",
	HandlerType: (*AnnouncementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCard",
			Handler:    _Announcement_CreateCard_Handler,
		},
		{
			MethodName: "GetCard",
			Handler:    _Announcement_GetCard_Handler,
		},
		{
			MethodName: "CreateOffer",
			Handler:    _Announcement_CreateOffer_Handler,
		},
		{
			MethodName: "GetOffer",
			Handler:    _Announcement_GetOffer_Handler,
		},
		{
			MethodName: "GetCardProfile",
			Handler:    _Announcement_GetCardProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "announcement.proto",
}
