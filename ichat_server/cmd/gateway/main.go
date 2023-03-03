package main

import (
	"github.com/sirupsen/logrus"
	"ichat"
	"ichat/cmd/gateway/handler"
	"ichat/internal/websocket"
	"ichat/pkg/db"
	"time"
)

func main() {
	r()
}

func r() {

	envPath := "./conf.yaml"

	ichat.InitConf(envPath)

	db.InitMysql(ichat.MysqlConf)
	db.InitRedis(ichat.RedisConf)

	handler := handler.HandlerImpl{}
	ws := websocket.NewServer(envPath)
	ws.Acceptor = &handler
	ws.MessageListener = &handler
	ws.StateListener = &handler
	ws.BeforeAcceptor = &handler
	if err := ws.Run(); err != nil {
		logrus.Error(err)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}
