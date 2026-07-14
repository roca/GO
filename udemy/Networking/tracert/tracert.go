package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"time"
)

func getINterface( name string ) net.Interface {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				if i.Name == name {
					return i
				}
			}
		}
	}
	panic("interface not found ")
}

func listen2Client(conn *net.IPConn, hop int) {
	b := make([]byte,1024)
	n,_ := conn.Read(b)


	source := fmt.Sprintf("%d.%d.%d.%d",b[12],b[13],b[14],b[15])
	hlen := b[0]&0x0f*4
	b = b[hlen:n]

	//fmt.Println(b)

	if b[0] == 11 {
		fmt.Println("In Transit To: ", source, " In Hop:" , hop)
	}

	if b[0] == 0 {
		fmt.Println("Reached To: ", source," In Hop:" , hop)
		os.Exit(0)
	}



}

func main() {


	iface := getINterface(os.Args[1])

	addrs, err := iface.Addrs()

	var ip net.IP
	// handle err
	func() {
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				return
			case *net.IPAddr:
				ip = v.IP
				return
			}
			// process IP address
		}

	}()

	IPAddr, err := net.ResolveIPAddr("ip4", ip.String())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	IPconn, err := net.ListenIP("ip4:icmp", IPAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	i := 0
	add := ""
	fmt.Print("Enter the address to trace:")
	fmt.Scanf("%s\n",&add)

	opts := ipv4.NewConn(IPconn)

	ttl := 1

	for {

		if err := opts.SetTTL(ttl); err != nil {
			panic(err)
		}

		msg := icmp.Message{}
		msg.Type = ipv4.ICMPTypeEcho
		msg.Code = 0
		msg.Body =  &icmp.Echo{
			ID:   1,
			Seq:  i,
			Data: []byte("rajnikant12345"),
		}
		b,_ := msg.Marshal( nil )

		dst, err := net.ResolveIPAddr("ip4", add)
		if err != nil {
			panic(err)
		}

		IPconn.WriteTo( b , dst)
		listen2Client(IPconn, ttl )
		time.Sleep(time.Second*2)
		ttl++
	}
}