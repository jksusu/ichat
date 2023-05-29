package main

import (
	"fmt"
	"github.com/golang/glog"
	"ichat"
	"ichat/internal/gate"
	"ichat/internal/http"
)

func main() {
	if err := ichat.Init(); err != nil {
		glog.Fatalf("init fail %v", err)
	}
	if err := http.RegisterApi(); err != nil {
		glog.Fatalf("http run fail:%v", err)
	}
	if err := gate.NewWsServer().Start(); err != nil {
		glog.Fatalf("ws start fail:%v", err)
	}
	fmt.Println("over")
}
