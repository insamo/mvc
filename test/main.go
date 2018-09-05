package main

import (
	"github.com/insamo/mvc"
)

func main() {
	app := mvc.NewMVC()

	app.Configure()

	defer app.Close()

	app.Listen()
}
