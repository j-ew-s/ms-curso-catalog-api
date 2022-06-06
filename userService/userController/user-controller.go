package userController

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/j-ew-s/ms-curso-catalog-api/userServiceGRPC"
	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func GetToken(ctx *fasthttp.RequestCtx) {

	id := fmt.Sprintf("%v", ctx.UserValue("token"))

	resp, err := userServiceGRPC.GetByUserId(id)

	httpStatusCode := 200

	if err != nil {
		httpStatusCode = 500
		PrepareResponse(ctx, httpStatusCode, nil)
		return
	}

	PrepareResponse(ctx, httpStatusCode, resp)

}

func PrepareResponse(ctx *fasthttp.RequestCtx, code int, response interface{}) {

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	ctx.Response.SetStatusCode(code)

	start := time.Now()

	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), response)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}
