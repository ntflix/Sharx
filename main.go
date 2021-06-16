package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ntflix/sharx/serviceAdvertiser"
	"github.com/ntflix/sharx/socketHandler"
)

func main() {
	const port = 62335

	sa := new(serviceAdvertiser.ServiceAdvertiser)
	sa.Init(port)
	sa.StartService()

	sh := new(socketHandler.SocketHandler)
	sh.Init(":"+fmt.Sprint(port), "tcp")

	fmt.Printf("Listening on port " + fmt.Sprint(port))
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	sa.StopService()
}
