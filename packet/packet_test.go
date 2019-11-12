package packet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerHeartBeatPacket(t *testing.T) {
	assert := assert.New(t)
	packet := &Packet{
		Type: 'H',
	}
	bs := packet.Bytes()
	str := fmt.Sprintf("%x", bs)
	assert.Equal("000148", str)
}

func TestClientHeartBeatPacket(t *testing.T) {
	assert := assert.New(t)
	packet := &Packet{
		Type: 'R',
	}
	bs := packet.Bytes()
	str := fmt.Sprintf("%x", bs)
	assert.Equal("000152", str)
}

func TestDebugPacketWithoutText(t *testing.T) {
	assert := assert.New(t)
	packet := &Packet{
		Type: '+',
	}
	bs := packet.Bytes()
	str := fmt.Sprintf("%x", bs)
	assert.Equal("00012b", str)
}

func TestDebugPacketWithText(t *testing.T) {
	assert := assert.New(t)
	packet := &Packet{
		Type:    '+',
		Payload: "test",
	}
	bs := packet.Bytes()
	str := fmt.Sprintf("%x", bs)
	assert.Equal("00052b74657374", str)
}

func TestUnsequencedDataPacketWithoutText(t *testing.T) {
	assert := assert.New(t)
	packet := &Packet{
		Type: 'U',
	}
	bs := packet.Bytes()
	str := fmt.Sprintf("%x", bs)
	assert.Equal("000155", str)
}

func TestUnsequencedDataPacketWithText(t *testing.T) {
	assert := assert.New(t)
	packet := &Packet{
		Type:    'U',
		Payload: `{"status":"success"}`,
	}
	bs := packet.Bytes()
	str := fmt.Sprintf("%x", bs)
	assert.Equal("0015557b22737461747573223a2273756363657373227d", str)
}
