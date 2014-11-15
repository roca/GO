// exe_1a
package main

import (
	"fmt"
)

func main() {

	i := 20

	addr := &i

	fmt.Println(&i, addr)
}
