package ichat

import (
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	MysqlConf *MysqlConfig
	RedisConf *RedisConfig
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

func InitConf(filePath string) error {
	viper := viper.New()
	path, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	allConf := struct {
		Mysql *MysqlConfig
		Redis *RedisConfig
	}{}
	if err := viper.Unmarshal(&allConf); err != nil {
		return err
	}
	MysqlConf = allConf.Mysql
	RedisConf = allConf.Redis

	return nil
}
