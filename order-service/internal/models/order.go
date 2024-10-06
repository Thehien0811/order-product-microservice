package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id"`
	ProductID string             `bson:"product_id"`
	Quantity  int32              `bson:"quantity"`
}
