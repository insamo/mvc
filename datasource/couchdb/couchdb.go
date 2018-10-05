package couchdb

import (
	"fmt"

	_ "github.com/go-kivik/couchdb" // init
	"github.com/go-kivik/kivik"
	"github.com/insamo/mvc/datasource/transactions/nosql"
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/kataras/golog"
)

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {

	instances := b.Environment.NoSqlInstances()

	for _, instance := range instances {
		driver := b.Environment.NoSql(instance).GetString("driver")

		if driver != "couch" {
			continue
		}

		dsn := b.Environment.NoSql(instance).GetString("host") + ":" +
			b.Environment.NoSql(instance).GetString("port")

		c, err := kivik.New(driver, dsn)

		if err != nil {
			golog.Errorf("Failed connect to nosql server: %s \n", err)
			fmt.Errorf("Failed connect to nosql server: %s \n", err)
		}

		b.NxFactory[instance] = nosql.NewTransactionFactory(c, nil)
	}
}
