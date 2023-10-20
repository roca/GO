package basic

import (
	"log"
	"project01/hello"
)

func BasicHello() {
	h := hello.Hello{
		Name: "Clark Kent",
	}
	log.Println(&h)
}
