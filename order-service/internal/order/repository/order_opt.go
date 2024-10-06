package repository

import (
	"context"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
)

type CreateOrderOption struct {
	ID       string
	ProductID     string
	Quantity int32
}

type Repository interface {
	CreateOrder(ctx context.Context, opt CreateOrderOption) (model.Order, error)
	DetailOrder(ctx context.Context, id string) (model.Order, error)
}
