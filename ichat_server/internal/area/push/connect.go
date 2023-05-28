package push

import (
	"ichat/internal/format"
	"ichat/internal/gate"
	"ichat/pkg/cache/connect"
)

type push struct {
	Err error
}

// 获取Fd连接节点信息
func (g *push) getNode() {

}

func Notice(r *format.R) (err error) {
	var connId string
	r.Type = format.NOTICE
	if connId, err = connect.GetConnIdByUid(r.To); err != nil {
		return
	}
	return gate.Gateway.Send(connId, r)
}

func Resp(r *format.R) (err error) {
	var connId string
	r.Type = format.RESPONSE
	if connId, err = connect.GetConnIdByUid(r.From); err != nil {
		return
	}
	return gate.Gateway.Send(connId, r)
}
