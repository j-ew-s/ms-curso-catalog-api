package api

import (
	"encoding/json"
	"fmt"
	"time"

	productModel "github.com/j-ew-s/ms-curso-catalog-api/product-services/models"
	userGRPC "github.com/j-ew-s/ms-curso-catalog-api/user-services/grpc"
	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func GetProducts(ctx *fasthttp.RequestCtx) {

	token := string(ctx.Request.Header.Peek("user-token"))

	fmt.Println(fmt.Sprintf("Token a ser buscado: %s", token))

	resp, err := userGRPC.IsTokenValid(token)
	httpStatusCode := 200

	if err != nil {
		fmt.Println(fmt.Sprintf("gRPC DIAL falhou: %v", err))
		httpStatusCode = 500
		PrepareResponse(ctx, httpStatusCode, nil)
		return
	}
	product := &productModel.Product{}
	fmt.Println(fmt.Sprintf("Resposta: %v", resp))

	if resp.Status == true {
		product.Id = "123890"
		product.Name = "Shoes"
		product.Price = 100.30
	} else if resp.Status == false {
		httpStatusCode = 401
	}

	PrepareResponse(ctx, httpStatusCode, product)

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
