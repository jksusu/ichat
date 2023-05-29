# websocket即时通讯服务器

## 支持的消息类型

```
消息通过json格式进行交互
如果消息格式不符合规范会自动丢弃

预先设置两个用户
739613
ws://127.0.0.1:9501/ws?token=c3293857bc70d6ac198086386f33f78c
469264
ws://127.0.0.1:9501/ws?token=537b469a2b3811b87d64420e0b003bb5

客户端定时发送心跳
{"type":4}  

私聊 
{"reqId":1,"from":"739613","to":"469264","type":1,"route":"ic.talk.to.user","Data":{"content":"我是a","type":1,"extra":""}}
{"reqId":1,"from":"469264","to":"739613","type":1,"route":"ic.talk.to.user","Data":{"content":"我是b","type":1,"extra":""}}

```