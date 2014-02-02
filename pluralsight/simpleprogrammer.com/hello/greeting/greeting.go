package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if isFormal {
		do(message)
	}
	do(alternate)

}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	//fmt.Println(len(greeting))
	message, alternate = greeting+" "+name, "HEY! "+name
	return
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
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
