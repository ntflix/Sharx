package socketHandler

import (
	"bufio"
	"fmt"
	"net"
)

type SocketHandler struct {
	Address  string
	Protocol string
	IP       *net.IP
}

func (sh *SocketHandler) Init(address string, protocol string) {
	sh.Address = address
	sh.Protocol = protocol

	// Listen on protocol and port
	link, _ := net.Listen(sh.Protocol, sh.Address)

	// Accept connections
	connection, _ := link.Accept()

	// Run loop forever or until interrupt
	for {
		// process connection data
		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print("Message Received:", string(message))
	}
}
