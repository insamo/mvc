package routes

import (
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/kataras/iris"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.Get("/ping", func(context iris.Context) {
		context.JSON(iris.Map{
			"app":     b.AppName,
			"status":  context.GetStatusCode(),
			"message": "pong",
		})
	})
}
