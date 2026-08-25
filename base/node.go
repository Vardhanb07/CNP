package base

type Node interface {
	Name() string
	SetNext(nodes ...Node)
	GetNext() []Node
	Transfer(packet *Packet)
}
