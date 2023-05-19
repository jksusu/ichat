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

// LogicClient is the client API for Logic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogicClient interface {
	MessageSync(ctx context.Context, in *MessageSyncReq, opts ...grpc.CallOption) (*MessageSyncResp, error)
	MessageRecall(ctx context.Context, in *MessageRecallReq, opts ...grpc.CallOption) (*MessageRecallResp, error)
	MessageAck(ctx context.Context, in *MessageAckReq, opts ...grpc.CallOption) (*MessageAckResp, error)
	TalkToUser(ctx context.Context, in *TalkToUserReq, opts ...grpc.CallOption) (*TalkToUserResp, error)
	TalkToGroup(ctx context.Context, in *TalkToGroupReq, opts ...grpc.CallOption) (*TalkToGroupResp, error)
	TalkToRoom(ctx context.Context, in *TalkToRoomReq, opts ...grpc.CallOption) (*TalkToRoomResp, error)
	// 通知
	NoticeToUser(ctx context.Context, in *NoticeToUserReq, opts ...grpc.CallOption) (*NoticeToUserResp, error)
	NoticeToGroup(ctx context.Context, in *NoticeToGroupReq, opts ...grpc.CallOption) (*NoticeToGroupResp, error)
	// 会话
	SessionListGet(ctx context.Context, in *SessionListGetReq, opts ...grpc.CallOption) (*SessionListGetResp, error)
	SessionRemove(ctx context.Context, in *SessionRemoveReq, opts ...grpc.CallOption) (*SessionRemoveResp, error)
	SessionSetting(ctx context.Context, in *SessionSettingReq, opts ...grpc.CallOption) (*SessionSettingResp, error)
}

type logicClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicClient(cc grpc.ClientConnInterface) LogicClient {
	return &logicClient{cc}
}

