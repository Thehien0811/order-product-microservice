package mongo

import (
	"github.com/Thehien0811/order-product-microservice/internal/user/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type implRepository struct {
	db mongo.Database
}

var _ repository.Repository = implRepository{}

func New(
	db mongo.Database,
) repository.Repository {
	return implRepository{
		db: db,
	}
}
