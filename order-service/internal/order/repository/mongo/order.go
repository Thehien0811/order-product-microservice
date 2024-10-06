package mongo

import (
	"context"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/order/repository"
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
