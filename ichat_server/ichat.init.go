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
	Other   struct {
		nodeId int64 `json:"nodeId"`
	}
}

type GrpcClient struct {
	Logic    pb.LogicClient
	Relation pb.RelationClient
	Gateway  pb.GatewayClient
}

func Init() (err error) {
	flag.Lookup("v").Value.Set("0")
	flag.Lookup("stderrthreshold").Value.Set("INFO")
	flag.Lookup("log_dir").Value.Set("../log/")
	flag.Set("logtostderr", "true")
	glog.Flush()

	viper := viper.New()
	viper.SetConfigFile("../conf.yaml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	viper.Unmarshal(&GlobalConf)

	idgen.NewIDGenerator(GlobalConf.Other.nodeId)
	cache.InitRedis(GlobalConf.Redis)
	model.InitMysql(GlobalConf.Mysql)

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
