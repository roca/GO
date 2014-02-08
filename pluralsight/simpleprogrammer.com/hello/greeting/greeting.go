package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {

	for _, s := range salutation {
		message, alternate := CreateMessage(s.Name, s.Greeting)

		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}

}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	prefixMap["Joe"] = "Jr "

	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "

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
