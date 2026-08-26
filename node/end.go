package node

import (
	"fmt"

	"github.com/Vardhanb07/CNP/base"
	"go.uber.org/zap"
)

type EndNode struct {
	priority    uint8
	name        string
	nextNodes   *base.LStore[base.Node]
	packetStore *base.LStore[*base.Packet]
}

func NewEndNode(priority uint8, name string) *EndNode {
	nn := base.NewLStore[base.Node]()
	ps := base.NewLStore[*base.Packet]()
	return &EndNode{priority: priority, name: name, nextNodes: nn, packetStore: ps}
}

func (en *EndNode) Name() string {
	return en.name
}

func (en *EndNode) SetNext(nodes ...base.Node) {
	if len(nodes) != 1 {
		panic("end node can only have one next node")
	}
	en.nextNodes.Set(nodes[0])
}

func (en *EndNode) GetNext() []base.Node {
	return en.nextNodes.Get()
}

func (en *EndNode) Transfer(packet *base.Packet) {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	nodes := en.nextNodes.Get()
	en.packetStore.Set(packet)
	if len(nodes) == 0 {
		for _, p := range en.packetStore.Get() {
			logger.Info(
				"packetStore is empty, no transfer",
				zap.String("name", en.Name()),
				zap.String("type", "end"),
				zap.String("packet::data", string(p.Data)),
				zap.Int("packet::priority", p.GetPriority()),
			)
		}
		return
	}
	for _, p := range en.packetStore.Get() {
		for _, node := range nodes {
			logger.Info(
				fmt.Sprintf("transfer from %v to %v", en.Name(), node.Name()),
				zap.String("name", en.Name()),
				zap.String("type", "end"),
				zap.String("packet::data", string(p.Data)),
				zap.Int("packet::priority", p.GetPriority()),
			)
			node.Transfer(p)
		}
	}
	en.packetStore.Discard()
}
