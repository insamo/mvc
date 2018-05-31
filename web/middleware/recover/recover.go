package recover

import (
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/kataras/iris/middleware/recover"
)

// New returns a new handler which adds some headers
// describing the application, i.e the owner, the startup time.
func New(b *bootstrap.Bootstrapper) {
	b.Use(recover.New())
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	New(b)
}
