package mvc

import (
	"github.com/insamo/mvc/core"
	"github.com/insamo/mvc/logger"
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/insamo/mvc/web/middleware/database"
	"github.com/insamo/mvc/web/middleware/error_handler"
	"github.com/insamo/mvc/web/middleware/identity"
	"github.com/insamo/mvc/web/middleware/recover"
	"github.com/insamo/mvc/web/routes"
)

func NewMVC(cfgs ...bootstrap.Configurator) *bootstrap.Bootstrapper {
	// Load environment
	env := core.NewConfig()

	// Create app
	app := bootstrap.New(
		env.Server("main").GetString("name"),
		env.Server("main").GetString("owner"),
		cfgs...,
	)

	app.Bootstrap(env)

	app.Configure(
		//etag.Configure,
		// Middlewares
		//jwt.Configure,
		identity.Configure,
		//cors.Configure,

		logger.ConfigureApplicationLogger,
		logger.ConfigureRequestLogger,
		logger.ConfigureDatabaseLogger,

		database.Configure,

		// Be after all
		routes.Configure,
		error_handler.Configure,
		recover.Configure,
	)

	return app
}

//func main() {
//	// PrintVersion
//	printVersionInfo()
//
//	env := core.NewConfig()
//
//	app := newApp(env)
//	defer app.Close()
//
//	// Start server
//	app.Listen(
//		env.Addr("main"),
//		iris.WithConfiguration(iris.TOML(".iris.toml")),
//	)
//
//	// Gracefull shutdown
//	//go func() {
//	//	ch := make(chan os.Signal, 1)
//	//	signal.Notify(ch,
//	//		// kill -SIGINT XXXX or Ctrl+c
//	//		os.Interrupt,
//	//		syscall.SIGINT, // register that too, it should be ok
//	//		// os.Kill  is equivalent with the syscall.Kill
//	//		os.Kill,
//	//		syscall.SIGKILL, // register that too, it should be ok
//	//		// kill -SIGTERM XXXX
//	//		syscall.SIGTERM,
//	//	)
//	//	select {
//	//	case <-ch:
//	//		println("shutdown...")
//	//
//	//		timeout := 5 * time.Second
//	//		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
//	//		defer cancel()
//	//
//	//		app.Shutdown(ctx)
//	//	}
//	//}()
//}
