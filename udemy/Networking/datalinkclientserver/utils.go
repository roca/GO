package datalinkclientserver

import (
	"bytes"
	"net"
	"syscall"
	"errors"
)


func GetInterface( name string ) ( iface net.Interface, err error ) {
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


func CreateDataLinkSocket(  iface net.Interface, raw bool ) ( fd int,  err error ) {

	s := syscall.SOCK_DGRAM

	if raw {
		s = syscall.SOCK_RAW
	}

	fd, err = syscall.Socket(syscall.AF_PACKET, s, 0 )
	if err != nil {
		return
	}

	if err != nil {
		return -1,err
	}

	if err = syscall.BindToDevice(fd, iface.Name); err != nil {
		return
	}

	return
}


/*
Bind socket to send receive over interface and assign a specific protocol.
*/
func BindSocketToAddress( fd int, protocol uint16, iface net.Interface  ) *syscall.SockaddrLinklayer {

	pv := uint16(protocol >> 8) | uint16( protocol << 8 )

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


// choose destination mac for sensing packet to. we can use the link layer structure bind to socket.
func CreateDestination( target string ,ll *syscall.SockaddrLinklayer ) *syscall.SockaddrLinklayer {


	hwaddress,err  := net.ParseMAC(target)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(hwaddress); i++ {
		ll.Addr[i] = hwaddress[i]
	}

	return ll

}




