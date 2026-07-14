package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const delimiter = '\n'

func main() {
	fmt.Print("\nEnter your strings: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(delimiter)

	if err != nil {
		fmt.Println(err)
		return
	}

	// convert sting
	input = strings.Replace(input, "\n", "", -1)

	fmt.Println(input)
	inputstring := strings.ToLower(strings.TrimSpace(input)) //make input to lower so that we can find any string not matter upper or lower

	if strings.HasPrefix(inputstring, "i") && strings.Contains(inputstring, "a") && strings.HasSuffix(inputstring, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
