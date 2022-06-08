package userService

import (
	"github.com/buaazp/fasthttprouter"
	controller "github.com/j-ew-s/ms-curso-catalog-api/user-services/controller"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/token/:token", controller.GetToken)
	//router.GET("/:id", recipescontroller.Create)
}
