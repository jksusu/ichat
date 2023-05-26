package main

import (
	"github.com/golang/glog"
	"ichat"
	"ichat/internal/area/gate"
	_ "net/http/pprof"
)

func main() {
	/*go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()*/
	if err := ichat.Init(); err != nil {
		glog.Fatalf("init fail %s", err)
	}
	if err := gate.NewWsServer().Start(); err != nil {
		glog.Fatalf("ws start fail:%v", err)
	}
}
