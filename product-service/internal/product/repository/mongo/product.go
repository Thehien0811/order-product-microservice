package mongo

import (
	"context"
	"fmt"
	"log"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/product/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	order = "products"
)

func (repo implRepository) getProductCollection() mongo.Collection {
	return *repo.db.Collection(order)
}

func (repo implRepository) CreateProduct(ctx context.Context, opt repository.CreateProductOption) (model.Product, error) {
	col := repo.getProductCollection()

	s := repo.buildProductModel(opt)

	_, err := col.InsertOne(ctx, s)
	if err != nil {
		return model.Product{}, err
	}

	return s, nil
}

func (repo implRepository) DetailProduct(ctx context.Context, id string) (model.Product, error) {
	col := repo.getProductCollection()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	cur, err := col.Find(ctx, bson.M{"_id": objID})
	if err != nil {
		return model.Product{}, err
	}
	var orders []model.Product
	err = cur.All(ctx, &orders)
	if err != nil {
		return model.Product{}, err
	}

	return orders[0], nil
}

func (repo implRepository) UpdateProduct(ctx context.Context, product model.Product) (bool, error) {
	col := repo.getProductCollection()

	// Xác định tiêu chí và cập nhật
	filter := bson.M{"_id": product.ID}
	update := bson.M{"$set": bson.M{"quantity": product.Quantity}}
	updateResult, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, fmt.Errorf("could not update product quantity: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		return false, fmt.Errorf("no product found with ID: %v", product.ID.Hex())
	}

	return true, nil
}
