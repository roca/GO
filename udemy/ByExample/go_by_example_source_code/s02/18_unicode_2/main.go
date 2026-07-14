// File name: ...\s02\18_unicode_2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	fmt.Println("여보세요")
	fmt.Println("\"OK\"")
	fmt.Println(`\"OK\"`)
	fmt.Println(`"OK"`)

	fmt.Println(`
	<!DOCTYPE html>
	<html>
	<body>	
		<h1>My First Heading</h1>
	</body>
	</html>`)
}
