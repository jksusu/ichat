package ichat_ws

import (
	"encoding/json"
	"sync/atomic"
)

type Status struct {

	//在线数量
	OnlinesConnectNumber int64 `json:"onlinesConnectNumber"`

	//分发消息数量
	ReadMessageNumber  int64 `json:"readMessageNumber"`
	WriteMessageNumber int64 `json:"writeMessageNumber"`

	//失败数量
	ReadMessageFailNumber  int64 `json:"readMessageFailNumber"`
	WriteMessageFailNumber int64 `json:"writeMessageFailNumber"`
}

var GlobalStatus *Status

func InitStatusManager() (err error) {
	GlobalStatus = &Status{}

	return
}

func OnlinesConnectNumber_INC() {
	atomic.AddInt64(&GlobalStatus.OnlinesConnectNumber, 1)
}

func OnlinesConnectNumber_DEC() {
	atomic.AddInt64(&GlobalStatus.OnlinesConnectNumber, -1)
}

func ReadMessageNumber_INC() {
	atomic.AddInt64(&GlobalStatus.ReadMessageNumber, 1)
}

func ReadMessageNumber_DEC() {
	atomic.AddInt64(&GlobalStatus.ReadMessageNumber, -1)
}
func WriteMessageNumber_INC() {
	atomic.AddInt64(&GlobalStatus.WriteMessageNumber, 1)
}

func WriteMessageNumber_DEC() {
	atomic.AddInt64(&GlobalStatus.WriteMessageNumber, -1)
}
func ReadMessageFailNumber_INC() {
	atomic.AddInt64(&GlobalStatus.ReadMessageFailNumber, 1)
}

func ReadMessageFailNumber_DEC() {
	atomic.AddInt64(&GlobalStatus.ReadMessageFailNumber, -1)
}
func WriteMessageFailNumber_INC() {
	atomic.AddInt64(&GlobalStatus.WriteMessageFailNumber, 1)
}

func WriteMessageFailNumber_DEC() {
	atomic.AddInt64(&GlobalStatus.WriteMessageFailNumber, -1)
}

// 输出数据
func (s *Status) Export() (data []byte, err error) {
	return json.Marshal(GlobalStatus)
}
