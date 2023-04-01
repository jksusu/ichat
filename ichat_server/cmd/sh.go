package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ichat"
	"ichat/cmd/gateway/handler"
	"ichat/cmd/gateway/store"
	"ichat/cmd/http/app/router"
	"ichat/internal/websocket"
	"ichat/pkg/db"
	"ichat/pkg/ichat_model"
	"net/http"
	"time"
)

func main() {
	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	go httpServer()
	runGateway()
	//app := &cli.App{
	//	Name:  "ichat start",
	//	Usage: "启动服务",
	//	Action: func(*cli.Context) error {
	//		fmt.Println("boom! I say!")
	//		go runGateway()
	//		go httpServer()
	//		return nil
	//	},
	//}
	//if err := app.Run(os.Args); err != nil {
	//	log.Fatal(err)
	//}
}

func runGateway() {
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

func httpServer() {
	r := gin.Default()
	ichat.InitConf("./conf.yaml")
	db.InitMysql(ichat.MysqlConf)
	db.InitRedis(ichat.RedisConf)
	//初始化路由
	router.Routers(r)
	//初始化ichat模型
	ichat_model.DB = db.DB
	r.Run("127.0.0.1:80")
}