func (c *logicClient) MessageSync(ctx context.Context, in *MessageSyncReq, opts ...grpc.CallOption) (*MessageSyncResp, error) {
	out := new(MessageSyncResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/MessageSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) MessageRecall(ctx context.Context, in *MessageRecallReq, opts ...grpc.CallOption) (*MessageRecallResp, error) {
	out := new(MessageRecallResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/MessageRecall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) MessageAck(ctx context.Context, in *MessageAckReq, opts ...grpc.CallOption) (*MessageAckResp, error) {
	out := new(MessageAckResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/MessageAck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) TalkToUser(ctx context.Context, in *TalkToUserReq, opts ...grpc.CallOption) (*TalkToUserResp, error) {
	out := new(TalkToUserResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/TalkToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) TalkToGroup(ctx context.Context, in *TalkToGroupReq, opts ...grpc.CallOption) (*TalkToGroupResp, error) {
	out := new(TalkToGroupResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/TalkToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) TalkToRoom(ctx context.Context, in *TalkToRoomReq, opts ...grpc.CallOption) (*TalkToRoomResp, error) {
	out := new(TalkToRoomResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/TalkToRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) NoticeToUser(ctx context.Context, in *NoticeToUserReq, opts ...grpc.CallOption) (*NoticeToUserResp, error) {
	out := new(NoticeToUserResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/NoticeToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) NoticeToGroup(ctx context.Context, in *NoticeToGroupReq, opts ...grpc.CallOption) (*NoticeToGroupResp, error) {
	out := new(NoticeToGroupResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/NoticeToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) SessionListGet(ctx context.Context, in *SessionListGetReq, opts ...grpc.CallOption) (*SessionListGetResp, error) {
	out := new(SessionListGetResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/SessionListGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) SessionRemove(ctx context.Context, in *SessionRemoveReq, opts ...grpc.CallOption) (*SessionRemoveResp, error) {
	out := new(SessionRemoveResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/SessionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) SessionSetting(ctx context.Context, in *SessionSettingReq, opts ...grpc.CallOption) (*SessionSettingResp, error) {
	out := new(SessionSettingResp)
	err := c.cc.Invoke(ctx, "/pb.Logic/SessionSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServer is the server API for Logic service.
// All implementations must embed UnimplementedLogicServer
// for forward compatibility
type LogicServer interface {
	MessageSync(context.Context, *MessageSyncReq) (*MessageSyncResp, error)
	MessageRecall(context.Context, *MessageRecallReq) (*MessageRecallResp, error)
	MessageAck(context.Context, *MessageAckReq) (*MessageAckResp, error)
	TalkToUser(context.Context, *TalkToUserReq) (*TalkToUserResp, error)
	TalkToGroup(context.Context, *TalkToGroupReq) (*TalkToGroupResp, error)
	TalkToRoom(context.Context, *TalkToRoomReq) (*TalkToRoomResp, error)
	// 通知
	NoticeToUser(context.Context, *NoticeToUserReq) (*NoticeToUserResp, error)
	NoticeToGroup(context.Context, *NoticeToGroupReq) (*NoticeToGroupResp, error)
	// 会话
	SessionListGet(context.Context, *SessionListGetReq) (*SessionListGetResp, error)
	SessionRemove(context.Context, *SessionRemoveReq) (*SessionRemoveResp, error)
	SessionSetting(context.Context, *SessionSettingReq) (*SessionSettingResp, error)
	mustEmbedUnimplementedLogicServer()
}

// UnimplementedLogicServer must be embedded to have forward compatible implementations.
type UnimplementedLogicServer struct {
}

func (UnimplementedLogicServer) MessageSync(context.Context, *MessageSyncReq) (*MessageSyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageSync not implemented")
}
func (UnimplementedLogicServer) MessageRecall(context.Context, *MessageRecallReq) (*MessageRecallResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageRecall not implemented")
}
func (UnimplementedLogicServer) MessageAck(context.Context, *MessageAckReq) (*MessageAckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageAck not implemented")
}
func (UnimplementedLogicServer) TalkToUser(context.Context, *TalkToUserReq) (*TalkToUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TalkToUser not implemented")
}
func (UnimplementedLogicServer) TalkToGroup(context.Context, *TalkToGroupReq) (*TalkToGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TalkToGroup not implemented")
}
func (UnimplementedLogicServer) TalkToRoom(context.Context, *TalkToRoomReq) (*TalkToRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TalkToRoom not implemented")
}
func (UnimplementedLogicServer) NoticeToUser(context.Context, *NoticeToUserReq) (*NoticeToUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoticeToUser not implemented")
}
func (UnimplementedLogicServer) NoticeToGroup(context.Context, *NoticeToGroupReq) (*NoticeToGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoticeToGroup not implemented")
}
func (UnimplementedLogicServer) SessionListGet(context.Context, *SessionListGetReq) (*SessionListGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionListGet not implemented")
}
func (UnimplementedLogicServer) SessionRemove(context.Context, *SessionRemoveReq) (*SessionRemoveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionRemove not implemented")
}
func (UnimplementedLogicServer) SessionSetting(context.Context, *SessionSettingReq) (*SessionSettingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionSetting not implemented")
}
func (UnimplementedLogicServer) mustEmbedUnimplementedLogicServer() {}

// UnsafeLogicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogicServer will
// result in compilation errors.
type UnsafeLogicServer interface {
	mustEmbedUnimplementedLogicServer()
}

func RegisterLogicServer(s grpc.ServiceRegistrar, srv LogicServer) {
	s.RegisterService(&Logic_ServiceDesc, srv)
}

func _Logic_MessageSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageSyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).MessageSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/MessageSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).MessageSync(ctx, req.(*MessageSyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_MessageRecall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRecallReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).MessageRecall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/MessageRecall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).MessageRecall(ctx, req.(*MessageRecallReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_MessageAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageAckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).MessageAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/MessageAck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).MessageAck(ctx, req.(*MessageAckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_TalkToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TalkToUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).TalkToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/TalkToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).TalkToUser(ctx, req.(*TalkToUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_TalkToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TalkToGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).TalkToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/TalkToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).TalkToGroup(ctx, req.(*TalkToGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_TalkToRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TalkToRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).TalkToRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/TalkToRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).TalkToRoom(ctx, req.(*TalkToRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_NoticeToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoticeToUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).NoticeToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/NoticeToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).NoticeToUser(ctx, req.(*NoticeToUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_NoticeToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoticeToGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).NoticeToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/NoticeToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).NoticeToGroup(ctx, req.(*NoticeToGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_SessionListGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionListGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).SessionListGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/SessionListGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).SessionListGet(ctx, req.(*SessionListGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_SessionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).SessionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/SessionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).SessionRemove(ctx, req.(*SessionRemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_SessionSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionSettingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).SessionSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Logic/SessionSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).SessionSetting(ctx, req.(*SessionSettingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Logic_ServiceDesc is the grpc.ServiceDesc for Logic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Logic",
	HandlerType: (*LogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageSync",
			Handler:    _Logic_MessageSync_Handler,
		},
		{
			MethodName: "MessageRecall",
			Handler:    _Logic_MessageRecall_Handler,
		},
		{
			MethodName: "MessageAck",
			Handler:    _Logic_MessageAck_Handler,
		},
		{
			MethodName: "TalkToUser",
			Handler:    _Logic_TalkToUser_Handler,
		},
		{
			MethodName: "TalkToGroup",
			Handler:    _Logic_TalkToGroup_Handler,
		},
		{
			MethodName: "TalkToRoom",
			Handler:    _Logic_TalkToRoom_Handler,
		},
		{
			MethodName: "NoticeToUser",
			Handler:    _Logic_NoticeToUser_Handler,
		},
		{
			MethodName: "NoticeToGroup",
			Handler:    _Logic_NoticeToGroup_Handler,
		},
		{
			MethodName: "SessionListGet",
			Handler:    _Logic_SessionListGet_Handler,
		},
		{
			MethodName: "SessionRemove",
			Handler:    _Logic_SessionRemove_Handler,
		},
		{
			MethodName: "SessionSetting",
			Handler:    _Logic_SessionSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ichat.logic.proto",
}
