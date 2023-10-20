package basic

import (
	"log"
	"project01/hello/protogen"
)

func BasicHello() {
	h := hello.Hello{
		Name: "Clark Kent",
	}
	log.Println(&h)
}
