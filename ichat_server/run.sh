#!/bin/sh

#定义各个Go项目的路径
gateway=./cmd/gateway/
logic=./cmd/logic/

#启动各个项目
cd gateway && go run main.go
cd logic && go run main.go