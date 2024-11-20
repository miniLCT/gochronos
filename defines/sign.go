package defines

import (
	"encoding/json"
	"errors"
)

const (
	FieldUUID      = "uuid"
	FieldGraphName = "graph_name"
	FieldNodeName  = "node_name"
)

var (
	// ErrNodeStoped is one operator in node stoped, graph will be stoped, message will be discard
	ErrNodeStoped = errors.New("operator stoped")
)

type NodeSign struct {
	GraphName string `json:"graph_name"` // 图名称
	NodeName  string `json:"node_name"`  // 节点名称
	UUID      string `json:"uuid"`       // 唯一标识
}

// TODO 不应该存在这种check的逻辑，直接提供一个初始化函数，强制校验
func (ns *NodeSign) Check() error {
	if ns.GraphName == "" {
		return errors.New("empty graph name")
	}
	if ns.NodeName == "" {
		return errors.New("empty node name")
	}
	if ns.UUID == "" {
		return errors.New("empty uuid")
	}
	return nil
}

func (ns *NodeSign) String() string {
	bs, _ := json.Marshal(ns)
	return string(bs)
}
