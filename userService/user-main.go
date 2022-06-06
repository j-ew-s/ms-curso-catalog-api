package userService

import (
	"github.com/buaazp/fasthttprouter"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/token/:token", userController.GetToken)
	//router.GET("/:id", recipescontroller.Create)
}
