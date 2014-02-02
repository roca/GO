package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")

	do(message)
	do(alternate)

}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	//fmt.Println(len(greeting))
	message, alternate = greeting[1]+" "+name, "HEY! "+name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func PrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
func main() {

	var s = Salutation{"Bob", "Hello"}
	Greet(s, CreatePrintFunction("!!!"))
}
