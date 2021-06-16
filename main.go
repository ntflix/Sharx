package main

import (
	"bufio"
	"os"

	"github.com/ntflix/sharx/serviceAdvertiser"
)

func main() {
	sa := new(serviceAdvertiser.ServiceAdvertiser)
	sa.Init()
	sa.StartService()

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	sa.StopService()
}
