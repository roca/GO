package basic

import (
	"log"
	hello "project01/hello/protogen"
)

func BasicHello() {
	h := hello.Hello{
		Name: "Clark Kent",
	}
	log.Println(&h)
}
