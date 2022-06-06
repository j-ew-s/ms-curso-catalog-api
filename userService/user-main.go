package userService

import (
	"github.com/buaazp/fasthttprouter"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/token/:token", GetToken)
	//router.GET("/:id", recipescontroller.Create)
}
