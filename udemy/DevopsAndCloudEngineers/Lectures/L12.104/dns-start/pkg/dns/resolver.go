package dns

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"strings"

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

	response, err := dnsQuery(getRootServers(), question)
	if err != nil {
		return err
	}
	response.Header.ID = header.ID
	responseBuf, err := response.Pack()
	if err != nil {
		return err
	}

	_, err = pc.WriteTo(responseBuf, addr)
	if err != nil {
		return err
	}

	return nil
}

func dnsQuery(servers []net.IP, question dnsmessage.Question) (*dnsmessage.Message, error) {
	fmt.Printf("Question %+v\n", question)

	message := dnsmessage.Message{}
	for i := 0; i < 3; i++ {
		dnsAnswer, header, err := outgoingDnsQuery(servers, question)
		if err != nil {
			return nil, err
		}

		parsedAnswers, err := dnsAnswer.AllAnswers()
		if err != nil {
			return nil, err
		}
		if header.Authoritative {
			message.Header.Response = true
			message.Answers = parsedAnswers
			return &message, nil
		}

		authorities, err := dnsAnswer.AllAuthorities()
		if err != nil {
			return nil, err
		}

		if len(authorities) == 0 {
			message.Header.RCode = dnsmessage.RCodeNameError
			return &message, nil
		}

		nameservers := make([]string, len(authorities))
		for k, authority := range authorities {
			if authority.Header.Type == dnsmessage.TypeNS {
				nameservers[k] = authority.Body.(*dnsmessage.NSResource).NS.String()
			}
		}

		additionals, err := dnsAnswer.AllAdditionals()
		if err != nil {
			return nil, err
		}
		newResolverServersFound := false
		resolvedServers := []net.IP{}
		for _, additional := range additionals {
			if additional.Header.Type == dnsmessage.TypeA {
				for _, nameserver := range nameservers {
					if additional.Header.Name.String() == nameserver {
						newResolverServersFound = true
						resolvedServers = append(resolvedServers, additional.Body.(*dnsmessage.AResource).A[:])
					}
				}
			}
		}

		servers = resolvedServers

		if !newResolverServersFound {
			for _, nameserver := range nameservers {
				if !newResolverServersFound {
					response, err := dnsQuery(getRootServers(), dnsmessage.Question{Name: dnsmessage.MustNewName(nameserver), Type: dnsmessage.TypeA, Class: dnsmessage.ClassINET})
					if err != nil {
						fmt.Printf("Warning: lookup of nameserver %s failed: %s\n", nameserver, err)
					} else {
						newResolverServersFound = true
						for _, answer := range response.Answers {
							if answer.Header.Type == dnsmessage.TypeA {
								servers = append(servers, answer.Body.(*dnsmessage.AResource).A[:])
							}
						}
					}
				}

			}
			
		}

	}

	message.Header.RCode = dnsmessage.RCodeServerFailure
	return &message, nil
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

func getRootServers() []net.IP {
	rootServers := []net.IP{}
	for _, rootServer := range strings.Split(ROOT_SERVERS, ",") {
		rootServers = append(rootServers, net.ParseIP(rootServer))
	}
	return rootServers
}
