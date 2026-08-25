package node

import (
	"github.com/Vardhanb07/CNP/base"
	"go.uber.org/zap"
)

type RouterNode struct {
	name        string
	packetStore *base.Store[base.Packet]
	nodeStore   *base.Store[base.Node]
}

func NewRouterNode(name string, limit int) *RouterNode {
	ps := base.NewStore[base.Packet](limit)
	ns := base.NewStore[base.Node](limit)
	return &RouterNode{name: name, packetStore: ps, nodeStore: ns}
}

func (rn *RouterNode) Name() string {
	return rn.name
}

func (rn *RouterNode) SetNext(nodes ...base.Node) {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	for _, n := range nodes {
		ok := rn.nodeStore.Set(n)
		if !ok {
			logger.Info(
				"nodeStore.Set failed",
				zap.String("name", rn.Name()),
				zap.String("type", "router"),
			)
		}
	}
}

func (rn *RouterNode) GetNext() []base.Node {
	return rn.nodeStore.Get()
}

func (rn *RouterNode) Transfer(packet *base.Packet) {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	nodes := rn.nodeStore.Get()
	if len(nodes) == 0 {
		logger.Info(
			"nodeStore is empty, no transfer",
			zap.String("name", rn.Name()),
			zap.String("type", "router"),
		)
		return
	}
	ok := rn.packetStore.Set(*packet)
	if !ok {
		logger.Info(
			"packetStore is full",
			zap.String("name", rn.Name()),
			zap.String("type", "router"),
		)
		for _, p := range rn.packetStore.Get() {
			for _, node := range nodes {
				node.Transfer(&p)
			}
		}
	}
}
