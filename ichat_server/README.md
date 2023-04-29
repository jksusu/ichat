# websocket即时通讯服务器

## 支持的消息类型

```
消息通过json格式进行交互
如果消息格式不符合规范会自动丢弃

预先设置两个用户
ws://127.0.0.1:9501/ws?token=c3293857bc70d6ac198086386f33f78c 739613
ws://127.0.0.1:9501/ws?token=537b469a2b3811b87d64420e0b003bb5 469264

客户端定时发送心跳
{"type":4}  

添加好友
{"reqId":1,"type":1,"from":"469264","route":"ichat.relation.friends.apply","data":{"to":"","remark":"dd"}}
{"reqId":1,"type":1,"from":"739613","route":"ichat.relation.friends.apply.reply","data":{"msgId":"1644963360732614657","status":2}}

```