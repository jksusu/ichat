package push

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"ichat/internal/format"
	"ichat/pkg/cache/connect"
	"ichat/pkg/ichat_ws"

	"github.com/gorilla/websocket"
)

type push struct {
	Err error
}

var Push = new(push)

// 获取Fd连接节点信息
func (g *push) getNode() {

}

func (g *push) Notice(r *format.R) *push {
	var connId string
	r.Type = format.NOTICE
	connId, g.Err = connect.GetConnIdByUid(r.To)
	if g.Err != nil {
		return g
	}
	g.Err = g.send(connId, r)
	return g
}

func (g *push) Resp(r *format.R) *push {
	var connId string
	r.Type = format.RESPONSE
	if connId, g.Err = connect.GetConnIdByUid(r.From); g.Err != nil {
		return g
	}
	g.Err = g.send(connId, r)
	return g
}

// 所有下发消息都走send
func (g *push) send(connId string, r *format.R) (err error) {
	buf := new(bytes.Buffer)
	if err = gob.NewEncoder(buf).Encode(r); err != nil {
		return
	}
	var d []byte
	if d, err = json.Marshal(r); err != nil {
		return
	}
	conn, ok := ichat_ws.GlobalConnManager.Get(connId)
	if !ok {
		return errors.New("connection not connId")
	}
	return conn.Push(&ichat_ws.WSMessage{MsgType: websocket.TextMessage, MsgData: d})
}
