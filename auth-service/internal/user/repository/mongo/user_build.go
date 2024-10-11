package mongo

import (
	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/user/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo implRepository) buildUserModel(opt repository.CreateUserOption) model.User {
	s := model.User{
		ID:       primitive.NewObjectID(),
		Username: opt.Username,
		Password: opt.Password,
	}
	return s
}
