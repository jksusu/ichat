package main

import (
	"github.com/sirupsen/logrus"
	"ichat"
	"ichat/gateway/handler"
	"ichat/pkg/db"
	"ichat/pkg/ichat_model"
	"ichat/pkg/ichat_tool/idgen"
	"ichat/pkg/ichat_websocket"
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
	ws := ichat_websocket.NewServer(envPath)
	ws.Acceptor = &handler
	ws.MessageListener = &handler
	ws.StateListener = &handler
	ws.BeforeAcceptor = &handler
	if err := ws.Run(); err != nil {
		logrus.Error(err)
	}
	//初始化雪花id生成
	idgen.NewIDGenerator(1)

	db.InitMysql(ichat.MysqlConf)
	db.InitRedis(ichat.RedisConf)
	//初始化ichat模型
	ichat_model.DB = db.DB
	for {
		time.Sleep(1 * time.Second)
	}
}
