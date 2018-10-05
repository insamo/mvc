package main

import (
	"fmt"

	"github.com/couchbase/gocb"

	"github.com/insamo/mvc"
)

func main() {
	app := mvc.NewMVC()

	app.Configure()

	defer app.Close()

	txExtra := app.NxFactory["couchbase"].BeginNewTransaction()

	db := txExtra.DataSource("schedule-test").(*gocb.Bucket)

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

	_, err := db.Upsert("test", doc, 0)

	fmt.Println(err)

	app.Listen()
}
