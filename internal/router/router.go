package router

import (
	"wsp_go/internal/handlers"

	"github.com/fasthttp/router"
)

func InitRouter() *router.Router {
	r := router.New()
	r.GET("/", handlers.WelcomePage)
	return r
}
