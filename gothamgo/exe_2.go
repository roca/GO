// exe_2
package main

import (
	"fmt"
)

type Speaker interface {
	SayHello() string
}

type English struct {
}

type Chinese struct {
}

func (e *English) SayHello() string {
	return "Hello World"
}

func (c *Chinese) SayHello() string {
	return "你好世界"
}

func SayHello(s Speaker) {
	fmt.Println(s.SayHello())
}
func main() {

	Speaker1 := &English{}
	fmt.Println(Speaker1.SayHello())

	Speaker2 := &Chinese{}
	fmt.Println(Speaker2.SayHello())
}
