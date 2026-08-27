package base

type Packet struct {
	data     []byte
	priority int
}

func NewPacket(data []byte, priority int) *Packet {
	return &Packet{data: data, priority: priority}
}

func (p *Packet) GetPriority() int {
	return p.priority
}

func (p *Packet) GetData() []byte {
	return p.data
}
