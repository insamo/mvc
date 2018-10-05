package database

import (
	"bufio"
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"time"
	"unicode"

	"github.com/insamo/mvc/datasource"
	"github.com/insamo/mvc/datasource/transactions/sql"

	_ "github.com/insamo/mvc/dialect/mssql2008" // init
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // init
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/kataras/golog"
)

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
			l := DatabaseLogger{log.New(b.DatabaseLogFile, "\r\n", 0)}
			db.SetLogger(l)
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

		b.TxFactory[instance] = sql.NewTransactionFactory(db, queries)
	}
}

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

var (
	defaultLogger            = DatabaseLogger{log.New(os.Stdout, "\r\n", 0)}
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

var NowFunc = func() time.Time {
	return time.Now()
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

var DatabaseLogFormatter = func(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
			currentTime     = NowFunc().Format("2006-01-02 15:04:05")
			source          = fmt.Sprintf("(%v)", values[1])
		)

		messages = []interface{}{source, currentTime}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf(" [%.2fms] \n", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, sql)
			messages = append(messages, fmt.Sprintf(" [%v] ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned "))
		} else {
			messages = append(messages, "")
			messages = append(messages, values[2:]...)
			messages = append(messages, "")
		}
	}

	return
}

type DatabaseLoggerInterface interface {
	Print(v ...interface{})
}

// LogWriter log writer interface
type LogWriter interface {
	Println(v ...interface{})
}

// Logger default logger
type DatabaseLogger struct {
	LogWriter
}

// Print format & print log
func (databaseLogger DatabaseLogger) Print(values ...interface{}) {
	databaseLogger.Println(DatabaseLogFormatter(values...)...)
}
