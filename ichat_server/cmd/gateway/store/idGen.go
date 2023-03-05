package store

import "github.com/bwmarrin/snowflake"

type IDGenerator struct {
	Node *snowflake.Node
}

var Gen *IDGenerator

func NewIDGenerator(nodeID int64) error {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return err
	}
	Gen = &IDGenerator{Node: node}

	return nil
}
