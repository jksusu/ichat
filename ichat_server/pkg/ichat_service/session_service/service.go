package session_service

import (
	"github.com/sirupsen/logrus"
	cacheSessionList "ichat/pkg/ichat_cache/session/lists"
	"ichat/pkg/ichat_model"
	"time"
)

type SessionList struct {
	SessionUID   string    `json:"session_uid"`
	Nickname     string    `json:"nickname"`
	LastMsg      string    `json:"last_msg"`
	LastMsgType  int       `json:"last_msg_type"` //最后一条消息类型
	LastTime     time.Time `json:"last_time"`
	HeadPortrait string    `json:"head_portrait"`
	Topping      int       `json:"topping"` //根据数字排序越大排前面
	UnReadNumber int       `json:"un_read_number"`
}

func GetSessionAll(UID string) []*SessionList {
	lists, err := cacheSessionList.GetAll(UID)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	var user ichat_model.Users
	users, err := user.FindInIds(lists)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if len(users) == 0 {
		logrus.Info("data is [] 0")
		return nil
	}

	var l = make([]*SessionList, len(users))

	for i, list := range lists {
		for _, user := range users {
			if list == user.UID {
				l[i] = &SessionList{
					SessionUID:   user.UID,
					Nickname:     user.Nickname,
					HeadPortrait: user.HeadPortrait,
					Topping:      0,
					UnReadNumber: GetUnread(),
				}
			}
		}
	}

	return l
}

func GetUnread() int {
	return 100
}

func GetCountUnread() {

}
