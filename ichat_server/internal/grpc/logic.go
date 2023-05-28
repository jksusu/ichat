package grpc

import (
	"context"
	"ichat/pkg/protocol/pb"
)

type LogicServer struct {
	pb.UnsafeLogicServer
}

func (*LogicServer) MessageSync(ctx context.Context, request *pb.MessageSyncReq) (*pb.MessageSyncResp, error) {
	return nil, nil
}

func (*LogicServer) MessageRecall(ctx context.Context, request *pb.MessageRecallReq) (*pb.MessageRecallResp, error) {
	return nil, nil
}

func (*LogicServer) MessageAck(ctx context.Context, request *pb.MessageAckReq) (*pb.MessageAckResp, error) {
	return nil, nil
}

func (*LogicServer) TalkToUser(ctx context.Context, r *pb.TalkToUserReq) (*pb.TalkToUserResp, error) {
	return nil, nil
}

func (*LogicServer) TalkToGroup(ctx context.Context, request *pb.TalkToGroupReq) (*pb.TalkToGroupResp, error) {
	return nil, nil
}

func (*LogicServer) TalkToRoom(ctx context.Context, request *pb.TalkToRoomReq) (*pb.TalkToRoomResp, error) {
	return nil, nil
}

func (*LogicServer) NoticeToUser(ctx context.Context, request *pb.NoticeToUserReq) (*pb.NoticeToUserResp, error) {
	return nil, nil
}

func (*LogicServer) NoticeToGroup(ctx context.Context, request *pb.NoticeToGroupReq) (*pb.NoticeToGroupResp, error) {
	return nil, nil
}

func (*LogicServer) SessionListGet(ctx context.Context, request *pb.SessionListGetReq) (*pb.SessionListGetResp, error) {
	return nil, nil
}

func (*LogicServer) SessionRemove(ctx context.Context, request *pb.SessionRemoveReq) (*pb.SessionRemoveResp, error) {
	return nil, nil
}
func (*LogicServer) SessionSetting(ctx context.Context, request *pb.SessionSettingReq) (*pb.SessionSettingResp, error) {
	return nil, nil
}
