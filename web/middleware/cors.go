package cors

import (
	"github.com/insamo/mvc/web/bootstrap"
	"github.com/iris-contrib/middleware/cors"
)

// Configure cross-origin resource sharing middleware
func Configure(b *bootstrap.Bootstrapper) {
	crs := cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
		Debug:            false,
	})
	b.UseGlobal(crs)
}
