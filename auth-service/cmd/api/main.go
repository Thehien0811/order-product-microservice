package main

import (
	"context"
	"log"
	"time"

	"github.com/Thehien0811/order-product-microservice/internal/httpserver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Cấu hình kết nối với MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Kết nối đến MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Kiểm tra kết nối
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("users")

	srv := httpserver.New(db)
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
