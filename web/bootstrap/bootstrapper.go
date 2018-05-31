package bootstrap

import (
	"time"

	"github.com/insamo/mvc/core"
	"github.com/insamo/mvc/datasource"

	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName            string
	AppOwner           string
	AppSpawnDate       time.Time
	ApplicationLogFile *os.File
	RequestLogFile     *os.File
	DatabaseLogFile    *os.File
	RequestContext     *context.Handler
	TxFactory          map[string]datasource.TransactionFactory
	Environment        core.Config
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

	//var refreshEvery = 10 * time.Second

	//// same as:
	//b.Use(func(ctx iris.Context) {
	//	now := time.Now()
	//	if modified, err := ctx.CheckIfModifiedSince(now.Add(-refreshEvery)); !modified && err == nil {
	//		ctx.WriteNotModified()
	//		return
	//	}
	//
	//	ctx.SetLastModified(now)
	//
	//	ctx.Next()
	//})

	// Initialize transaction map
	b.TxFactory = make(map[string]datasource.TransactionFactory)

	return b
}

func (b *Bootstrapper) Close() {
	for _, instance := range b.Environment.DatabaseInstances() {
		b.TxFactory[instance].Close()
	}
	b.RequestLogFile.Close()
	b.DatabaseLogFile.Close()
	b.ApplicationLogFile.Close()
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen() {
	addr := b.Environment.Addr("main")
	cfgs := iris.WithConfiguration(iris.TOML(".iris.toml"))

	b.Run(iris.Addr(addr), cfgs)
}
