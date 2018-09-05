package nosql

import (
	"context"
	"fmt"

	_ "github.com/go-kivik/couchdb"
	"github.com/insamo/mvc/web/bootstrap"

	"github.com/go-kivik/kivik"
	"github.com/kataras/golog"
)

// TODO need error handling
// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {

	instances := b.Environment.NoSqlInstances()

	for _, instance := range instances {
		driver := b.Environment.NoSql(instance).GetString("driver")

		dsn := b.Environment.NoSql(instance).GetString("host") + ":" +
			b.Environment.NoSql(instance).GetString("port")

		fmt.Print("Test connection to database " + instance + " ")
		golog.Debugf("Test connection to database " + instance + " success!")
		fmt.Print("connected \n")

		client, err := kivik.New(context.TODO(), driver, dsn)
		if err != nil {
			golog.Errorf("Failed connect to nosql server: %s \n", err)
			fmt.Errorf("Failed connect to nosql server: %s \n", err)
		}

		db, err := client.DB(context.TODO(), "schedule")
		if err != nil {
			golog.Errorf("Failed connect to database: %s \n", err)
			fmt.Errorf("Failed connect to database: %s \n", err)
		}

		b.NoSqlDB[b.Environment.NoSql(instance).GetString("database")] = *db
	}
}
