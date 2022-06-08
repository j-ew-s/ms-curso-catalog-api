package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/j-ew-s/ms-curso-catalog-api/user-services/controller"
	"github.com/valyala/fasthttp"
)

func main() {
	CreateHTTPServer()
}

func CreateHTTPServer() {

	router := fasthttprouter.New()

	controller.GetToken(router)

	fasthttp.ListenAndServe("5002", CORS(router.Handler))

}

var (
	corsAllowHeaders     = "*"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

// CORS handles CORS
func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}
