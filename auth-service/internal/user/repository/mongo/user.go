package mongo

import (
	"context"
	"log"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/user/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	users = "users"
)

func (repo implRepository) getUserCollection() mongo.Collection {
	return *repo.db.Collection(users)
}

func (repo implRepository) CreateUser(ctx context.Context, opt repository.CreateUserOption) (model.User, error) {
	col := repo.getUserCollection()

	s := repo.buildUserModel(opt)

	_, err := col.InsertOne(ctx, s)
	if err != nil {
		return model.User{}, err
	}

	return s, nil
}

func (repo implRepository) DetailUser(ctx context.Context, id string) (model.User, error) {
	col := repo.getUserCollection()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	cur, err := col.Find(ctx, bson.M{"_id": objID})
	if err != nil {
		return model.User{}, err
	}
	var users []model.User
	err = cur.All(ctx, &users)
	if err != nil {
		return model.User{}, err
	}

	return users[0], nil
}

func (repo implRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	col := repo.getUserCollection()

	cur, err := col.Find(ctx, bson.M{"username": username})
	if err != nil {
		return model.User{}, err
	}
	var users []model.User
	err = cur.All(ctx, &users)
	if err != nil {
		return model.User{}, err
	}

	return users[0], nil
}
