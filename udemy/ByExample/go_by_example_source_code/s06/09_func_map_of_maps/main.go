// File name: ...\s06\09_func_map_of_maps\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	// capitalCity := map[string]string{
	// 	"Canada":  "Ottawa",
	// 	"France":  "Paris",
	// 	"Germany": "Berlin",
	// }
	// fmt.Println(capitalCity["Canada"])

	employees := map[string]map[string]string{

		"BT": map[string]string{
			"firstName": "Blake",
			"lastName":  "Travis",
		},
		"PC": map[string]string{
			"firstName": "Parker",
			"lastName":  "Cooper",
		},
		"DC": map[string]string{
			"firstName": "Dakota",
			"lastName":  "Carrington",
		},
	}

	if emp, ok := employees["PC"]; ok {
		fmt.Println(emp["firstName"], emp["lastName"])
	}

	for initials, emp := range employees {
		fmt.Println(initials, emp["firstName"], emp["lastName"])
	}
}
