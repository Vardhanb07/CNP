package base

import (
	"container/heap"
)

type PacketPriorityQueue struct {
	pq *PriorityQueue
}

func NewPacketPriorityQueue() *PacketPriorityQueue {
	pq := NewPriorityQueue()
	heap.Init(pq)
	return &PacketPriorityQueue{pq: pq}
}

func (ppq *PacketPriorityQueue) Push(p *Packet) {
	item := &Item{
		p: p,
	}
	heap.Push(ppq.pq, item)
}

func (ppq *PacketPriorityQueue) Pop() (*Packet, bool) {
	if ppq.pq.Len() <= 0 {
		return nil, false
	}
	item := heap.Pop(ppq.pq).(*Item)
	return item.p, true
}
