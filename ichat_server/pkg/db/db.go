package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"ichat"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func init() {
	InitMysql(ichat.MysqlConf)
	InitRedis(ichat.RedisConf)
}

func InitMysql(conf *ichat.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		panic(err)
	} else {
		logrus.Infof("mysql database:%s connection succeeded", conf.Database)
	}
	return
}

func InitRedis(conf *ichat.RedisConfig) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Database,
	})
	logrus.Infof("redis database:%s:%d connection succeeded", conf.Host, conf.Port)
}
