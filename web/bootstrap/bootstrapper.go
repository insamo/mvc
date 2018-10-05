package bootstrap

import (
	"time"

	"github.com/insamo/mvc/core"

	"os"

	"github.com/insamo/mvc/datasource/transactions/nosql"
	"github.com/insamo/mvc/datasource/transactions/sql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/rakanalh/scheduler"
	"github.com/rakanalh/scheduler/storage"
)

// Configurator func
type Configurator func(*Bootstrapper)

// Bootstrapper struct
type Bootstrapper struct {
	*iris.Application
	AppName            string
	AppOwner           string
	AppSpawnDate       time.Time
	ApplicationLogFile *os.File
	RequestLogFile     *os.File
	DatabaseLogFile    *os.File
	RequestContext     *context.Handler
	TxFactory          map[string]sql.TransactionFactory
	NxFactory          map[string]nosql.TransactionFactory
	Environment        core.Config
	Scheduler          scheduler.Scheduler
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap(environment core.Config) *Bootstrapper {
	b.Environment = environment

	// Initialize transaction map
	b.TxFactory = make(map[string]sql.TransactionFactory)
	b.NxFactory = make(map[string]nosql.TransactionFactory)

	// Initialize scheduler
	s := storage.NewMemoryStorage()
	b.Scheduler = scheduler.New(s)
	b.Scheduler.Start()

	return b
}

// Close bootstraper
func (b *Bootstrapper) Close() {
	for _, instance := range b.Environment.DatabaseInstances() {
		b.TxFactory[instance].Close()
	}
	b.RequestLogFile.Close()
	b.DatabaseLogFile.Close()
	b.ApplicationLogFile.Close()
	b.Scheduler.Clear()
	b.Scheduler.Stop()
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen() {
	addr := b.Environment.Addr("main")
	cfgs := iris.WithConfiguration(iris.TOML(".iris.toml"))

	b.Run(iris.Addr(addr), cfgs)
}
