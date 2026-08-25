package node

import (
	"github.com/Vardhanb07/CNP/base"
	"go.uber.org/zap"
)

type EndNode struct {
	priority  uint8
	name      string
	nextNodes base.Store[base.Node]
}

func NewEndNode(priority uint8, name string) *EndNode {
	return &EndNode{priority: priority, name: name}
}

func (en *EndNode) Name() string {
	return en.name
}

func (en *EndNode) SetNext(nodes ...base.Node) {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	if len(nodes) != 1 {
		panic("end node can only have one next node")
	}
	ok := en.nextNodes.Set(nodes[0])
	if !ok {
		logger.Info(
			"nextNodes.Set failed",
			zap.String("name", en.Name()),
			zap.String("type", "end"),
		)
	}
}

func (en *EndNode) GetNext() []base.Node {
	return en.nextNodes.Get()
}

func (en *EndNode) Transfer(packet *base.Packet) {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	nodes := en.nextNodes.Get()
	if len(nodes) == 0 {
		logger.Info(
			"nodeStore is empty, no transfer",
			zap.String("name", en.Name()),
			zap.String("type", "end"),
			zap.Any("packet", packet),
		)
		return
	}
	for _, node := range nodes {
		node.Transfer(packet)
	}
}
