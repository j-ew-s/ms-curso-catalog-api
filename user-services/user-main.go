package userService

import (
	"github.com/buaazp/fasthttprouter"
	userApi "github.com/j-ew-s/ms-curso-catalog-api/user-services/api"
)

func SetRoutes(router *fasthttprouter.Router) {

	router.GET("/token/:token", userApi.GetToken)
	//router.GET("/:id", recipescontroller.Create)
}
