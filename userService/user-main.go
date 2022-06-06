package userService

import (
	"github.com/buaazp/fasthttprouter"
	userService "github.com/j-ew-s/ms-curso-catalog-api/userService/controllers"
)

func setRoutes(router *fasthttprouter.Router) {

	router.
		router.GET("/token/:token", userService.GetToken)
	//router.GET("/:id", recipescontroller.Create)
}
