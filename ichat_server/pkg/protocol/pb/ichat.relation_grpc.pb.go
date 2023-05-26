// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// RelationClient is the client API for Relation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationClient interface {
	FriendsApply(ctx context.Context, in *FriendsApplyRequest, opts ...grpc.CallOption) (*FriendsApplyResponse, error)
	FriendsApplyReply(ctx context.Context, in *FriendsApplyReplyRequest, opts ...grpc.CallOption) (*FriendsApplyReplyResponse, error)
	FriendsEdit(ctx context.Context, in *FriendsEditRequest, opts ...grpc.CallOption) (*FriendsEditResponse, error)
	FriendsDel(ctx context.Context, in *FriendsDelRequest, opts ...grpc.CallOption) (*FriendsDelResponse, error)
	GroupCreate(ctx context.Context, in *GroupCreateRequest, opts ...grpc.CallOption) (*GroupCreateResponse, error)
	GroupEdit(ctx context.Context, in *GroupEditRequest, opts ...grpc.CallOption) (*GroupEditResponse, error)
	GroupJoinApply(ctx context.Context, in *GroupJoinApplyRequest, opts ...grpc.CallOption) (*GroupJoinApplyResponse, error)
	GroupJoinApplyApprove(ctx context.Context, in *GroupJoinApplyApproveRequest, opts ...grpc.CallOption) (*GroupJoinApplyApproveResponse, error)
	GroupExit(ctx context.Context, in *GroupExitRequest, opts ...grpc.CallOption) (*GroupExitResponse, error)
	GroupMemberRemove(ctx context.Context, in *GroupMemberRemoveRequest, opts ...grpc.CallOption) (*GroupMemberRemoveResponse, error)
	GroupMemberSetUp(ctx context.Context, in *GroupMemberSetUpRequest, opts ...grpc.CallOption) (*GroupMemberSetUpResponse, error)
}

type relationClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationClient(cc grpc.ClientConnInterface) RelationClient {
	return &relationClient{cc}
}

