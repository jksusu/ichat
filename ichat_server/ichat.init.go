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

type (
	Conf struct {
		Mysql   *model.MysqlConf
		Redis   *cache.RedisConf
		Gateway *ichat_ws.Config
		Tcp     *ichat_ws.Config
		Other   struct {
			NodeId   int64  `json:"nodeId"`
			HttpAddr string `json:"httpPort"`
		}
	}

	GrpcClient struct {
		Logic    pb.LogicClient
		Relation pb.RelationClient
		Gateway  pb.GatewayClient
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

func Init() (err error) {
	flag.Lookup("v").Value.Set("0")
	flag.Lookup("stderrthreshold").Value.Set("INFO")
	flag.Lookup("log_dir").Value.Set("../log/")
	flag.Set("logtostderr", "true")
	glog.Flush()

	viper := viper.New()
	viper.SetConfigFile("conf.yaml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	viper.Unmarshal(&GlobalConf)

	idgen.NewIDGenerator(GlobalConf.Other.NodeId)
	cache.InitRedis(GlobalConf.Redis)
	model.InitMysql(GlobalConf.Mysql)

	return err
}
