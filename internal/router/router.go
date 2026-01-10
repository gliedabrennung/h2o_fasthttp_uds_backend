package router

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func InitRouter() *router.Router {
	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		_, err := fmt.Fprintf(ctx, "Fasthttp Http Server")
		if err != nil {
			return
		}
	})
	return r
}
