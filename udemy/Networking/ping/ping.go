package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func getINterface(name string) net.Interface {
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

func listen2Client(conn *net.IPConn) {
	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	fmt.Println("++++++++++++++++++++++++++++++++ IP Packet ++++++++++++++++++++++++++++++++")
	fmt.Println("Protocol Version:", b[0]>>4)
	fmt.Println("Header Length:", b[0]&0x0f*4)
	fmt.Println("Total Length:", binary.BigEndian.Uint16(b[2:4]))
	fmt.Printf("Source Address: %d.%d.%d.%d\n", b[12], b[13], b[14], b[15])
	fmt.Printf("Destination Address: %d.%d.%d.%d\n", b[16], b[17], b[18], b[19])
	fmt.Println("++++++++++++++++++++++++++++++++ ICMP Packet ++++++++++++++++++++++++++++++++")
	hlen := b[0] & 0x0f * 4
	b = b[hlen:n]
	fmt.Println("Type: ", b[0])
	fmt.Println("Code: ", b[1])
	fmt.Println("Message: ", string(b[4:]))
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

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
	fmt.Print("Enter the address to ping:")
	fmt.Scanf("%s\n", &add)
	for {
		msg := icmp.Message{}
		msg.Type = ipv4.ICMPTypeEcho
		msg.Code = 0
		msg.Body = &icmp.Echo{
			ID:   1,
			Seq:  i,
			Data: []byte("abcdefghijklmnopqrstuvwxyz1234567890"),
		}
		b, _ := msg.Marshal(nil)

		dst, err := net.ResolveIPAddr("ip4", add)
		if err != nil {
			panic(err)
		}
		IPconn.WriteTo(b, dst)
		start := time.Now()
		listen2Client(IPconn)
		end := time.Since(start).Milliseconds()
		fmt.Println("Time taken to the destination is:", end, "milliseconds")

		time.Sleep(time.Second * 2)
		i++
	}
}
