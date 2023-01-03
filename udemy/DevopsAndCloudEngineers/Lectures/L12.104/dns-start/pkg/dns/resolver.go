package dns

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"

	"golang.org/x/net/dns/dnsmessage"
)

const ROOT_SERVERS = "198.41.0.4,199.9.14.201,192.33.4.12,199.7.91.13,192.203.230.10,192.5.5.241,192.112.36.4,198.97.190.53"

func HandlePacket(pc net.PacketConn, addr net.Addr, buf []byte) {
	if err := handlePacket(pc, addr, buf); err != nil {
		fmt.Printf("Error handling packet: [%s]: %s\n", addr.String(), err)
	}
}

func handlePacket(pc net.PacketConn, addr net.Addr, buf []byte) error {
	p := dnsmessage.Parser{}
	header, err := p.Start(buf)
	if err != nil {
		return err
	}
	question, err := p.Question()
	if err != nil {
		return err
	}

	return fmt.Errorf("not implemented yet")
}

func outgoingDnsQuery(servers []net.IP, question dnsmessage.Question) (*dnsmessage.Parser, *dnsmessage.Header, error) {
	fmt.Printf("New outgoing dns query for %s, servers: %+v\n", question.Name.String(), servers)
	randomID, err := RandUint16()
	if err != nil {
		return nil, nil, err
	}

	message := dnsmessage.Message{
		Header: dnsmessage.Header{
			ID:       randomID,
			Response: false,
			OpCode:   dnsmessage.OpCode(0),
		},
		Questions: []dnsmessage.Question{question},
	}

	buf, err := message.Pack()
	if err != nil {
		return nil, nil, err
	}

	var conn net.Conn
	for _, server := range servers {
		conn, err = net.Dial("udp", server.String()+":53")
		if err == nil {
			break
		}
	}
	if conn == nil {
		return nil, nil, fmt.Errorf("Failed to make connection to servers: %s", err)
	}

	_, err = conn.Write(buf)
	if err != nil {
		return nil, nil, err
	}

	answer := make([]byte, 512)
	n, err := bufio.NewReader(conn).Read(answer)
	if err != nil {
		return nil, nil, err
	}

	conn.Close()

	var parser dnsmessage.Parser

	header, err := parser.Start(answer[:n])
	if err != nil {
		return nil, nil, fmt.Errorf("parser start error: %s", err)
	}

	if header.ID != randomID {
		return nil, nil, fmt.Errorf("ID in response does not match")
	}

	questions, err := parser.AllQuestions()
	if err != nil {
		return nil, nil, fmt.Errorf("parser questions error: %s", err)
	}

	if len(questions) != len(message.Questions) {
		return nil, nil, fmt.Errorf("number of questions in response does not match")
	}

	err = parser.SkipAllQuestions()
	if err != nil {
		return nil, nil, fmt.Errorf("parser skip questions error: %s", err)
	}

	return &parser, &header, nil
}

func RandUint16() (uint16, error) {
	max := ^uint16(0)
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}

	return uint16(randomNumber.Int64()), nil
}
