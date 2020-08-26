package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"syscall"
)

// ArpHeader specifies the header for an ARP message.
type ArpHeader struct {
	HardwareType          uint16
	ProtocolType          uint16
	HardwareAddressLength uint8
	ProtocolAddressLength uint8
	Opcode                uint16
}

// ArpMessage represents an ARP message.
type ArpMessage struct {
	ArpHeader
	senderHardwareAddress []byte
	senderProtocolAddress []byte
	targetHardwareAddress []byte
	targetProtocolAddress []byte
}

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

const (
	opARPRequest = 1
	opARPReply   = 2
)

var (
	ethernetBroadcast = net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	ethernetZero      = net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

// bytes returns the wire representation of the ARP message.
func (m *ArpMessage) bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, m.ArpHeader); err != nil {
		return nil, fmt.Errorf("binary write failed: %v", err)
	}
	buf.Write(m.senderHardwareAddress)
	buf.Write(m.senderProtocolAddress)
	buf.Write(m.targetHardwareAddress)
	buf.Write(m.targetProtocolAddress)
	//pad := make([]byte,18)
	//buf.Write(pad)

	return buf.Bytes(), nil
}

// bytes returns the wire representation of the ARP message.
func (m *ArpMessage) Read(b []byte) {

	m.HardwareType = binary.BigEndian.Uint16(b[:2])
	m.ProtocolType = binary.BigEndian.Uint16(b[2:4])
	m.HardwareAddressLength = b[4]
	m.ProtocolAddressLength = b[5]
	m.Opcode = binary.BigEndian.Uint16(b[6:8])

	fmt.Println("Hardware Type:", m.HardwareType)
	fmt.Printf("Protocol Type: 0x%x\n", m.ProtocolType)
	fmt.Println("Hardware Length:", m.HardwareAddressLength)
	fmt.Println("Protocol Address Length:", m.ProtocolAddressLength)
	fmt.Println("Opcode:", m.Opcode)

	fmt.Println("Sender Hardware Address", net.HardwareAddr(b[8:14]).String())
	fmt.Println("Sender IP Address", net.IP(b[14:18]))
	fmt.Println("Target Hardware Address", net.HardwareAddr(b[18:24]).String())
	fmt.Println("Target IP Address", net.IP(b[24:28]))

}

func main() {

	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_DGRAM, int((syscall.ETH_P_ARP)))
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	iface := getINterface(os.Args[2])

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

	ipd := net.IP{}
	ipd.UnmarshalText([]byte(os.Args[1]))

	m := &ArpMessage{
		ArpHeader{
			1,            // IEEE 802
			0x0800,       // Ethernet
			6,            // 48-bit MAC Address
			net.IPv4len,  // 32-bit IPv4 Address
			opARPRequest, // ARP Reply
		},
		iface.HardwareAddr,
		ip.To4(),
		ethernetZero,
		ipd.To4(),
	}

	if err := syscall.BindToDevice(fd, iface.Name); err != nil {
		panic(err)
	}

	v := syscall.ETH_P_ARP

	pv := uint16(v>>8) | uint16(v<<8)

	ll := syscall.SockaddrLinklayer{
		Protocol: uint16(pv),
		Ifindex:  iface.Index,
		Pkttype:  0, // syscall.PACKET_HOST
		Hatype:   m.HardwareType,
		Halen:    m.HardwareAddressLength,
	}

	target := ethernetBroadcast

	for i := 0; i < len(target); i++ {
		ll.Addr[i] = target[i]
	}

	b, err := m.bytes()
	if err != nil {
		panic(err)
	}

	if err := syscall.Bind(fd, &ll); err != nil {
		panic(err)
	}
	if err := syscall.Sendto(fd, b, 0, &ll); err != nil {
		panic(err)
	}

	b = make([]byte, 1024)
	n, _, err := syscall.Recvfrom(fd, b, 0)
	if err != nil {
		panic(err)
	}

	in := bytes.Buffer{}
	in.Write(b[:n])

	incomm := ArpMessage{}

	incomm.Read(b[:n])

}
