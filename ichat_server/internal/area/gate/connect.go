package gate

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"ichat/internal/format"
	"ichat/pkg/ichat_cache/connect"
	"ichat/pkg/ichat_websocket"
)

type gate struct {
	Conn *ichat_websocket.Connect

	Err error
}

var Gateway = new(gate)

func (g *gate) Notice(r *format.R) *gate {
	var connId string
	r.Type = format.NOTICE
	connId, g.Err = connect.GetConnIdByUid(r.To)
	if g.Err != nil {
		return g
	}
	g.Err = g.send(connId, r)
	return g
}

func (g *gate) Resp(r *format.R) *gate {
	var connId string
	r.Type = format.RESPONSE
	if connId, g.Err = connect.GetConnIdByUid(r.From); g.Err != nil {
		return g
	}
	g.Err = g.send(connId, r)
	return g
}

// send 私有的所有下发消息都走send
func (g *gate) send(connId string, r *format.R) (err error) {
	buf := new(bytes.Buffer)
	if err = gob.NewEncoder(buf).Encode(r); err != nil {
		return
	}
	var d []byte
	if d, err = json.Marshal(r); err != nil {
		return
	}
	conn, ok := ichat_websocket.GlobalConnManager.GetConn(connId)
	if !ok {
		return errors.New("connection not connId")
	}
	return conn.Push(&ichat_websocket.WSMessage{MsgType: websocket.TextMessage, MsgData: d})
}
