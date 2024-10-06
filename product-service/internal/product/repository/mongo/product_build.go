package mongo

import (
	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/product/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo implRepository) buildProductModel(opt repository.CreateProductOption) model.Product {
	s := model.Product{
		ID:       primitive.NewObjectID(),
		Name:     opt.Name,
		Quantity: opt.Quantity,
	}
	return s
}
