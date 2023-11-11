package application

import (
	"fmt"
	"log"
	"strings"
)

type HelloService struct{}

func (a *HelloService) GenerateHello(name string) string {
	return "Hello " + name
}

func (a *HelloService) GenerateManyHellos(name string, count int) []string {
	var hellos []string
	for i := 0; i < count; i++ {
		hellos = append(hellos, fmt.Sprintf("[%d] %s", i, a.GenerateHello(name)))
	}
	return hellos
}

func (a *HelloService) GenerateHelloToEveryone(names []string) string {
	var hellos []string
	for _, name := range names {
		hellos = append(hellos, a.GenerateHello(name))
	}
	log.Println(hellos)
	return strings.Join(hellos, ", ")
}
