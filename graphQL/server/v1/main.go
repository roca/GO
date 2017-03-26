package v1

import "github.com/GOCODE/graphQL/server"

func SERVER() {
	server.Version = 1
	server.Start()
}
