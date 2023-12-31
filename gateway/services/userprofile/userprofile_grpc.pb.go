// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: userprofile.proto

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

// UserProfileClient is the client API for UserProfile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProfileClient interface {
	ProfileDetails(ctx context.Context, in *ProfileDetailsRequest, opts ...grpc.CallOption) (*ProfileDetailsResponse, error)
	EditProfile(ctx context.Context, in *EditProfileRequest, opts ...grpc.CallOption) (*EditProfileResponse, error)
	UploadImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	PublicProfile(ctx context.Context, in *PublicProfileRequest, opts ...grpc.CallOption) (*PublicProfileResponse, error)
	PublicProfileHost(ctx context.Context, in *PublicProfileHostRequest, opts ...grpc.CallOption) (*PublicProfileHostResponse, error)
	AddToChatList(ctx context.Context, in *AddChatListRequest, opts ...grpc.CallOption) (*AddChatListResponse, error)
	GetChatList(ctx context.Context, in *ChatListRequest, opts ...grpc.CallOption) (*ChatListResponse, error)
}

type userProfileClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProfileClient(cc grpc.ClientConnInterface) UserProfileClient {
	return &userProfileClient{cc}
}

func (c *userProfileClient) ProfileDetails(ctx context.Context, in *ProfileDetailsRequest, opts ...grpc.CallOption) (*ProfileDetailsResponse, error) {
	out := new(ProfileDetailsResponse)
	err := c.cc.Invoke(ctx, "/UserProfile/ProfileDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) EditProfile(ctx context.Context, in *EditProfileRequest, opts ...grpc.CallOption) (*EditProfileResponse, error) {
	out := new(EditProfileResponse)
	err := c.cc.Invoke(ctx, "/UserProfile/EditProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) UploadImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/UserProfile/UploadImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) PublicProfile(ctx context.Context, in *PublicProfileRequest, opts ...grpc.CallOption) (*PublicProfileResponse, error) {
	out := new(PublicProfileResponse)
	err := c.cc.Invoke(ctx, "/UserProfile/PublicProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) PublicProfileHost(ctx context.Context, in *PublicProfileHostRequest, opts ...grpc.CallOption) (*PublicProfileHostResponse, error) {
	out := new(PublicProfileHostResponse)
	err := c.cc.Invoke(ctx, "/UserProfile/PublicProfileHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) AddToChatList(ctx context.Context, in *AddChatListRequest, opts ...grpc.CallOption) (*AddChatListResponse, error) {
	out := new(AddChatListResponse)
	err := c.cc.Invoke(ctx, "/UserProfile/AddToChatList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) GetChatList(ctx context.Context, in *ChatListRequest, opts ...grpc.CallOption) (*ChatListResponse, error) {
	out := new(ChatListResponse)
	err := c.cc.Invoke(ctx, "/UserProfile/GetChatList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfileServer is the server API for UserProfile service.
// All implementations must embed UnimplementedUserProfileServer
// for forward compatibility
type UserProfileServer interface {
	ProfileDetails(context.Context, *ProfileDetailsRequest) (*ProfileDetailsResponse, error)
	EditProfile(context.Context, *EditProfileRequest) (*EditProfileResponse, error)
	UploadImage(context.Context, *ImageRequest) (*ImageResponse, error)
	PublicProfile(context.Context, *PublicProfileRequest) (*PublicProfileResponse, error)
	PublicProfileHost(context.Context, *PublicProfileHostRequest) (*PublicProfileHostResponse, error)
	AddToChatList(context.Context, *AddChatListRequest) (*AddChatListResponse, error)
	GetChatList(context.Context, *ChatListRequest) (*ChatListResponse, error)
	mustEmbedUnimplementedUserProfileServer()
}

// UnimplementedUserProfileServer must be embedded to have forward compatible implementations.
type UnimplementedUserProfileServer struct {
}

func (UnimplementedUserProfileServer) ProfileDetails(context.Context, *ProfileDetailsRequest) (*ProfileDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProfileDetails not implemented")
}
func (UnimplementedUserProfileServer) EditProfile(context.Context, *EditProfileRequest) (*EditProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedUserProfileServer) UploadImage(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedUserProfileServer) PublicProfile(context.Context, *PublicProfileRequest) (*PublicProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicProfile not implemented")
}
func (UnimplementedUserProfileServer) PublicProfileHost(context.Context, *PublicProfileHostRequest) (*PublicProfileHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicProfileHost not implemented")
}
func (UnimplementedUserProfileServer) AddToChatList(context.Context, *AddChatListRequest) (*AddChatListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToChatList not implemented")
}
func (UnimplementedUserProfileServer) GetChatList(context.Context, *ChatListRequest) (*ChatListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatList not implemented")
}
func (UnimplementedUserProfileServer) mustEmbedUnimplementedUserProfileServer() {}

// UnsafeUserProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProfileServer will
// result in compilation errors.
type UnsafeUserProfileServer interface {
	mustEmbedUnimplementedUserProfileServer()
}

func RegisterUserProfileServer(s grpc.ServiceRegistrar, srv UserProfileServer) {
	s.RegisterService(&UserProfile_ServiceDesc, srv)
}

func _UserProfile_ProfileDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).ProfileDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfile/ProfileDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).ProfileDetails(ctx, req.(*ProfileDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfile/EditProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).EditProfile(ctx, req.(*EditProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfile/UploadImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).UploadImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_PublicProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).PublicProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfile/PublicProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).PublicProfile(ctx, req.(*PublicProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_PublicProfileHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicProfileHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).PublicProfileHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfile/PublicProfileHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).PublicProfileHost(ctx, req.(*PublicProfileHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_AddToChatList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChatListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).AddToChatList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfile/AddToChatList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).AddToChatList(ctx, req.(*AddChatListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_GetChatList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).GetChatList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfile/GetChatList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).GetChatList(ctx, req.(*ChatListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserProfile_ServiceDesc is the grpc.ServiceDesc for UserProfile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProfile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserProfile",
	HandlerType: (*UserProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfileDetails",
			Handler:    _UserProfile_ProfileDetails_Handler,
		},
		{
			MethodName: "EditProfile",
			Handler:    _UserProfile_EditProfile_Handler,
		},
		{
			MethodName: "UploadImage",
			Handler:    _UserProfile_UploadImage_Handler,
		},
		{
			MethodName: "PublicProfile",
			Handler:    _UserProfile_PublicProfile_Handler,
		},
		{
			MethodName: "PublicProfileHost",
			Handler:    _UserProfile_PublicProfileHost_Handler,
		},
		{
			MethodName: "AddToChatList",
			Handler:    _UserProfile_AddToChatList_Handler,
		},
		{
			MethodName: "GetChatList",
			Handler:    _UserProfile_GetChatList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userprofile.proto",
}
