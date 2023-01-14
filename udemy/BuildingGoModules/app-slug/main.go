package main

import (
	"log"

	"github.com/roca/GO/tree/staging/udemy/BuildingGoModules/toolkit"
)

func main() {
	toSlug := "NOW!!!? is the time 123"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(slugified)
}
