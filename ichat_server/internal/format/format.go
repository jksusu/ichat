package format

import "encoding/json"

const (
	REQUEST  = 1 //请求
	RESPONSE = 2 //回复
	NOTICE   = 3 //通知
)

type R struct {
	MsgId int64       `json:"msgId"` //消息ID
	ReqId int         `json:"reqId"` //请求ID
	From  string      `json:"from"`  //来自用户 uid 不能为空
	To    string      `json:"to"`    //接收者uid 可为空
	Type  int         `json:"type"`  //请求类型
	Route string      `json:"route"` //路由
	Data  interface{} `json:"data"`  //数据
}

type ErrResp struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

func Decode(data interface{}, v any) (err error) {
	var p []byte
	p, err = json.Marshal(data)
	err = json.Unmarshal(p, v)
	return err
}
