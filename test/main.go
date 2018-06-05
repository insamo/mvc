package main

import (
	"github.com/insamo/mvc"
)

func main() {
	app := mvc.NewMVC()

	app.Listen()
}
