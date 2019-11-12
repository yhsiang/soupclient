package packet

import (
	"encoding/binary"
)

type Packet struct {
	// Length  int
	Type    byte
	Payload string
}

func (p Packet) Bytes() []byte {
	var data []byte

	data = append(data, p.Type)
	data = append(data, []byte(p.Payload)...)
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(len(data)))

	return append(bs, data...)
}
