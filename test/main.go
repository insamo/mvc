package main

import (
	"context"
	"fmt"

	"github.com/go-kivik/kivik"
	"github.com/insamo/mvc"
)

func main() {
	app := mvc.NewMVC()

	app.Configure()

	defer app.Close()

	txExtra := app.CoachFactory["main"].BeginNewTransaction()

	db := txExtra.DataSource("catalog").(*kivik.DB)

	//
	//db, err := client.DB(context.TODO(), "catalog")
	//if err != nil {
	//	golog.Errorf("Failed connect to database: %s \n", err)
	//	fmt.Errorf("Failed connect to database: %s \n", err)
	//}
	//
	doc := map[string]interface{}{
		"_id":  "username",
		"name": "Insamo",
	}

	_, err := db.Put(context.TODO(), "1", doc)

	fmt.Println(err)

	app.Listen()
}
