package unread_service

import "time"

type UnReadMessage struct {
	UID       string    `json:"uid"`
	TO        string    `json:"to"`
	MsgId     int64     `json:"msg_id"`
	MsgType   string    `json:"msg_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func AddFriend() {

}
