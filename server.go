package main

import (
	"github.com/buaazp/fasthttprouter"
	catalogtSerivces "github.com/j-ew-s/ms-curso-catalog-api/catalog-services"
	database "github.com/j-ew-s/ms-curso-catalog-api/database"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	dbConn := SetDataBase()

	catService := catalogtSerivces.CatalogServiceMain(&dbConn)
	catalogtSerivces.SetRoutes(router, catService)

	fasthttp.ListenAndServe(":5200", CORS(router.Handler))
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

func SetDataBase() database.MongoDBConfig {

	conn := database.SetupMongoDBConfig("127.0.0.1", 27017, "mongoadmin", "secret")

	return conn
}
