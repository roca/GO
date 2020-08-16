package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"syscall"
)

func main() {
	ifaceName := ""
	protocol := ""
	fmt.Print("Enter Interface name: ")
	fmt.Scanf("%s\n", &ifaceName)

	if ifaceName == "" {
		ifaceName = "enp0s8"
	}

	iface, err := GetInterface(ifaceName)

	if err != nil {
		panic(err)
	}

	fmt.Print("Enter protocol: ")
	fmt.Scanf("%s\n", &protocol)

	pp, _ := hex.DecodeString(protocol)

	fd, err := CreateDataLinkSocket(iface)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	BindSocketToAddress(fd, binary.BigEndian.Uint16(pp), iface)

	rb := make([]byte, 8192)

	for {
		n, _, err := syscall.Recvfrom(fd, rb, 0)

		if err != nil {
			panic(err)
		}

		destmac := net.HardwareAddr(rb[:6])
		srcmac := net.HardwareAddr(rb[6:12])

		fmt.Println("Destination Mac => ", destmac.String(), "|| Source Mac => ", srcmac.String())
		typ := binary.BigEndian.Uint16(rb[12:14])
		fmt.Printf("Type: %X\n", typ)
		fmt.Println("Data: ", string(rb[14:n]))
		fmt.Println("==================================================================================")

	}
}

func GetInterface(name string) (iface net.Interface, err error) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				if i.Name == name {
					iface = i
					return
				}
			}
		}
	}
	return net.Interface{}, errors.New("No interface found")
}

func CreateDataLinkSocket(iface net.Interface) (fd int, err error) {
	fd, err = syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 0)

	if err != nil {
		return -1, err
	}

	if err = syscall.BindToDevice(fd, iface.Name); err != nil {
		return
	}

	return
}

/*
Bind socket to send receive over interface and assign a specific protocol.
*/
func BindSocketToAddress(fd int, protocol uint16, iface net.Interface) *syscall.SockaddrLinklayer {

	pv := uint16(protocol>>8) | uint16(protocol<<8)

	ll := &syscall.SockaddrLinklayer{
		Protocol: uint16(pv),
		Ifindex:  iface.Index,
		Pkttype:  0, // syscall.PACKET_HOST
		Hatype:   1,
		Halen:    6,
	}

	if err := syscall.Bind(fd, ll); err != nil {
		panic(err)
	}

	return ll

}
