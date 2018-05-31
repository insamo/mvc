package identity

import (
	"time"

	"github.com/insamo/mvc/web/bootstrap"

	"github.com/kataras/iris"
)

// New returns a new handler which adds some headers
// describing the application, i.e the owner, the startup time.
func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		// response headers
		//TODO set ACCEPT
		ctx.Header("App-Name", b.AppName)
		ctx.Header("App-Owner", b.AppOwner)
		ctx.Header("App-Since", time.Since(b.AppSpawnDate).String())
		// TODO set from config contenttype
		ctx.ContentType("application/json; charset=UTF-8")
		ctx.Next()
	}
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	b.UseGlobal(h)
}
