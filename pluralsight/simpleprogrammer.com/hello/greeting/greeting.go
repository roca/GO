package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool, times int) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	i := 0
	for {
		if i >= times {
			break
		}

		if i%2 == 0 {
			i++
			continue
		}
		if prefix := GetPrefix(salutation.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}

		i++
	}

}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Bob":
		prefix = "Mr "
	case name == "Joe", name == "Amy", len(name) == 10:
		prefix = "Dr "
	case name == "Mary":
		prefix = "Mrs "
	default:
		prefix = "Dude "
	}

	return
}

func TypeSwitchTest(x interface{}) {

	switch t := x; x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutaion")
	default:
		fmt.Print("Unknown type: ")
		fmt.Println(t)
	}
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
