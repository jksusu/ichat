package main

import (
	"github.com/sirupsen/logrus"
	"ichat"
	"ichat/cmd/gateway/handler"
	"ichat/cmd/gateway/store"
	"ichat/internal/websocket"
	"ichat/pkg/db"
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

	envPath := "./conf.yaml"

	ichat.InitConf(envPath)

	handler := handler.HandlerImpl{}
	ws := websocket.NewServer(envPath)
	ws.Acceptor = &handler
	ws.MessageListener = &handler
	ws.StateListener = &handler
	ws.BeforeAcceptor = &handler
	if err := ws.Run(); err != nil {
		logrus.Error(err)
	}
	//初始化雪花id生成
	store.NewIDGenerator(100)

	db.InitMysql(ichat.MysqlConf)
	db.InitRedis(ichat.RedisConf)

	for {
		time.Sleep(1 * time.Second)
	}
}
