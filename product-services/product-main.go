package productMain

import (
	"github.com/buaazp/fasthttprouter"
	productAPI "github.com/j-ew-s/ms-curso-catalog-api/product-services/api"
	"github.com/j-ew-s/ms-curso-catalog-api/shared"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/", shared.AuthSessionValidator(productAPI.GetProducts))

}
