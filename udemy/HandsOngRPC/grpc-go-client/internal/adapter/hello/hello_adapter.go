package hello

import "grpc-go-client/internal/port"

type HelloAdapter struct {
	helloClient port.HelloClientPort
}
