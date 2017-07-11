package main

import (
	"log"

	"github.com/GOCODE/buffalo/hello_go/actions"
	"github.com/gobuffalo/envy"
)

func main() {
	port := envy.Get("PORT", "3000")
	app := actions.App()
	log.Fatal(app.Start(port))
}
