package main

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/soupclient/packet"
)

func mockServer(t *testing.T) {
	serverHB := packet.Packet{
		Type: 'H',
	}
	hb := packet.Packet{
		Type: 'R',
	}
	debug := packet.Packet{
		Type:    '+',
		Payload: "debug test",
	}
	data1 := packet.Packet{
		Type:    'U',
		Payload: `{"message": "hello world"}`,
	}
	data2 := packet.Packet{
		Type:    'U',
		Payload: `{"message": "answer is 42"}`,
	}

	testcases := [][]byte{
		hb.Bytes(),
		hb.Bytes(),
		debug.Bytes(),
		hb.Bytes(),
		hb.Bytes(),
		hb.Bytes(),
		data1.Bytes(),
		hb.Bytes(),
		data2.Bytes(),
		hb.Bytes(),
		hb.Bytes(),
		hb.Bytes(),
		hb.Bytes(),
	}

	l, err := net.Listen("tcp", ":3000")
	assert.NoError(t, err)
	defer l.Close()

	buf := make([]byte, 1024)
	i := 0
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			reqLen, err := conn.Read(buf)
			assert.NoError(t, err)
			_, err = conn.Write(serverHB.Bytes())
			assert.NoError(t, err)
			assert.Equal(t, testcases[i], buf[:reqLen])
			i++
			if i == len(testcases) {
				return
			}
		}
	}
}

func TestMain(t *testing.T) {
	go mockServer(t)
	runClient(":3000")
}
