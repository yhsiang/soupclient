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

func TestPacketTypeName(t *testing.T) {
	assert := assert.New(t)
	serverHeartBeat := &Packet{
		Type: 'H',
	}
	clientHeartBeat := &Packet{
		Type: 'R',
	}
	debug := &Packet{
		Type: '+',
	}
	data := &Packet{
		Type: 'U',
	}
	assert.Equal("server heartbeat", serverHeartBeat.TypeName())
	assert.Equal("client heartbeat", clientHeartBeat.TypeName())
	assert.Equal("debug", debug.TypeName())
	assert.Equal("unsequenced data", data.TypeName())
}

func TestNew(t *testing.T) {
	assert := assert.New(t)
	serverHeartBeat := &Packet{
		Type: 'H',
	}
	p := NewPacket(serverHeartBeat.Bytes())

	assert.Equal(serverHeartBeat.TypeName(), p.TypeName())

	data := &Packet{
		Type:    'U',
		Payload: `{"status":"success"}`,
	}
	p2 := NewPacket(data.Bytes())

	assert.Equal(data.TypeName(), p2.TypeName())
	assert.Equal(data.Payload, p2.Payload)

}
