package couchbase

import (
	"fmt"

	"github.com/couchbase/gocb"
	"github.com/insamo/mvc/datasource/transactions/nosql"
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/kataras/golog"
)

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {

	instances := b.Environment.NoSqlInstances()

	for _, instance := range instances {
		driver := b.Environment.NoSql(instance).GetString("driver")

		if driver != "couchbase" {
			continue
		}

		c, err := gocb.Connect(b.Environment.NoSql(instance).GetString("host"))

		if err != nil {
			golog.Errorf("Cluster connect with error %v", err)
			fmt.Errorf("Cluster connect with error %v", err)
		}

		err = c.Authenticate(gocb.PasswordAuthenticator{
			Username: b.Environment.NoSql(instance).GetString("username"),
			Password: b.Environment.NoSql(instance).GetString("password"),
		})
		if err != nil {
			golog.Errorf("Cluster auth with error %v", err)
			fmt.Errorf("Cluster auth with error %v", err)
		}

		b.NxFactory[instance] = nosql.NewTransactionFactory(c, nil)
	}
}
