package gate

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"ichat/internal/format"
	"ichat/pkg/cache/connect"
	"ichat/pkg/ichat_ws"
)

type gate struct {
	Conn *ichat_ws.Connect

	Err error
}

var Gateway = new(gate)

func (g *gate) Notice(r *format.R) *gate {
	var (
		connId string
		err    error
	)
	r.Type = format.NOTICE
	if connId, err = connect.GetConnIdByUid(r.To); err != nil {
		return g
	}
	g.Err = g.send(connId, r)
	return g
}

func (g *gate) Resp(r *format.R) *gate {
	var (
		connId string
		err    error
	)
	r.Type = format.RESPONSE
	if connId, err = connect.GetConnIdByUid(r.From); err != nil {
		return g
	}
	g.Err = g.send(connId, r)
	return g
}

func (g *gate) send(connId string, r *format.R) (err error) {
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