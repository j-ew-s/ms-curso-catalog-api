package productMain

import (
	"github.com/buaazp/fasthttprouter"
	productAPI "github.com/j-ew-s/ms-curso-catalog-api/product-services/api"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/", productAPI.GetProducts)

}
