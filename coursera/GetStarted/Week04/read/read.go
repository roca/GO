/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Person : Each field will be a string of size 20 (characters).
type Person struct {
	fname [20]byte
	lname [20]byte
}

// People ...
type People []Person

func main() {

	var filePath string

	fmt.Print("Please enter the full path name of the text file : ")
	fmt.Scan(&filePath)

	people, err := ReadPeopleFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	for _, person := range people {
		// %s format will convert the array of bytes to a string
		fmt.Printf("FirstName: %s , LastName: %s\n", person.fname, person.lname)
	}

}

// ReadPeopleFromFile : reads people from a file
func ReadPeopleFromFile(filePath string) (People, error) {
	people := People{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		if len(values) > 1 {
			firstName := makeByteArrayFromString(values[0])
			lastName := makeByteArrayFromString(values[1])

			people = append(people, Person{fname: firstName, lname: lastName})
		}
	}

	return people, nil
}

func makeByteArrayFromString(str string) [20]byte {
	var arr [20]byte
	for k, v := range []byte(str) {
		arr[k] = byte(v)
	}
	return arr
}
