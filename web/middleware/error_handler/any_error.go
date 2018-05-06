package error_handler

import (
	"bitbucket.org/insamo/mvc/web/bootstrap"

	"github.com/kataras/iris"
)

// New returns a new handler which adds some headers
// describing the application, i.e the owner, the startup time.
func New(b *bootstrap.Bootstrapper) {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}
		ctx.JSON(err)
		return
	})
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	New(b)
}
