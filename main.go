package main

import (
	"log"

	"github.com/gobuffalo/gothrecipe/actions"
)

func main() {
	app := actions.App()
	log.Fatal(app.Serve())
}
