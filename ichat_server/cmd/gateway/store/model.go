package store

const (
	MESSAGE_CONTENT = "chat_message_content"
	MESSAGE_INDEX   = "chat_message_index"
)

type ChatMessageContent struct {
	ID       int64  `gorm:"primarykey"`
	Type     int    `gorm:"default:1"`
	Content  string `gorm:"size:3000;not null"`
	Extra    string `gorm:"size:500"`
	SendTime int64  `gorm:"index"`
}

type ChatMessageIndex struct {
	ID       int64  `gorm:"primarykey"`
	A        string `gorm:"index;size:60;not null;comment:账户a"`
	B        string `gorm:"size:60;not null;comment:账户b"`
	ISend    int    `gorm:"default:0;not null;comment:是否为账户a发送"`
	MsgId    int64  `gorm:"not null;comment:关联消息内容表中的ID"`
	GroupId  string `gorm:"size:30;comment:群ID，单聊情况为空"`
	SendTime int64  `gorm:"index;not null;comment:消息发送时间"`
}
