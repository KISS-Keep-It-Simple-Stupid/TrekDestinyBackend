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
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	GetMyPost(ctx context.Context, in *GetMyPostRequest, opts ...grpc.CallOption) (*GetMyPostResponse, error)
	GetPostHost(ctx context.Context, in *GetPostHostRequest, opts ...grpc.CallOption) (*GetPostHostResponse, error)
	AcceptOffer(ctx context.Context, in *AcceptOfferRequest, opts ...grpc.CallOption) (*AcceptOfferResponse, error)
	RejectOffer(ctx context.Context, in *RejectOfferRequest, opts ...grpc.CallOption) (*RejectOfferResponse, error)
	EditAnnouncement(ctx context.Context, in *EditAnnouncementRequest, opts ...grpc.CallOption) (*EditAnnouncementResponse, error)
	DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, opts ...grpc.CallOption) (*DeleteAnnouncementResponse, error)
	EditPost(ctx context.Context, in *EditPostRequest, opts ...grpc.CallOption) (*EditPostResponse, error)
	UploadHostHouseImage(ctx context.Context, in *HostHouseImageRequest, opts ...grpc.CallOption) (*HostHouseImageResponse, error)
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

func (c *announcementClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/Announcement/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) GetMyPost(ctx context.Context, in *GetMyPostRequest, opts ...grpc.CallOption) (*GetMyPostResponse, error) {
	out := new(GetMyPostResponse)
	err := c.cc.Invoke(ctx, "/Announcement/GetMyPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) GetPostHost(ctx context.Context, in *GetPostHostRequest, opts ...grpc.CallOption) (*GetPostHostResponse, error) {
	out := new(GetPostHostResponse)
	err := c.cc.Invoke(ctx, "/Announcement/GetPostHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) AcceptOffer(ctx context.Context, in *AcceptOfferRequest, opts ...grpc.CallOption) (*AcceptOfferResponse, error) {
	out := new(AcceptOfferResponse)
	err := c.cc.Invoke(ctx, "/Announcement/AcceptOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) RejectOffer(ctx context.Context, in *RejectOfferRequest, opts ...grpc.CallOption) (*RejectOfferResponse, error) {
	out := new(RejectOfferResponse)
	err := c.cc.Invoke(ctx, "/Announcement/RejectOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) EditAnnouncement(ctx context.Context, in *EditAnnouncementRequest, opts ...grpc.CallOption) (*EditAnnouncementResponse, error) {
	out := new(EditAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/Announcement/EditAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, opts ...grpc.CallOption) (*DeleteAnnouncementResponse, error) {
	out := new(DeleteAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/Announcement/DeleteAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) EditPost(ctx context.Context, in *EditPostRequest, opts ...grpc.CallOption) (*EditPostResponse, error) {
	out := new(EditPostResponse)
	err := c.cc.Invoke(ctx, "/Announcement/EditPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementClient) UploadHostHouseImage(ctx context.Context, in *HostHouseImageRequest, opts ...grpc.CallOption) (*HostHouseImageResponse, error) {
	out := new(HostHouseImageResponse)
	err := c.cc.Invoke(ctx, "/Announcement/UploadHostHouseImage", in, out, opts...)
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
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	GetMyPost(context.Context, *GetMyPostRequest) (*GetMyPostResponse, error)
	GetPostHost(context.Context, *GetPostHostRequest) (*GetPostHostResponse, error)
	AcceptOffer(context.Context, *AcceptOfferRequest) (*AcceptOfferResponse, error)
	RejectOffer(context.Context, *RejectOfferRequest) (*RejectOfferResponse, error)
	EditAnnouncement(context.Context, *EditAnnouncementRequest) (*EditAnnouncementResponse, error)
	DeleteAnnouncement(context.Context, *DeleteAnnouncementRequest) (*DeleteAnnouncementResponse, error)
	EditPost(context.Context, *EditPostRequest) (*EditPostResponse, error)
	UploadHostHouseImage(context.Context, *HostHouseImageRequest) (*HostHouseImageResponse, error)
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
func (UnimplementedAnnouncementServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedAnnouncementServer) GetMyPost(context.Context, *GetMyPostRequest) (*GetMyPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyPost not implemented")
}
func (UnimplementedAnnouncementServer) GetPostHost(context.Context, *GetPostHostRequest) (*GetPostHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostHost not implemented")
}
func (UnimplementedAnnouncementServer) AcceptOffer(context.Context, *AcceptOfferRequest) (*AcceptOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptOffer not implemented")
}
func (UnimplementedAnnouncementServer) RejectOffer(context.Context, *RejectOfferRequest) (*RejectOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectOffer not implemented")
}
func (UnimplementedAnnouncementServer) EditAnnouncement(context.Context, *EditAnnouncementRequest) (*EditAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAnnouncement not implemented")
}
func (UnimplementedAnnouncementServer) DeleteAnnouncement(context.Context, *DeleteAnnouncementRequest) (*DeleteAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnnouncement not implemented")
}
func (UnimplementedAnnouncementServer) EditPost(context.Context, *EditPostRequest) (*EditPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPost not implemented")
}
func (UnimplementedAnnouncementServer) UploadHostHouseImage(context.Context, *HostHouseImageRequest) (*HostHouseImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadHostHouseImage not implemented")
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

func _Announcement_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_GetMyPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).GetMyPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/GetMyPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).GetMyPost(ctx, req.(*GetMyPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_GetPostHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).GetPostHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/GetPostHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).GetPostHost(ctx, req.(*GetPostHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_AcceptOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).AcceptOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/AcceptOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).AcceptOffer(ctx, req.(*AcceptOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_RejectOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).RejectOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/RejectOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).RejectOffer(ctx, req.(*RejectOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_EditAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).EditAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/EditAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).EditAnnouncement(ctx, req.(*EditAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_DeleteAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).DeleteAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/DeleteAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).DeleteAnnouncement(ctx, req.(*DeleteAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_EditPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).EditPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/EditPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).EditPost(ctx, req.(*EditPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Announcement_UploadHostHouseImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostHouseImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServer).UploadHostHouseImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Announcement/UploadHostHouseImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServer).UploadHostHouseImage(ctx, req.(*HostHouseImageRequest))
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
		{
			MethodName: "CreatePost",
			Handler:    _Announcement_CreatePost_Handler,
		},
		{
			MethodName: "GetMyPost",
			Handler:    _Announcement_GetMyPost_Handler,
		},
		{
			MethodName: "GetPostHost",
			Handler:    _Announcement_GetPostHost_Handler,
		},
		{
			MethodName: "AcceptOffer",
			Handler:    _Announcement_AcceptOffer_Handler,
		},
		{
			MethodName: "RejectOffer",
			Handler:    _Announcement_RejectOffer_Handler,
		},
		{
			MethodName: "EditAnnouncement",
			Handler:    _Announcement_EditAnnouncement_Handler,
		},
		{
			MethodName: "DeleteAnnouncement",
			Handler:    _Announcement_DeleteAnnouncement_Handler,
		},
		{
			MethodName: "EditPost",
			Handler:    _Announcement_EditPost_Handler,
		},
		{
			MethodName: "UploadHostHouseImage",
			Handler:    _Announcement_UploadHostHouseImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "announcement.proto",
}
