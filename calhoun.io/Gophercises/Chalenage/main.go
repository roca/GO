package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func chdir(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Error:", "Directory not found")
	}
}

func command(cmd string) {
	fmt.Println("Command:", cmd)
	command := strings.Split(cmd, " ")
	exe := exec.Command(command[0], command[1:]...)
	stdout, err := exe.Output()

	if err != nil {
		switch command[0] {
		case "mkdir":
			fmt.Println("Error: ", "Directory already exists")
		default:
			fmt.Println("Error: ", err.Error())
		}
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func main() {
	for {
		var input1, input2 string
		fmt.Scanf("%s %s\n", &input1, &input2)
		switch input1 {
		case "quit":
			os.Exit(1)
		case "pwd":
			command(input1)
		case "ls":
			if input2 == "" {
				command(input1)
			} else {
				command(fmt.Sprintf("%s %s", input1, strings.ToUpper(input2)))
			}
		case "mkdir":
			command(fmt.Sprintf("%s %s", input1, input2))
		case "cd":
			chdir(input2)
		case "touch":
			command(fmt.Sprintf("%s %s", input1, input2))
		default:
			fmt.Println("Command not recognized:", input1)
		}

	}
}
