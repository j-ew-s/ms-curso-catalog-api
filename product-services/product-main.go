package productMain

import (
	"github.com/buaazp/fasthttprouter"
	productAPI "github.com/j-ew-s/ms-curso-catalog-api/product-services/api"
	sessionServices "github.com/j-ew-s/ms-curso-catalog-api/session-services"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/", sessionServices.AuthSessionValidator(productAPI.GetProducts))

}
