package ichat

import (
	"flag"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"ichat/pkg/cache"
	"ichat/pkg/ichat_ws"
	"ichat/pkg/model"
	"ichat/pkg/protocol/pb"
	"ichat/pkg/tools/idgen"
)

var GlobalConf *Conf

type Conf struct {
	Mysql   *model.MysqlConf
	Redis   *cache.RedisConf
	Gateway *ichat_ws.Config
	Tcp     *ichat_ws.Config
}

type GrpcClient struct {
	Logic    pb.LogicClient
	Relation pb.RelationClient
	Gateway  pb.GatewayClient
}

func Init() (err error) {
	flag.Parse()
	defer glog.Flush()

	viper := viper.New()
	viper.SetConfigFile("../../conf.yaml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	viper.Unmarshal(&GlobalConf)

	//雪花id生成器
	idgen.NewIDGenerator(1)
	err = model.InitMysql(GlobalConf.Mysql)
	err = cache.InitRedis(GlobalConf.Redis)
	return err
}

func InitGrpc() {
	/*var (
		listen net.Listener
		err    error
		server = grpc.NewServer()
	)
	pb.RegisterGatewayServer(server, &api.GatewayServer{})
	if listen, err = net.Listen("tcp", GrpcConf.Push); err != nil {
		panic(err)
	}
	glog.Info("grpc server listening at %s", GrpcConf.LogicListenAddr)
	go func() {
		if err = server.Serve(listen); err != nil {
			glog.Fatalf("failed to serve: %v", err)
		}
	}()*/
}
