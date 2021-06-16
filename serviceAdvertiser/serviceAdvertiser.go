package serviceAdvertiser

import (
	"net"
	"os"

	"github.com/hashicorp/mdns"
)

type ServiceAdvertiser struct {
	Host string
	Info []string

	IPAddresses []net.IP
	service     *mdns.MDNSService
	server      *mdns.Server
}

func (sa *ServiceAdvertiser) Init(port int) {
	// Set up our service export
	sa.Host, _ = os.Hostname()
	sa.Info = []string{}
	sa.IPAddresses = []net.IP{} // Set our IP addresses to nil so it advertises on all IP interfaces

	sa.service, _ = mdns.NewMDNSService(
		sa.Host,
		"_sharx._tcp",
		"local.",
		"",
		port,
		sa.IPAddresses,
		sa.Info,
	)
}

func (sa *ServiceAdvertiser) StartService() {
	sa.server, _ = mdns.NewServer(&mdns.Config{Zone: sa.service})
}

func (sa *ServiceAdvertiser) StopService() {
	sa.server.Shutdown()
}
