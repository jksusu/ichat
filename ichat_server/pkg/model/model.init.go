package model

import (
	"fmt"
	"github.com/golang/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type MysqlConf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Charset  string `json:"charset"`
}

var DB *gorm.DB

func InitMysql(conf *MysqlConf) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		glog.Fatalln(err)
	}
	glog.Infof("\u001B[1;31;47mmysql database:%s connection success\u001B[0m", conf.Database)
	return
}

// 未读消息列表
type UnreadMessageList struct {
	MsgId     int64     `json:"msg_id"`
	UID       string    `json:"uid"`
	MsgType   int       `json:"msg_type"`
	CreatedAt time.Time `json:"created_at"`
}
