package main

import (
	"fmt"

	"04_02b/business"

	"golang.org/x/exp/slices"
)

// main is our simple "playground" for the course.
// Note, that in production code, it is a good practice to keep the main function short.
func main() {
	// Create three different energy offers of kineteco
	solar2k := business.Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := business.Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := business.Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Printf(solar2k.Print())
	fmt.Printf(solar3k.Print())
	fmt.Printf(windwest.Print())
	fmt.Printf("our first generic function: %s\n", business.PrintGeneric(solar2k))
	fmt.Printf("our first generic function with wind: %s\n", business.PrintGeneric(windwest))

	// Print functions for 01_04
	ss := []business.Solar{solar2k, solar3k}
	business.PrintSlice[business.Solar](ss)
	business.PrintSlice[business.Wind]([]business.Wind{windwest, windwest})

	// Print functions for 03_02
	ss2 := business.SolarSlice{solar2k, solar3k}
	business.PrintSlice(ss2)
	business.PrintSlice2(ss2)

	fmt.Printf("index: %d\n",slices.Index(ss, solar2k))
	business.SortByCost(ss)
	fmt.Printf("index: %d\n",slices.Index(ss, solar2k))

}
