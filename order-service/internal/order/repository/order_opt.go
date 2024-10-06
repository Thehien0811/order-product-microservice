package repository

import (
	"context"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
)

type CreateOrderOption struct {
	ID       string
	Name     string
	Quantity int32
}

type Repository interface {
	CreateOrder(ctx context.Context, opt CreateOrderOption) (model.Order, error)
}
