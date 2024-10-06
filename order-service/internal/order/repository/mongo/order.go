package mongo

import (
	"context"
	"log"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/order/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	order = "orders"
)

func (repo implRepository) getOrderCollection() mongo.Collection {
	return *repo.db.Collection(order)
}

func (repo implRepository) CreateOrder(ctx context.Context, opt repository.CreateOrderOption) (model.Order, error) {
	col := repo.getOrderCollection()

	s := repo.buildOrderModel(opt)

	_, err := col.InsertOne(ctx, s)
	if err != nil {
		return model.Order{}, err
	}

	return s, nil
}

func (repo implRepository) DetailOrder(ctx context.Context, id string) (model.Order, error) {
	col := repo.getOrderCollection()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	cur, err := col.Find(ctx, bson.M{"_id": objID})
	if err != nil {
		return model.Order{}, err
	}
	var orders []model.Order
	err = cur.All(ctx, &orders)
	if err != nil {
		return model.Order{}, err
	}

	return orders[0], nil

}
