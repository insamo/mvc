package etag

import (
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

const IfNoneMatchHeaderKey = "If-None-Match"

// New returns a new handler which adds some headers
// describing the application, i.e the owner, the startup time.
func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		etag := ctx.Values().GetString(context.ETagHeaderKey)
		ctx.Header(context.ETagHeaderKey, etag)
		if match := ctx.GetHeader(IfNoneMatchHeaderKey); match == etag {
			ctx.WriteNotModified()
			return
		}
		ctx.Next()
	}
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	b.DoneGlobal(h)
}
