package ssh

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func StartServer(privateKey, authorizedKeys []byte) error {
	authorizedKeysMap := map[string]bool{}
	for len(authorizedKeys) > 0 {
		pubKey, _, _, rest, err := ssh.ParseAuthorizedKey(authorizedKeys)
		if err != nil {
			return fmt.Errorf("Failed to parse authorized_keys, err: %s", err)
		}
		authorizedKeysMap[string(pubKey.Marshal())] = true
		authorizedKeys = rest
	}

	// An SSH server is represented by a ServerConfig, which holds
	// certificate details and handles authentication of ServerConns.
	config := &ssh.ServerConfig{
		PublicKeyCallback: func(c ssh.ConnMetadata, pubKey ssh.PublicKey) (*ssh.Permissions, error) {
			if authorizedKeysMap[string(pubKey.Marshal())] {
				return &ssh.Permissions{
					// Record the public key used for authentication.
					Extensions: map[string]string{
						"pubkey-fp": ssh.FingerprintSHA256(pubKey),
					},
				}, nil
			}
			return nil, fmt.Errorf("unknown public key for %q", c.User())
		},
	}

	private, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("Failed to parse private key: ", err)
	}

	config.AddHostKey(private)

	// Once a ServerConfig has been configured, connections can be
	// accepted.
	listener, err := net.Listen("tcp", "0.0.0.0:2022")
	if err != nil {
		return fmt.Errorf("failed to listen for connection: %s", err)
	}

	fmt.Println("Starting Server...")
	for {
		nConn, err := listener.Accept()
		if err != nil {
			fmt.Printf("failed to accept incoming connection: %s\n", err)
		}

		// Before use, a handshake must be performed on the incoming
		// net.Conn.
		conn, _, _, err := ssh.NewServerConn(nConn, config)
		if err != nil {
			fmt.Printf("failed to handshake: %s\n", err)
		}
		if conn != nil && conn.Permissions != nil {
			log.Printf("logged in with key %s", conn.Permissions.Extensions["pubkey-fp"])
		}
	}

}
