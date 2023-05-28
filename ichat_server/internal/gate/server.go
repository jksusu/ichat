package gate

import (
	"ichat"
	"ichat/pkg/ichat_ws"
)

func NewWsServer() *ichat_ws.Server {
	ws := ichat_ws.NewServer()
	ws.Config = ichat.GlobalConf.Gateway
	ws.Acceptor = &HandlerImpl{}
	ws.MessageListener = &HandlerImpl{}
	ws.StateListener = &HandlerImpl{}
	ws.BeforeAcceptor = &HandlerImpl{}
	return ws
}
