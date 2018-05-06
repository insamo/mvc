package etag

import (
	"bitbucket.org/insamo/mvc/web/bootstrap"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

const ifNoneMatchHeaderKey = "If-None-Match"

// New returns a new handler which adds some headers
// describing the application, i.e the owner, the startup time.
func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx context.Context) {
		//fmt.Println(ctx.Values())
		//etag := ctx.Values().GetString("hash")
		//fmt.Println(1)
		//fmt.Println(ctx.GetHeader("test"))
		//key := ctx.Request().URL.Path
		//ctx.Header(context.ETagHeaderKey, etag)
		//if match := ctx.GetHeader(ifNoneMatchHeaderKey); match == etag {
		//	ctx.WriteNotModified()
		//	return
		//}

		//ctx.Next()
	}

}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	b.DoneGlobal(h)
}
