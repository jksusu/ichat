package idgen

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

type IDGenerator struct {
	Node *snowflake.Node
}

var Gen *IDGenerator

func NewIDGenerator(nodeId int64) error {
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		return err
	}
	Gen = &IDGenerator{Node: node}
	return nil
}

// 生成消息id
func GenChatMsgId() int64 {
	return Gen.Node.Generate().Int64()
}

// 业务 + 节点id + 数字
func GenConnectId() string {
	return fmt.Sprintf("%d%s", Gen.Node.Generate(), "b")
}
