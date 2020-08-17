package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/rajnikant12345/gotutorials/datalinkclientserver"
	"net"
	"os"
	"syscall"
)


func send( fd int, ll syscall.Sockaddr, iface net.Interface  ) {

	for {
		fmt.Print("Please enter message:")
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		b := line + " From : " + iface.HardwareAddr.String()
		if err := syscall.Sendto(fd, []byte(b), 0, ll); err != nil {
			panic(err)
		}
	}
}


func main() {
	ifaceName := "eth0"
	protocol := "0003"
	destMac := "ff:ff:ff:ff:ff:ff"

	if len(os.Args) != 4 {
		fmt.Println("Usage ./rec <interfacename> <protocol> <destmac>")
		fmt.Println("Values defaults to eth0 , 0003 and ff:ff:ff:ff:ff:ff")
	}else {
		ifaceName = os.Args[1]
		protocol = os.Args[2]
		destMac  = os.Args[3]
	}

	iface, err := datalinkclientserver.GetInterface(ifaceName)

	if err != nil {
		panic(err)
	}

	pp,_ := hex.DecodeString(protocol)


	fd,err := datalinkclientserver.CreateDataLinkSocket(iface, false)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	ll := datalinkclientserver.BindSocketToAddress( fd, binary.BigEndian.Uint16(pp) ,iface )

	ll1 := datalinkclientserver.CreateDestination( destMac , ll )

	send(fd,ll1, iface  )

	//time.Sleep(time.Second*100000)

	/*rb := make([]byte, 8192)

	for {
		n,_,err := syscall.Recvfrom(fd, rb, 0 )

		if err != nil {
			panic(err)
		}

		fmt.Println( rb[:n]  )

		destmac := net.HardwareAddr(rb[:6])
		srcmac := net.HardwareAddr(rb[6:12])

		fmt.Println("Destination Mac => ",destmac.String(), "|| Source Mac => ",srcmac.String())
		typ := binary.BigEndian.Uint16( rb[12:14] )
		fmt.Printf("Type: %X\n",typ)
		fmt.Println("Data: ", (rb[14:n]))
		fmt.Println("==================================================================================")

	}*/
}
