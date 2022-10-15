package catalogtMain

import (
	"github.com/buaazp/fasthttprouter"
	catalogAPI "github.com/j-ew-s/ms-curso-catalog-api/catalog/api"
	sessionServices "github.com/j-ew-s/ms-curso-catalog-api/session-services"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/", sessionServices.AuthSessionValidator(catalogAPI.GetProducts))

}
