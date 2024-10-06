package repository

import (
	"context"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
)

type CreateProductOption struct {
	ID       string
	Name     string
	Quantity int32
}

type Repository interface {
	CreateProduct(ctx context.Context, opt CreateProductOption) (model.Product, error)
	DetailProduct(ctx context.Context, id string) (model.Product, error)
}
