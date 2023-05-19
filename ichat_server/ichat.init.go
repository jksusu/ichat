package ichat

import (
	"github.com/spf13/viper"
	"ichat/pkg/protocol/pb"
	"path/filepath"
)

var (
	MysqlConf *MysqlConfig
	RedisConf *RedisConfig
	GrpcConf  *GrpcConfig
	RpcClient *GrpcClient
)

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

func init() {
	InitConf("./conf.yaml")
}

func InitConf(filePath string) error {
	viper := viper.New()
	path, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	viper.SetConfigFile(path)
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
	MysqlConf = allConf.Mysql
	RedisConf = allConf.Redis
	GrpcConf = allConf.Grpc

	return nil
}
