package repository

import (
	"context"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
)

type CreateUserOption struct {
	Username string
	Password string
}

type Repository interface {
	CreateUser(ctx context.Context, opt CreateUserOption) (model.User, error)
	DetailUser(ctx context.Context, id string) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}
