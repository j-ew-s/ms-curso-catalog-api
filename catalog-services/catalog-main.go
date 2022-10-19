package catalogtSerivces

import (
	"github.com/buaazp/fasthttprouter"
	authServices "github.com/j-ew-s/ms-curso-catalog-api/auth-services"
	"github.com/j-ew-s/ms-curso-catalog-api/catalog-services/api"
	database "github.com/j-ew-s/ms-curso-catalog-api/database"
)

type Server struct {
	*api.CatalogService
}

func CatalogServiceMain(conn *database.MongoDBConfig) *api.CatalogService {
	catalogService := &api.CatalogService{
		DBConn: conn,
	}

	return catalogService
}

func SetRoutes(router *fasthttprouter.Router, client *api.CatalogService) {

	router.GET("/", authServices.AuthSessionValidator(client.GetProducts))
	router.GET("/:id", authServices.AuthSessionValidator(client.GetProductById))
	router.POST("/", authServices.AuthSessionValidator(client.InsertProduct))

}
