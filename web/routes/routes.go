package routes

import (
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/kataras/iris"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	b.Get("/ping", func(context iris.Context) {
		return
	})

	//b.Get("/generate", func(context iris.Context) {
	//	mySigningKey := []byte("My Secret")
	//
	//	type MyCustomClaims struct {
	//		AppName string `json:"foo"`
	//		jwt.StandardClaims
	//	}
	//
	//	// Create the Claims
	//	claims := MyCustomClaims{
	//		"kpreact",
	//		jwt.StandardClaims{
	//			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	//			Issuer:    "test",
	//		},
	//	}
	//
	//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//	ss, err := token.SignedString(mySigningKey)
	//	fmt.Printf("%v %v", ss, err)
	//
	//	context.WriteString(ss)
	//})
}
