package ichat_websocket

import (
	"encoding/json"
	"path/filepath"

	"github.com/spf13/viper"
)

var GlobalConfig *Config

type Config struct {
	WsServiceId        string `json:"serviceid"`          //服务id
	WsServiceName      string `json:"servicename"`        //服务名
	WsPort             int    `json:"wsport"`             //端口
	WsReadTimeout      int    `json:"wsreadtimeout"`      //毫秒读超时时间
	WsWriteTimeout     int    `json:"wswritetimeout"`     //毫秒写超时时间
	WsWriteChannelSize int    `json:"wswritechannelsize"` //写通道最大数量
	WsReadChannelSize  int    `json:"wsreadchannelsize"`  //读通道最大数量
	WsHeartbeatTimeout int    `json:"wsheartbeattimeout"` //秒心跳超时时间，超过这个时间触发重连
}

func InitConfig(filePath string) error {
	viper := viper.New()
	path, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	viper.SetConfigFile(path)
	viper.ReadInConfig()
	maps := viper.GetStringMap("gate")
	jsonArr, err := json.Marshal(maps)
	if err != nil {
		return err
	}
	conf := &Config{}
	err = json.Unmarshal(jsonArr, &conf)
	if err != nil {
		return err
	}

	GlobalConfig = conf

	return nil
}
