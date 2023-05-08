package api

import (
	"context"
	"ichat/pkg/protocol/pb"
)

type RelationServer struct {
	pb.UnsafeRelationServer
}

func (*RelationServer) FriendsApply(ctx context.Context, request *pb.FriendsApplyRequest) (*pb.FriendsApplyResponse, error) {
	return nil, nil
}
func (*RelationServer) FriendsApplyReply(ctx context.Context, request *pb.FriendsApplyReplyRequest) (*pb.FriendsApplyReplyResponse, error) {
	return nil, nil
}

func (*RelationServer) FriendsDel(ctx context.Context, request *pb.FriendsDelRequest) (*pb.FriendsDelResponse, error) {
	return nil, nil
}
func (*RelationServer) FriendsEdit(ctx context.Context, request *pb.FriendsEditRequest) (*pb.FriendsEditResponse, error) {
	return nil, nil
}
func (*RelationServer) GroupCreate(ctx context.Context, request *pb.GroupCreateRequest) (*pb.GroupCreateResponse, error) {
	return nil, nil
}
func (*RelationServer) GroupEdit(ctx context.Context, request *pb.GroupEditRequest) (*pb.GroupEditResponse, error) {
	return nil, nil
}
func (*RelationServer) GroupJoinApply(ctx context.Context, request *pb.GroupJoinApplyRequest) (*pb.GroupJoinApplyResponse, error) {
	return nil, nil
}
func (*RelationServer) GroupJoinApplyApprove(ctx context.Context, request *pb.GroupJoinApplyApproveRequest) (*pb.GroupJoinApplyApproveResponse, error) {
	return nil, nil
}
func (*RelationServer) GroupExit(ctx context.Context, request *pb.GroupExitRequest) (*pb.GroupExitResponse, error) {
	return nil, nil
}

func (*RelationServer) GroupMemberRemove(ctx context.Context, request *pb.GroupMemberRemoveRequest) (*pb.GroupMemberRemoveResponse, error) {
	return nil, nil
}

func (*RelationServer) GroupMemberSetUp(ctx context.Context, request *pb.GroupMemberSetUpRequest) (*pb.GroupMemberSetUpResponse, error) {
	return nil, nil
}
