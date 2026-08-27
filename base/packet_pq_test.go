package base_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/Vardhanb07/CNP/base"
)

func TestPacketPriorityQueue(t *testing.T) {
	ppq := base.NewPacketPriorityQueue()
	var lp []*base.Packet
	for i := 10; i >= 0; i-- {
		lp = append(lp, base.NewPacket(fmt.Appendf([]byte(""), "p%v", i), i))
	}
	for _, p := range lp {
		ppq.Push(p)
	}
	var nlp []int
	for {
		p, ok := ppq.Pop()
		if !ok {
			break
		}
		nlp = append(nlp, p.GetPriority())
	}
	if !slices.IsSorted(nlp) {
		t.Errorf("priority queue fails in sorting packets by priority")
	}
}
