package userService

import (
	"github.com/buaazp/fasthttprouter"
	userController "github.com/j-ew-s/ms-curso-catalog-api/userService"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/token/:token", userController.GetToken)
	//router.GET("/:id", recipescontroller.Create)
}
