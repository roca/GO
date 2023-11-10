package application

import "fmt"

type HelloService struct{}

func (a *HelloService) GenerateHello(name string) string {
	return "Hello " + name
}

func (a *HelloService) GenerateManyHellos(name string, count int) []string {
	var hellos []string
	for i := 0; i < count; i++ {
		hellos = append(hellos, fmt.Sprintf("[%d] %s", i,a.GenerateHello(name)))
	}
	return hellos
}