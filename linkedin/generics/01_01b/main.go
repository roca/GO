package main

import (
	"01_01b/business"
	"fmt"
)

// main is our simple "playground" for the course.
// Note, that in production code, it is a good practice to keep the main function short.
func main() {
	// Create three different energy offers of kineteco
	solar2k := business.Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := business.Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := business.Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Println(solar3k.Print())
	fmt.Println(solar2k.Print())
	fmt.Println(windwest.Print())

	fmt.Println(business.PrintGeneric(solar2k))
	fmt.Println(business.PrintGeneric(solar3k))
	fmt.Println(business.PrintGeneric(windwest))
}
