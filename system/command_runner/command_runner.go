// command_runner
package main


type Command string

type Executable interface {
	execute()
}

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
