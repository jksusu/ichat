# websocket即时通讯服务器


## 支持的消息类型
```
消息通过json格式进行交互
如果消息格式不符合规范会自动丢弃

1.心跳
{"type":"PING"}  客户端定时发送心跳
{"type":"PONG"}  服务端收到心跳后返回信息


2.CHAT类型表示上层业务需要的消息

{
	"type": "CHAT",
	"data": {
		"msgv": "1",
		"msgId": 0,
		"seq": 0,
		"route": "chat.user.message",
		"from": "1212121",
		"to": "1212121",
		"content": "hello",
		"type": 1,
		"Extra": ""
	}
}
3.消息ack
```

## 待优化点
> 连接map打散到更小的bucket中，减少推送遍历
## 消息ack
> 根据 route 中的内容来确定 状态
> * chat.message.server.ack 服务端已收到并已经缓存
> * chat.message.arrive.ack 已经送达，但还未读
> * chat.message.read.ack   已读回执

## 唯一数生成
>1. 一个字节8个bit