func (c *relationClient) FriendsApply(ctx context.Context, in *FriendsApplyRequest, opts ...grpc.CallOption) (*FriendsApplyResponse, error) {
	out := new(FriendsApplyResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/FriendsApply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) FriendsApplyReply(ctx context.Context, in *FriendsApplyReplyRequest, opts ...grpc.CallOption) (*FriendsApplyReplyResponse, error) {
	out := new(FriendsApplyReplyResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/FriendsApplyReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) FriendsEdit(ctx context.Context, in *FriendsEditRequest, opts ...grpc.CallOption) (*FriendsEditResponse, error) {
	out := new(FriendsEditResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/FriendsEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) FriendsDel(ctx context.Context, in *FriendsDelRequest, opts ...grpc.CallOption) (*FriendsDelResponse, error) {
	out := new(FriendsDelResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/FriendsDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GroupCreate(ctx context.Context, in *GroupCreateRequest, opts ...grpc.CallOption) (*GroupCreateResponse, error) {
	out := new(GroupCreateResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/GroupCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GroupEdit(ctx context.Context, in *GroupEditRequest, opts ...grpc.CallOption) (*GroupEditResponse, error) {
	out := new(GroupEditResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/GroupEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GroupJoinApply(ctx context.Context, in *GroupJoinApplyRequest, opts ...grpc.CallOption) (*GroupJoinApplyResponse, error) {
	out := new(GroupJoinApplyResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/GroupJoinApply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GroupJoinApplyApprove(ctx context.Context, in *GroupJoinApplyApproveRequest, opts ...grpc.CallOption) (*GroupJoinApplyApproveResponse, error) {
	out := new(GroupJoinApplyApproveResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/GroupJoinApplyApprove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GroupExit(ctx context.Context, in *GroupExitRequest, opts ...grpc.CallOption) (*GroupExitResponse, error) {
	out := new(GroupExitResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/GroupExit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GroupMemberRemove(ctx context.Context, in *GroupMemberRemoveRequest, opts ...grpc.CallOption) (*GroupMemberRemoveResponse, error) {
	out := new(GroupMemberRemoveResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/GroupMemberRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GroupMemberSetUp(ctx context.Context, in *GroupMemberSetUpRequest, opts ...grpc.CallOption) (*GroupMemberSetUpResponse, error) {
	out := new(GroupMemberSetUpResponse)
	err := c.cc.Invoke(ctx, "/pb.Relation/GroupMemberSetUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServer is the server API for Relation service.
// All implementations must embed UnimplementedRelationServer
// for forward compatibility
type RelationServer interface {
	FriendsApply(context.Context, *FriendsApplyRequest) (*FriendsApplyResponse, error)
	FriendsApplyReply(context.Context, *FriendsApplyReplyRequest) (*FriendsApplyReplyResponse, error)
	FriendsEdit(context.Context, *FriendsEditRequest) (*FriendsEditResponse, error)
	FriendsDel(context.Context, *FriendsDelRequest) (*FriendsDelResponse, error)
	GroupCreate(context.Context, *GroupCreateRequest) (*GroupCreateResponse, error)
	GroupEdit(context.Context, *GroupEditRequest) (*GroupEditResponse, error)
	GroupJoinApply(context.Context, *GroupJoinApplyRequest) (*GroupJoinApplyResponse, error)
	GroupJoinApplyApprove(context.Context, *GroupJoinApplyApproveRequest) (*GroupJoinApplyApproveResponse, error)
	GroupExit(context.Context, *GroupExitRequest) (*GroupExitResponse, error)
	GroupMemberRemove(context.Context, *GroupMemberRemoveRequest) (*GroupMemberRemoveResponse, error)
	GroupMemberSetUp(context.Context, *GroupMemberSetUpRequest) (*GroupMemberSetUpResponse, error)
	mustEmbedUnimplementedRelationServer()
}

// UnimplementedRelationServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServer struct {
}

func (UnimplementedRelationServer) FriendsApply(context.Context, *FriendsApplyRequest) (*FriendsApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendsApply not implemented")
}
func (UnimplementedRelationServer) FriendsApplyReply(context.Context, *FriendsApplyReplyRequest) (*FriendsApplyReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendsApplyReply not implemented")
}
func (UnimplementedRelationServer) FriendsEdit(context.Context, *FriendsEditRequest) (*FriendsEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendsEdit not implemented")
}
func (UnimplementedRelationServer) FriendsDel(context.Context, *FriendsDelRequest) (*FriendsDelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendsDel not implemented")
}
func (UnimplementedRelationServer) GroupCreate(context.Context, *GroupCreateRequest) (*GroupCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupCreate not implemented")
}
func (UnimplementedRelationServer) GroupEdit(context.Context, *GroupEditRequest) (*GroupEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupEdit not implemented")
}
func (UnimplementedRelationServer) GroupJoinApply(context.Context, *GroupJoinApplyRequest) (*GroupJoinApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupJoinApply not implemented")
}
func (UnimplementedRelationServer) GroupJoinApplyApprove(context.Context, *GroupJoinApplyApproveRequest) (*GroupJoinApplyApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupJoinApplyApprove not implemented")
}
func (UnimplementedRelationServer) GroupExit(context.Context, *GroupExitRequest) (*GroupExitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupExit not implemented")
}
func (UnimplementedRelationServer) GroupMemberRemove(context.Context, *GroupMemberRemoveRequest) (*GroupMemberRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupMemberRemove not implemented")
}
func (UnimplementedRelationServer) GroupMemberSetUp(context.Context, *GroupMemberSetUpRequest) (*GroupMemberSetUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupMemberSetUp not implemented")
}
func (UnimplementedRelationServer) mustEmbedUnimplementedRelationServer() {}

// UnsafeRelationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServer will
// result in compilation errors.
type UnsafeRelationServer interface {
	mustEmbedUnimplementedRelationServer()
}

func RegisterRelationServer(s grpc.ServiceRegistrar, srv RelationServer) {
	s.RegisterService(&Relation_ServiceDesc, srv)
}

func _Relation_FriendsApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendsApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).FriendsApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/FriendsApply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).FriendsApply(ctx, req.(*FriendsApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_FriendsApplyReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendsApplyReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).FriendsApplyReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/FriendsApplyReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).FriendsApplyReply(ctx, req.(*FriendsApplyReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_FriendsEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendsEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).FriendsEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/FriendsEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).FriendsEdit(ctx, req.(*FriendsEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_FriendsDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendsDelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).FriendsDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/FriendsDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).FriendsDel(ctx, req.(*FriendsDelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GroupCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GroupCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/GroupCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GroupCreate(ctx, req.(*GroupCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GroupEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GroupEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/GroupEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GroupEdit(ctx, req.(*GroupEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GroupJoinApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupJoinApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GroupJoinApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/GroupJoinApply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GroupJoinApply(ctx, req.(*GroupJoinApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GroupJoinApplyApprove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupJoinApplyApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GroupJoinApplyApprove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/GroupJoinApplyApprove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GroupJoinApplyApprove(ctx, req.(*GroupJoinApplyApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GroupExit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupExitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GroupExit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/GroupExit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GroupExit(ctx, req.(*GroupExitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GroupMemberRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupMemberRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GroupMemberRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/GroupMemberRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GroupMemberRemove(ctx, req.(*GroupMemberRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GroupMemberSetUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupMemberSetUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GroupMemberSetUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Relation/GroupMemberSetUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GroupMemberSetUp(ctx, req.(*GroupMemberSetUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Relation_ServiceDesc is the grpc.ServiceDesc for Relation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Relation",
	HandlerType: (*RelationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FriendsApply",
			Handler:    _Relation_FriendsApply_Handler,
		},
		{
			MethodName: "FriendsApplyReply",
			Handler:    _Relation_FriendsApplyReply_Handler,
		},
		{
			MethodName: "FriendsEdit",
			Handler:    _Relation_FriendsEdit_Handler,
		},
		{
			MethodName: "FriendsDel",
			Handler:    _Relation_FriendsDel_Handler,
		},
		{
			MethodName: "GroupCreate",
			Handler:    _Relation_GroupCreate_Handler,
		},
		{
			MethodName: "GroupEdit",
			Handler:    _Relation_GroupEdit_Handler,
		},
		{
			MethodName: "GroupJoinApply",
			Handler:    _Relation_GroupJoinApply_Handler,
		},
		{
			MethodName: "GroupJoinApplyApprove",
			Handler:    _Relation_GroupJoinApplyApprove_Handler,
		},
		{
			MethodName: "GroupExit",
			Handler:    _Relation_GroupExit_Handler,
		},
		{
			MethodName: "GroupMemberRemove",
			Handler:    _Relation_GroupMemberRemove_Handler,
		},
		{
			MethodName: "GroupMemberSetUp",
			Handler:    _Relation_GroupMemberSetUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ichat.relation.proto",
}