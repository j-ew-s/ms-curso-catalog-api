package catalogtMain

import (
	"github.com/buaazp/fasthttprouter"
	authServices "github.com/j-ew-s/ms-curso-catalog-api/auth-services"
	catalogAPI "github.com/j-ew-s/ms-curso-catalog-api/catalog/api"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/", authServices.AuthSessionValidator(catalogAPI.GetProducts))
	router.GET("/:id", authServices.AuthSessionValidator(catalogAPI.GetProducts))

}
