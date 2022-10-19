package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/j-ew-s/ms-curso-catalog-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

func SetupMongoDBConfig(host string, port int, user string, password string) MongoDBConfig {

	mongoDBConfig := &MongoDBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}

	return *mongoDBConfig
}
func (mc MongoDBConfig) GetMongoDBURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/", mc.User, mc.Password, mc.Host, mc.Port)
}

func (mc MongoDBConfig) ProductCollection() mongo.Collection {

	client, err := mongo.NewClient(options.Client().ApplyURI(mc.GetMongoDBURI()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	productCollection := client.Database("catalog").Collection("product")

	return *productCollection
}

func (mc MongoDBConfig) GetProducts(filter interface{}) ([]models.Product, error) {

	products := make([]models.Product, 0)

	productCollection := mc.ProductCollection()
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	defer productCollection.Database().Client().Disconnect(ctx)

	cursor, err := productCollection.Find(ctx, filter)
	if err != nil {
		return products, err
	}

	for cursor.Next(ctx) {

		var p models.Product
		err := cursor.Decode(&p)
		if err != nil {
			return products, err
		}

		products = append(products, p)
	}

	if err := cursor.Err(); err != nil {
		return products, err
	}

	cursor.Close(ctx)

	return products, nil
}

func (mc MongoDBConfig) InsertProduct(product *models.Product) error {

	productCollection := mc.ProductCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer productCollection.Database().Client().Disconnect(ctx)

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := productCollection.InsertOne(ctx, product)

	if err != nil {
		return err
	}

	return nil
}
