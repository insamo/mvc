package jwt

import (
	"bitbucket.org/insamo/mvc/web/bootstrap"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

// New returns a new handler which adds jwt middleware
func New(b *bootstrap.Bootstrapper) *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(b.Environment.Core().GetString("secure.key")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		Debug:         true,
	})
	return jwtHandler
}

// Configure creates a new jwt middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	b.UseGlobal(h.Serve)
}
