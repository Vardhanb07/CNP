package base

type Packet struct {
	Data     []byte
	Priority int
}

func NewPacket(data []byte) *Packet {
	return &Packet{Data: data}
}

func (p *Packet) GetPriority() int {
	return p.Priority
}
