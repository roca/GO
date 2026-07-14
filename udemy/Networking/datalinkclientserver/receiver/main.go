package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/rajnikant12345/gotutorials/datalinkclientserver"
	"net"
	"os"
	"syscall"
	"time"
)


func send(iface net.Interface , fd int, ll syscall.Sockaddr) {

	for {
		time.Sleep(time.Second*2)
		b := "Hi This Is A Message From : " + iface.HardwareAddr.String()
		if err := syscall.Sendto(fd, []byte(b), 0, ll); err != nil {
			panic(err)
		}
	}
}

func main() {
	ifaceName := "eth0"
	protocol := "0003"

	if len(os.Args) != 3 {
		fmt.Println("Usage ./rec <interfacename> <protocol>")
		fmt.Println("Values defaults to eth0 and 0003")
	} else {
		ifaceName = os.Args[1]
		protocol = os.Args[2]
	}

	iface, err := datalinkclientserver.GetInterface(ifaceName)

	if err != nil {
		panic(err)
	}

	pp,_ := hex.DecodeString(protocol)


	fd,err := datalinkclientserver.CreateDataLinkSocket(iface, true)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	datalinkclientserver.BindSocketToAddress( fd, binary.BigEndian.Uint16(pp) ,iface )


	rb := make([]byte, 8192)

	for {
		n,_,err := syscall.Recvfrom(fd, rb, 0 )

		if err != nil {
			panic(err)
		}

		destmac := net.HardwareAddr(rb[:6])
		srcmac := net.HardwareAddr(rb[6:12])

		fmt.Println("Destination Mac => ",destmac.String(), "|| Source Mac => ",srcmac.String())
		typ := binary.BigEndian.Uint16( rb[12:14] )
		fmt.Printf("Type: %X\n",typ)
		fmt.Println("Data: ", string(rb[14:n]))
		fmt.Println("==================================================================================")

	}
}
