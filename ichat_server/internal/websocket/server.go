package websocket

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type BeforeAcceptor interface {
	BeforeAccept(*http.Request) error
}
type Acceptor interface {
	Accept(*Connect, *http.Request) (uint64, error)
}
type MessageListener interface {
	Receive(*Connect, *ChatMessage)
}
type StateListener interface {
	Disconnect(id string) error
}

type Server struct {
	sync.Once
	server *http.Server
	connId string
	BeforeAcceptor
	Acceptor
	MessageListener
	StateListener
}

var (
	GlobalServer *Server
	Upgrader     = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 新的websocket服务
func NewServer(path string) *Server {
	if err := InitConfig(path); err != nil {
		logrus.Error(err)
	}
	if err := InitConnManager(); err != nil {
		logrus.Error(err)
	}
	if err := InitStatusManager(); err != nil {
		logrus.Error(err)
	}
	return &Server{}
}

// 初始化websocket服务
func (s *Server) Run() (err error) {
	if s.Acceptor == nil || s.MessageListener == nil || s.StateListener == nil || s.BeforeAcceptor == nil {
		return errors.New("Must be achieved Acceptor MessageListener StateListener BeforeAcceptor interface")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", s.handleConnect)
	s.server = &http.Server{
		ReadTimeout:  time.Duration(GlobalConfig.WsReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(GlobalConfig.WsWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}
	var listen net.Listener
	if listen, err = net.Listen("tcp", fmt.Sprintf(":%d", GlobalConfig.WsPort)); err != nil {
		return
	}
	GlobalServer = s
	printStart()
	go s.server.Serve(listen)
	return
}

func (s *Server) handleConnect(w http.ResponseWriter, r *http.Request) {
	if err := s.BeforeAccept(r); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	websocket, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	id := fmt.Sprintf("%d%s", time.Now().UnixNano(), "b")

	conn := HttpUpgraderToWebsocket(id, websocket)

	s.Accept(conn, r)

	s.handler(conn)
}

func (s *Server) handler(c *Connect) {
	for {
		data, err := c.ReadMsg()
		if err != nil {
			goto CLOSE
		}
		if data == nil || data.MsgType != 1 {
			continue
		}
		msg, err := DecodeMessageBody(data.MsgData)
		if err != nil {
			goto CLOSE
		}

		switch msg.Type {
		case PINGTYPE:
			if err := c.Pong(); err != nil {
				logrus.Error(err)
				goto CLOSE
			}
		case CHATTYPE:
			if msg.Data != nil {
				s.Receive(c, msg.Data)
			}
		}
	}
CLOSE:
	c.Close()
	return
}

func printStart() {

	listenAddr := fmt.Sprintf(":%d", GlobalConfig.WsPort)
	serviceId := GlobalConfig.WsServiceId
	serviceName := GlobalConfig.WsServiceName

	fmt.Printf("\033[1;31;47m %s \033[0m\n", "gateway websocket server start success")
	fmt.Printf("serviceId    : %s\n", serviceId)
	fmt.Printf("serviceName  : %s\n", serviceName)
	fmt.Printf("httpRoute    : /ws\n")
	fmt.Printf("listen       : %s\n", listenAddr)
}

// SendChatMessage 发送聊天消息
func SendChatMessage(connId string, chat *ChatMessage) error {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, chat); err != nil {
		return err
	}
	conn, ok := GlobalConnManager.GetConn(connId)
	if !ok {
		return errors.New("SendChatMessage get conn error")
	}
	return conn.Push(&WSMessage{MsgType: websocket.TextMessage, MsgData: buf.Bytes()})
}
