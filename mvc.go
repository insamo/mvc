package mvc

import (
	"github.com/insamo/mvc/core"
	"github.com/insamo/mvc/datasource/database"
	"github.com/insamo/mvc/datasource/nosql"
	"github.com/insamo/mvc/logger"
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/insamo/mvc/web/middleware/error_handler"
	"github.com/insamo/mvc/web/middleware/identity"
	"github.com/insamo/mvc/web/middleware/recover"
	"github.com/insamo/mvc/web/routes"
)

// NewMVC configurate and run
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
		identity.Configure,

		logger.ConfigureApplicationLogger,
		logger.ConfigureRequestLogger,
		logger.ConfigureDatabaseLogger,

		database.Configure,
		nosql.Configure,

		// Be after all
		routes.Configure,
		error_handler.Configure,
		recover.Configure,
	)

	return app
}
