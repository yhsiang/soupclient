package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yhsiang/soupclient/packet"
)

func main() {
	flag.Parse()
	var server string

	if flag.NArg() > 0 {
		server = flag.Arg(0)
	} else {
		fmt.Printf("Please provide server host and port to connect.\n")
		fmt.Printf("Example: localhost:30010\n")
		return
	}
	hb := &packet.Packet{
		Type: 'R',
	}

	debug := &packet.Packet{
		Type:    '+',
		Payload: "debug test",
	}

	data1 := &packet.Packet{
		Type:    'U',
		Payload: `{"message": "hello world"}`,
	}

	data2 := &packet.Packet{
		Type:    'U',
		Payload: `{"message": "answer is 42"}`,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(1 * time.Second)
	debugTimer := time.NewTimer(1300 * time.Millisecond)
	data46Timer := time.NewTimer(4600 * time.Millisecond)
	data54Timer := time.NewTimer(5400 * time.Millisecond)
	closeTimer := time.NewTimer(10000 * time.Millisecond)
	defer ticker.Stop()

	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Printf("Failed to connect "+server+"\n", err)
		return
	}

	// send heartbeat at 0ms
	conn.Write(hb.Bytes())

	for {
		select {
		case <-debugTimer.C:
			conn.Write(debug.Bytes())
		case <-data46Timer.C:
			conn.Write(data1.Bytes())
		case <-data54Timer.C:
			conn.Write(data2.Bytes())
		case <-closeTimer.C:
			conn.Close()
			return
		case <-ticker.C:
			conn.Write(hb.Bytes())
		case <-sigs:
			return
		default:
		}
	}

}
