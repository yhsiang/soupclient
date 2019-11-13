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

func send(conn net.Conn, p *packet.Packet) error {
	name := p.TypeName()
	_, err := conn.Write(p.Bytes())
	if err != nil {
		fmt.Printf("Failed to send %s packet.\n", name)
		return err
	}
	fmt.Printf("Send %s to server.\n", name)
	return nil
}

func runClient(server string) {
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

	ticker := time.NewTicker(1 * time.Second)
	debugTimer := time.NewTimer(1300 * time.Millisecond)
	data46Timer := time.NewTimer(4600 * time.Millisecond)
	data54Timer := time.NewTimer(5400 * time.Millisecond)
	closeTimer := time.NewTimer(10000 * time.Millisecond)
	defer ticker.Stop()

	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Printf("Failed to connect %s.\n", server)
		return
	}
	fmt.Printf("Connect to server %s.\n", server)
	// send heartbeat at 0ms
	err = send(conn, &hb)
	if err != nil {
		return
	}

	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed receive packet from server.")
			return
		}
		serverPacket := packet.NewPacket(buf[:reqLen])
		fmt.Printf("Receive %s.\n", serverPacket.TypeName())
		if serverPacket.Type != 'H' {
			// preserve from none heart beat packet from server
			fmt.Printf("Payload: %s.\n", serverPacket.Payload)
		}

		select {
		case <-debugTimer.C:
			err = send(conn, &debug)
			if err != nil {
				return
			}
		case <-data46Timer.C:
			err = send(conn, &data1)
			if err != nil {
				return
			}
		case <-data54Timer.C:
			err = send(conn, &data2)
			if err != nil {
				return
			}
		case <-closeTimer.C:
			conn.Close()
			fmt.Printf("Close connection.\n")
			return
		case <-ticker.C:
			err = send(conn, &hb)
			if err != nil {
				return
			}
		}
	}

}

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

	runClient(server)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
