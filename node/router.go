package node

import (
	"fmt"

	"github.com/Vardhanb07/CNP/base"
	"go.uber.org/zap"
)

type RouterNode struct {
	name        string
	packetStore *base.Store[*base.Packet]
	packetQueue *base.PacketPriorityQueue
	nodeStore   *base.LStore[base.Node]
}

func NewRouterNode(name string, limit int) *RouterNode {
	ps := base.NewStore[*base.Packet](limit)
	ns := base.NewLStore[base.Node]()
	pq := base.NewPacketPriorityQueue()
	return &RouterNode{name: name, packetStore: ps, nodeStore: ns, packetQueue: pq}
}

func (rn *RouterNode) Name() string {
	return rn.name
}

func (rn *RouterNode) SetNext(nodes ...base.Node) {
	for _, n := range nodes {
		rn.nodeStore.Set(n)
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
	ok := rn.packetStore.Set(packet)
	if !ok {
		logger.Info(
			"packetStore is full",
			zap.String("name", rn.Name()),
			zap.String("type", "router"),
		)
		for _, p := range rn.packetStore.Get() {
			rn.packetQueue.Push(p)
		}
		for {
			p, ok := rn.packetQueue.Pop()
			if !ok {
				break
			}
			for _, node := range nodes {
				np := base.NewPacket(fmt.Appendf(p.GetData(), "from %v", rn.Name()), p.GetPriority())
				node.Transfer(np)
			}
		}
		rn.packetStore.Discard()
		rn.packetStore.Set(packet)
	}
	logger.Info(
		fmt.Sprintf("packet transfered to %v", rn.Name()),
		zap.String("name", rn.Name()),
		zap.String("type", "router"),
		zap.String("packet::data", string(packet.GetData())),
		zap.Int("packet::priority", packet.GetPriority()),
		zap.Int("limit::remaining", rn.packetStore.Limit()-rn.packetStore.Len()),
	)
}
