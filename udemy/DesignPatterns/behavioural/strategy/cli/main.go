package main

import (
	"flag"
	"log"
	"os"

	"example.com/strategy/shapes"
)

var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()

	activeStrategy, err := shapes.CreateShapePrintStrategy(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case shapes.CONSOLE_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("./image.jpg")
		if err != nil {
			log.Fatal("Error opening image file")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
