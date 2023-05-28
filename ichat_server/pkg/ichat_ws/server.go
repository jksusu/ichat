package ichat_ws

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"ichat/pkg/tools/idgen"
	"net"
	"net/http"
	"sync"
	"time"
)

type Config struct {
	ServiceId          string `json:"serviceid"`   //服务id
	ServiceName        string `json:"servicename"` //服务名
	WsPort             int    `json:"wsport"`      //端口
	WsRoute            string `json:"wsroute"`
	WsReadTimeout      int    `json:"wsreadtimeout"`      //毫秒读超时时间
	WsWriteTimeout     int    `json:"wswritetimeout"`     //毫秒写超时时间
	WsWriteChannelSize int    `json:"wswritechannelsize"` //写通道最大数量
	WsReadChannelSize  int    `json:"wsreadchannelsize"`  //读通道最大数量
	WsHeartbeatTimeout int    `json:"wsheartbeattimeout"` //秒心跳超时时间，超过这个时间触发重连
}

type Server struct {
	sync.Once
	server *http.Server
	connId string
	BeforeAcceptor
	Acceptor
	MessageListener
	StateListener
	*Config
	*Manager
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() (err error) {
	if s.Acceptor == nil || s.MessageListener == nil || s.StateListener == nil || s.BeforeAcceptor == nil {
		return errors.New("must be achieved Acceptor MessageListener StateListener BeforeAcceptor interface")
	}
	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("%s", s.WsRoute), s.handleConnect)
	s.server = &http.Server{
		ReadTimeout:  time.Duration(s.WsReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(s.WsWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}
	var listen net.Listener
	if listen, err = net.Listen("tcp", fmt.Sprintf(":%d", s.WsPort)); err != nil {
		return
	}

	s.Manager = &Manager{&sync.Map{}}
	GlobalConnManager = s.Manager
	glog.Info("\u001B[1;31;47mconn manager init success\u001B[0m")
	glog.Infof("\u001B[1;31;47mserverId:%s serverName:%s port:%d route:%s\u001B[0m", s.ServiceId, s.ServiceName, s.WsPort, s.WsRoute)

	go s.server.Serve(listen)
	for {
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) handleConnect(w http.ResponseWriter, r *http.Request) {
	if err := s.BeforeAccept(r); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	var (
		ws       *websocket.Conn
		err      error
		Upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
	)
	if ws, err = Upgrader.Upgrade(w, r, nil); err != nil {
		glog.Fatalf("ws conn upgrader fail %v", err)
	}
	id := idgen.GenConnectId()
	conn := HttpUpgraderToWebsocket(s, id, ws)
	glog.Infoln("new connect id:%v", id)

	defer conn.Close()

	s.Accept(conn, r)
	s.handler(conn)
}

func (s *Server) handler(c *Connect) {
	for {
		data, err := c.ReadMsg()
		if err != nil {
			c.Close()
			return
		}
		if data == nil || data.MsgType != 1 {
			continue
		}
		msg, err := DecodeMessageBody(data.MsgData)
		if err != nil {
			c.Close()
			return
		}
		switch msg.Type {
		case REQUESTPING:
			glog.Info("ping")
			if err = c.Pong(msg); err != nil {
				glog.Error(err)
				c.Close()
				return
			}
		case REQUEST:
			if data.MsgData != nil {
				s.Receive(c, data)
			}
		}
	}
}

type BeforeAcceptor interface {
	BeforeAccept(*http.Request) error
}
type Acceptor interface {
	Accept(*Connect, *http.Request) (uint64, error)
}
type MessageListener interface {
	Receive(*Connect, *WSMessage)
}
type StateListener interface {
	Disconnect(id string) error
}
