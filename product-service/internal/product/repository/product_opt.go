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

type UpdateProductFilter struct {
	ID       string
	Quantity string
}
type Repository interface {
	CreateProduct(ctx context.Context, opt CreateProductOption) (model.Product, error)
	DetailProduct(ctx context.Context, id string) (model.Product, error)
	UpdateProduct(ctx context.Context, product model.Product) (bool, error)
}
