package ichat

import (
	"github.com/spf13/viper"
	"ichat/pkg/ichat_ws"
	"ichat/pkg/model"
	"ichat/pkg/protocol/pb"
	"path/filepath"
)

var GlobalConf = new(IchatConf)

type IchatConf struct {
	MysqlConf *MysqlConfig
	RedisConf *RedisConfig
	WsConf    *ichat_ws.Config
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Charset  string `json:"charset"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type GrpcConfig struct {
	LogicListenAddr    string `json:"logicListenAddr"`
	MessageListenAddr  string `json:"messageListenAddr"`
	RelationListenAddr string `json:"relationListenAddr"`
	GatewayListenAddr  string `json:"gatewayListenAddr"`
}

type GrpcClient struct {
	Logic    pb.LogicClient
	Relation pb.RelationClient
	Gateway  pb.GatewayClient
}

func (c *IchatConf) InitConf(path string) error {
	viper := viper.New()
	file, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	viper.SetConfigFile(file)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	allConf := struct {
		Mysql *MysqlConfig
		Redis *RedisConfig
		Grpc  *GrpcConfig
	}{}
	if err = viper.Unmarshal(&allConf); err != nil {
		return err
	}
	c.MysqlConf = allConf.Mysql
	c.RedisConf = allConf.Redis
	return err
}

func InitDb() {
	model.InitMysql(GlobalConf.MysqlConf)
	model.InitMysql(GlobalConf.MysqlConf)
}
