package session

import (
	"ichat/pkg/model"
	"time"
)

type SessionList struct {
	To             string    `json:"to"`
	Nickname       string    `json:"nickname"`
	Username       string    `json:"username"`
	Avatar         string    `json:"avatar"`
	Topping        int       `json:"topping"`
	NoDisturb      int       `json:"no_disturb"`
	SessionType    int       `json:"session_type"`
	LastMsgContent string    `json:"last_msg_content"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func List(uid string) (data []*SessionList, err error) {
	list := model.SessionListModel.GetAll(uid)
	//查询用户详情
	userIds := make([]string, 0)
	groupIds := make([]string, 0)
	for _, v := range list {
		if v.SessionType == model.SESSION_TYPE_CONTACTS {
			userIds = append(userIds, v.Uid)
		}
		if v.SessionType == model.SESSION_TYPE_GROUP {
			groupIds = append(groupIds, v.Uid)
		}
	}
	users := make([]model.Users, 0)
	if len(userIds) > 0 {
		users, err = model.UserModel.FindInIds(userIds)
		if err != nil {
			return nil, err
		}
	}
	for _, v := range list {
		for _, vv := range users {
			if v.To == vv.UID {
				s := &SessionList{
					To:             vv.UID,
					Nickname:       vv.Nickname,
					Username:       vv.Username,
					Avatar:         vv.Avatar,
					Topping:        0,
					NoDisturb:      0,
					SessionType:    v.SessionType,
					LastMsgContent: "",
					UpdatedAt:      v.UpdatedAt,
				}
				data = append(data, s)
			}
		}
	}
	return data, err
}

func Add(uid, to string, sessionType int, msgId int64) error {
	if row, err := model.SessionListModel.Insert(&model.SessionList{
		Uid:         uid,
		To:          to,
		SessionType: sessionType,
		LastMsgId:   msgId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}); row == 0 || err != nil {
		return err
	}
	return nil
}

func Remove(to string) error {
	if row, err := model.SessionListModel.Delete(to); row == 0 || err != nil {
		return err
	}
	return nil
}

func Options(to string, sessionType, topping, noDisturb int) error {
	if sessionType == model.SESSION_TYPE_CONTACTS {
		//修改置顶或者是免打扰
		return nil
	}
	if sessionType == model.SESSION_TYPE_GROUP {
		//修改置顶或者是免打扰
		return nil

	}
	return nil
}
