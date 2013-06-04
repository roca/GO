// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv6

import (
	"net"
	"syscall"
)

// A Conn represents a network endpoint that uses IPv6 transport.
// It allows to set basic IP-level socket options such as traffic
// class and hop limit.
type Conn struct {
	genericOpt
}

type genericOpt struct {
	net.Conn
}

func (c *genericOpt) ok() bool { return c != nil && c.Conn != nil }

// PathMTU returns a path MTU value for the destination associated
// with the endpoint.
func (c *Conn) PathMTU() (int, error) {
	if !c.genericOpt.ok() {
		return 0, syscall.EINVAL
	}
	fd, err := c.genericOpt.sysfd()
	if err != nil {
		return 0, err
	}
	return ipv6PathMTU(fd)
}

// NewConn returns a new Conn.
func NewConn(c net.Conn) *Conn {
	return &Conn{
		genericOpt: genericOpt{Conn: c},
	}
}

// A PacketConn represents a packet network endpoint that uses IPv6
// transport.  It is used to control several IP-level socket options
// including IPv6 header manipulation.  It also provides datagram
// based network I/O methods specific to the IPv6 and higher layer
// protocols such as OSPF, GRE, and UDP.
type PacketConn struct {
	genericOpt
	dgramOpt
	payloadHandler
}

type dgramOpt struct {
	net.PacketConn
}

func (c *dgramOpt) ok() bool { return c != nil && c.PacketConn != nil }

// SetControlMessage allows to receive the per packet basis IP-level
// socket options.
func (c *PacketConn) SetControlMessage(cf ControlFlags, on bool) error {
	if !c.payloadHandler.ok() {
		return syscall.EINVAL
	}
	fd, err := c.payloadHandler.sysfd()
	if err != nil {
		return err
	}
	return setControlMessage(fd, &c.payloadHandler.rawOpt, cf, on)
}

// NewPacketConn returns a new PacketConn using c as its underlying
// transport.
func NewPacketConn(c net.PacketConn) *PacketConn {
	return &PacketConn{
		genericOpt:     genericOpt{Conn: c.(net.Conn)},
		dgramOpt:       dgramOpt{PacketConn: c},
		payloadHandler: payloadHandler{PacketConn: c},
	}
}
