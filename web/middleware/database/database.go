package database

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"bitbucket.org/insamo/mvc/datasource"
	_ "bitbucket.org/insamo/mvc/dialect/mssql"
	"bitbucket.org/insamo/mvc/web/bootstrap"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/denisenkom/go-mssqldb"

	"github.com/kataras/golog"
)

func testConnection(driver string, dsn string) bool {
	db, err := gorm.Open(driver, dsn)
	defer db.Close()

	if err != nil {
		fmt.Print("* ")
		return false
	}

	db.Close()

	return true
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {

	instances := b.Environment.DatabaseInstances()

	for _, instance := range instances {

		driver := b.Environment.Database(instance).GetString("driver")
		dsn := b.Environment.DSN(instance)

		fmt.Print("Test connection to database " + instance + " ")
		for testConnection(driver, dsn) == false {
			duration := time.Second * 2
			time.Sleep(duration)
		}
		golog.Debugf("Test connection to database " + instance + " success!")
		fmt.Print("connected \n")

		db, err := gorm.Open(b.Environment.Database(instance).GetString("driver"), b.Environment.DSN(instance))

		if err != nil {
			golog.Errorf("Failed connect to database: %s \n", err)
			fmt.Errorf("Failed connect to database: %s \n", err)
		}

		if err := db.DB().Ping(); err != nil {
			golog.Errorf("Failed ping to database: %s \n", err)
			fmt.Errorf("Failed ping to database: %s \n", err)
		}

		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)

		// Set database debug
		db.LogMode(b.Environment.Database(instance).GetBool("debug"))

		// Set logger if configured
		if b.DatabaseLogFile != nil {
			db.SetLogger(log.New(b.DatabaseLogFile, "", 0))
		}

		// Loading queries
		f, err := os.Open("storage/database/queries.sql")
		if err != nil {
			golog.Errorf("Failed to open queries file: %s \n", err)
			fmt.Errorf("Failed to open queries file: %s \n", err)
		}
		scanner := &datasource.Scanner{}
		queries := scanner.Run(bufio.NewScanner(f))
		f.Close()

		b.TxFactory[instance] = datasource.NewTransactionFactory(db, queries)
	}
}
