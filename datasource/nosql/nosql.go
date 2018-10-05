package nosql

import (
	"fmt"

	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/insamo/mvc/datasource/nosql/transactions"
	"github.com/insamo/mvc/web/bootstrap"
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

		fmt.Print("Test connection to server coach " + instance + " ")
		golog.Debugf("Test connection to database coach" + instance + " success!")
		fmt.Print("connected \n")

		client, err := kivik.New(driver, dsn)

		if err != nil {
			golog.Errorf("Failed connect to nosql server: %s \n", err)
			fmt.Errorf("Failed connect to nosql server: %s \n", err)
		}

		b.CoachFactory[instance] = nosql.NewTransactionFactory(client, nil)
	}
}
