# websocket即时通讯服务器

## 支持的消息类型

```
消息通过json格式进行交互
如果消息格式不符合规范会自动丢弃

预先设置两个用户
ws://127.0.0.1:9501/ws?token=c3293857bc70d6ac198086386f33f78c 739613
ws://127.0.0.1:9501/ws?token=537b469a2b3811b87d64420e0b003bb5 469264



1.心跳
{"type":4}  客户端定时发送心跳
{"type":4}  服务端收到心跳后返回信息


2.CHAT类型表示上层业务需要的消息
{
	"reqId": 1,
	"type": 1,
	"from": "739613",
	"route": "chat.user.message",
	"data": {
		"msgId": 1234,
		"to": "469264",
		"content": "dd",
		"type": 1,
		"send_time": 1678721586523,
		"extra": ""
	}
}

{"reqId":4,"type":1,"route":"chat.session.lists","from":"666666"}
3.消息ack
```

## 待优化点

> 连接map打散到更小的bucket中，减少推送遍历

## 消息ack

> 根据 route 中的内容来确定 状态
> * chat.message.server.ack 服务端已收到并已经缓存
> * chat.message.arrive.ack 已经送达，但还未读
> * chat.message.read.ack 已读回执

## 唯一数生成

> 1. 一个字节8个bit