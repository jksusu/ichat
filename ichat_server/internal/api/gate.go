package api

import (
	"context"
	"ichat/pkg/protocol/pb"
)

type GatewayServer struct {
	pb.UnsafeGatewayServer
}

func (*GatewayServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return nil, nil
}

func (*GatewayServer) Response(ctx context.Context, req *pb.ResponseReq) (*pb.Empty, error) {
	return nil, nil
}

func (*GatewayServer) Notice(ctx context.Context, req *pb.NoticeReq) (*pb.Empty, error) {
	return nil, nil
}
