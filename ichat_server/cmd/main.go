package main

import (
	"fmt"
	"github.com/golang/glog"
	"ichat"
	"ichat/internal/gate"
)

func main() {
	if err := ichat.Init(); err != nil {
		glog.Fatalf("init fail %s", err)
	}
	if err := gate.NewWsServer().Start(); err != nil {
		glog.Fatalf("ws start fail:%v", err)
	}
	fmt.Println("over")
}
