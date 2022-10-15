package api

import (
	models "github.com/j-ew-s/ms-curso-catalog-api/catalog/models"
	"github.com/j-ew-s/ms-curso-catalog-api/common"
	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func GetProducts(ctx *fasthttp.RequestCtx) {

	// 1  token := string(ctx.Request.Header.Peek("user-token"))

	// 2 fmt.Println(fmt.Sprintf("Token a ser buscado: %s", shared.GlobalSession.Token))

	// 2 resp, err := userGRPC.IsTokenValid(shared.GlobalSession.Token)
	httpStatusCode := 200

	/* 2
	if err != nil {
		fmt.Println(fmt.Sprintf("gRPC DIAL falhou: %v", err))
		httpStatusCode = 500
		PrepareResponse(ctx, httpStatusCode, nil)
		return
	}*/
	// 2 fmt.Println(fmt.Sprintf("Resposta: %v", resp))
	product := &models.Product{}

	// 2if resp.Status == true {
	product.Id = "123890"
	product.Name = "Shoes"
	product.Price = 100.30
	// 2} else if resp.Status == false {
	httpStatusCode = 401
	// 2}

	// added the shared
	common.PrepareResponse(ctx, httpStatusCode, product)

}

// 3

// func PrepareResponse(ctx *fasthttp.RequestCtx, code int, response interface{}) {

// 	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

// 	ctx.Response.SetStatusCode(code)

// 	start := time.Now()

// 	if err := json.NewEncoder(ctx).Encode(response); err != nil {
// 		elapsed := time.Since(start)
// 		fmt.Println(" ERROR : ", elapsed, err.Error(), response)
// 		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
// 	}

// }
