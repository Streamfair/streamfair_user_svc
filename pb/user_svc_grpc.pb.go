// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	user "github.com/Streamfair/streamfair_user_svc/pb/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserManagementServiceClient is the client API for UserManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementServiceClient interface {
	CreateUser(ctx context.Context, in *user.CreateUserRequest, opts ...grpc.CallOption) (*user.CreateUserResponse, error)
	LoginUser(ctx context.Context, in *user.LoginUserRequest, opts ...grpc.CallOption) (*user.LoginUserResponse, error)
}

type userManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementServiceClient(cc grpc.ClientConnInterface) UserManagementServiceClient {
	return &userManagementServiceClient{cc}
}

func (c *userManagementServiceClient) CreateUser(ctx context.Context, in *user.CreateUserRequest, opts ...grpc.CallOption) (*user.CreateUserResponse, error) {
	out := new(user.CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserManagementService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) LoginUser(ctx context.Context, in *user.LoginUserRequest, opts ...grpc.CallOption) (*user.LoginUserResponse, error) {
	out := new(user.LoginUserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserManagementService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServiceServer is the server API for UserManagementService service.
// All implementations must embed UnimplementedUserManagementServiceServer
// for forward compatibility
type UserManagementServiceServer interface {
	CreateUser(context.Context, *user.CreateUserRequest) (*user.CreateUserResponse, error)
	LoginUser(context.Context, *user.LoginUserRequest) (*user.LoginUserResponse, error)
	mustEmbedUnimplementedUserManagementServiceServer()
}

// UnimplementedUserManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServiceServer struct {
}

func (UnimplementedUserManagementServiceServer) CreateUser(context.Context, *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserManagementServiceServer) LoginUser(context.Context, *user.LoginUserRequest) (*user.LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUserManagementServiceServer) mustEmbedUnimplementedUserManagementServiceServer() {}

// UnsafeUserManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServiceServer will
// result in compilation errors.
type UnsafeUserManagementServiceServer interface {
	mustEmbedUnimplementedUserManagementServiceServer()
}

func RegisterUserManagementServiceServer(s grpc.ServiceRegistrar, srv UserManagementServiceServer) {
	s.RegisterService(&UserManagementService_ServiceDesc, srv)
}

func _UserManagementService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserManagementService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).CreateUser(ctx, req.(*user.CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserManagementService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).LoginUser(ctx, req.(*user.LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagementService_ServiceDesc is the grpc.ServiceDesc for UserManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserManagementService",
	HandlerType: (*UserManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserManagementService_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserManagementService_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_svc.proto",
}
