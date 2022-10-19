package models

import (
	"time"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Brand       string             `bson:"brand"`
	Price       float64            `bson:"Price"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
