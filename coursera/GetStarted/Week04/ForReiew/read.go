package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var fileName string
	var namesArray []names
	fmt.Println("Write the name of the file: ")
	fmt.Scan(&fileName)

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading the file: ", err)
	} else {

		dataString := string(data)

		fields := strings.Fields(dataString)

		var newName names
		for i, v := range fields {
			if i%2 == 0 {
				newName.lName = v
			} else {
				newName.fName = v
			}

			if newName.lName != "" && newName.fName != "" {
				namesArray = append(namesArray, newName)
				newName.fName = ""
				newName.lName = ""
			}

		}

		for i, v := range namesArray {
			fmt.Println(i, " --> ", v)
		}
	}
}

type names struct {
	fName string
	lName string
}
