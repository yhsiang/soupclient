package packet

import (
	"encoding/binary"
)

type Packet struct {
	Type    byte
	Payload string
}

var mapping map[byte]string

func init() {
	mapping = make(map[byte]string)

	mapping['H'] = "server heartbeat"
	mapping['R'] = "client heartbeat"
	mapping['+'] = "debug"
	mapping['U'] = "unsequenced data"
}

func (p Packet) Bytes() []byte {
	var data []byte

	data = append(data, p.Type)
	data = append(data, []byte(p.Payload)...)
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(len(data)))

	return append(bs, data...)
}

func (p Packet) TypeName() string {
	return mapping[p.Type]
}

func NewPacket(bs []byte) *Packet {
	p := &Packet{}

	packetLen := binary.BigEndian.Uint16(bs[:2])

	p.Type = bs[2]

	if packetLen > 1 {
		p.Payload = string(bs[3:])
	}

	return p
}
