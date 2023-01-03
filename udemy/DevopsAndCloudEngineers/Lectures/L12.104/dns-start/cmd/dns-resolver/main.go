package main

import (
	"fmt"
	"net"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/Lectures/L12.104/dns-start/pkg/dns"
)

func main() {
	fmt.Printf("Starting DNS Server...\n")

	packetConnection, err := net.ListenPacket("udp", ":53")
	if err != nil {
		panic(err)
	}
	defer packetConnection.Close()

	for {
		buf := make([]byte, 512)
		n, addr, err := packetConnection.ReadFrom(buf)
		if err != nil {
			fmt.Printf("Read error from [%s]: %s\n", addr.String(), err)
		}
		go dns.HandlePacket(packetConnection, addr, buf[:n])
	}
}
