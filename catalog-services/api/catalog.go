package api

import (
	"encoding/json"
	"fmt"

	"github.com/j-ew-s/ms-curso-catalog-api/common"
	"github.com/j-ew-s/ms-curso-catalog-api/database"
	"github.com/j-ew-s/ms-curso-catalog-api/models"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatalogService struct {
	DBConn *database.MongoDBConfig
}

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func (cs CatalogService) GetProducts(ctx *fasthttp.RequestCtx) {

	httpStatusCode := 200
	filter := bson.D{{}}
	products, err := cs.DBConn.GetProducts(filter)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	if len(products) < 1 {
		httpStatusCode = 404
	}

	common.PrepareResponse(ctx, httpStatusCode, products)

}

func (cs CatalogService) GetProductById(ctx *fasthttp.RequestCtx) {

	httpStatusCode := 200
	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	prodId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(fmt.Printf(" No Documents found :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}
	filter := bson.D{primitive.E{Key: "_id", Value: prodId}}

	products, err := cs.DBConn.GetProducts(filter)
	if err != nil {
		fmt.Println(fmt.Printf(" No Documents found :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	if len(products) < 1 {
		httpStatusCode = 404
	}

	common.PrepareResponse(ctx, httpStatusCode, products)

}

func (cs CatalogService) InsertProduct(ctx *fasthttp.RequestCtx) {

	httpStatusCode := 201

	product := &models.Product{}
	err := json.Unmarshal(ctx.PostBody(), &product)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	err = cs.DBConn.InsertProduct(product)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	common.PrepareResponse(ctx, httpStatusCode, nil)

}
