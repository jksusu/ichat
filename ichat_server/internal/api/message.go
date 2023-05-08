package api

import (
	"context"
	"ichat/internal/rpc/message"
	"ichat/pkg/protocol/pb"
)

type MessageServer struct {
	pb.UnsafeMessageServer
}

func (*MessageServer) MessageSync(ctx context.Context, request *pb.MessageSyncRequest) (*pb.MessageSyncResponse, error) {
	return nil, nil
}

func (*MessageServer) MessageRecall(ctx context.Context, request *pb.MessageRecallRequest) (*pb.MessageRecallResponse, error) {
	return nil, nil
}

func (*MessageServer) MessageAck(ctx context.Context, request *pb.MessageAckRequest) (*pb.MessageAckResponse, error) {
	return nil, nil
}

func (*MessageServer) TalkToUser(ctx context.Context, r *pb.TalkToUserRequest) (*pb.TalkToUserResponse, error) {
	return message.Talk.ToUser(ctx, r)
}

func (*MessageServer) TalkToGroup(ctx context.Context, request *pb.TalkToGroupRequest) (*pb.TalkToGroupResponse, error) {
	return nil, nil
}

func (*MessageServer) TalkToRoom(ctx context.Context, request *pb.TalkToRoomRequest) (*pb.TalkToRoomResponse, error) {
	return nil, nil
}

func (*MessageServer) SessionListGet(ctx context.Context, request *pb.SessionListGetRequest) (*pb.SessionListGetResponse, error) {
	return nil, nil
}

func (*MessageServer) SessionRemove(ctx context.Context, request *pb.SessionRemoveRequest) (*pb.SessionRemoveResponse, error) {
	return nil, nil
}
func (*MessageServer) SessionSetting(ctx context.Context, request *pb.SessionSettingRequest) (*pb.SessionSettingResponse, error) {
	return nil, nil
}
