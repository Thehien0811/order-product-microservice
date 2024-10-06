package mongo

import (
	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/order/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo implRepository) buildOrderModel(opt repository.CreateOrderOption) model.Order {
	s := model.Order{
		ID:       primitive.NewObjectID(),
		Name:     opt.Name,
		Quantity: opt.Quantity,
	}
	return s
}
