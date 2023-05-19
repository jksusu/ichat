package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ichat"
	"ichat/internal/api"
	"ichat/internal/gate/handler"
	"ichat/pkg/ichat_websocket"
	"ichat/pkg/protocol/pb"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()
	r()
}

func r() {
	envPath := "/conf.yaml"
	ichat.InitConf(envPath)
	ws := ichat_websocket.NewServer(envPath)
	ws.Acceptor = &handler.HandlerImpl{}
	ws.MessageListener = &handler.HandlerImpl{}
	ws.StateListener = &handler.HandlerImpl{}
	ws.BeforeAcceptor = &handler.HandlerImpl{}
	if err := ws.Run(); err != nil {
		logrus.Error(err)
	}
	server := grpc.NewServer()
	pb.RegisterGatewayServer(server, &api.GatewayServer{})
	listen, err := net.Listen("tcp", ichat.GrpcConf.GatewayListenAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("grpc server listening at %v", ichat.GrpcConf.LogicListenAddr)
	go func() {
		if err = server.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	//初始化雪花id生成
	//idgen.NewIDGenerator(1)
	//
	//db.InitMysql(ichat.MysqlConf)
	//db.InitRedis(ichat.RedisConf)
	////初始化ichat模型
	//ichat_model.DB = db.DB
	for {
		time.Sleep(1 * time.Second)
	}
}
