package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	CreateHTTPServer()
	//res, _ := userService.GetByUserId("1")
	//fmt.Println(res)
}

func CreateHTTPServer() {

	router := fasthttprouter.New()

	userService.setRoutes(&router)

	setRoutes(router)

	fasthttp.ListenAndServe("5002", CORS(router.Handler))

}

// Configure all routes for controllers
func setRoutes(router *fasthttprouter.Router) {

	router.GET("/", recipescontroller.Ping)
	router.GET("/:id", recipescontroller.Create)
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
