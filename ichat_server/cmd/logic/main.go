package main

import (
	"google.golang.org/grpc"
	"ichat"
	"ichat/internal/api"
	"ichat/pkg/protocol/pb"
	"log"
	"net"
	"time"
)

func main() {
	ListenLogic()
	ListenMessage()
	ListenRelation()
	RunClient()
	for {
		time.Sleep(1 * time.Second)
	}
}
func ListenLogic() {
	server := grpc.NewServer()
	pb.RegisterLogicServer(server, &api.LogicServer{})
	listen, err := net.Listen("tcp", ichat.GrpcConf.LogicListenAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at %v", ichat.GrpcConf.LogicListenAddr)
	go func() {
		if err = server.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func ListenMessage() {
	//server := grpc.NewServer()
	//pb.RegisterMessageServer(server, &api.MessageServer{})
	//listen, err := net.Listen("tcp", ichat.GrpcConf.MessageListenAddr)
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("server listening at %v", ichat.GrpcConf.MessageListenAddr)
	//go func() {
	//	if err = server.Serve(listen); err != nil {
	//		log.Fatalf("failed to serve: %v", err)
	//	}
	//}()
}

func ListenRelation() {
	server := grpc.NewServer()
	pb.RegisterRelationServer(server, &api.RelationServer{})
	listen, err := net.Listen("tcp", ichat.GrpcConf.RelationListenAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at %v", ichat.GrpcConf.RelationListenAddr)
	go func() {
		if err = server.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func RunClient() {
	lconn, err := grpc.Dial(ichat.GrpcConf.LogicListenAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ichat.RpcClient.Logic = pb.NewLogicClient(lconn)

	mconn, err := grpc.Dial(ichat.GrpcConf.MessageListenAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ichat.RpcClient.Message = pb.NewMessageClient(mconn)

	rconn, err := grpc.Dial(ichat.GrpcConf.RelationListenAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ichat.RpcClient.Relation = pb.NewRelationClient(rconn)
}
