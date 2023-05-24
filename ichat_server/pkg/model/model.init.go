package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"ichat"
	"time"
)

var DB *gorm.DB

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

type GLOBALMODEL struct {
	ID        int            `json:"id" gorm:"primarykey;comment:主键id"`
	CreatedAt time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"comment:删除时间"`
}

// 未读消息列表
type UnreadMessageList struct {
	MsgId     int64     `json:"msg_id"`
	UID       string    `json:"uid"`
	MsgType   int       `json:"msg_type"`
	CreatedAt time.Time `json:"created_at"`
}
