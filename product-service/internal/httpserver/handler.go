package httpserver

import (
	productHTTP "github.com/Thehien0811/order-product-microservice/internal/product/delivery/http"
	productUC "github.com/Thehien0811/order-product-microservice/internal/product/usecase"
	myKafka "github.com/Thehien0811/order-product-microservice/kafka"
	productRepo "github.com/Thehien0811/order-product-microservice/internal/product/repository/mongo"
)

func (srv HTTPServer) mapHandlers() error {
	api := srv.gin.Group("/api/v1")
	k := myKafka.New()
	productRepo := productRepo.New(*srv.db)
	productUC := productUC.New(k, productRepo)
	orderH := productHTTP.New(productUC)
	productHTTP.MapRoutes(api.Group("/products"), orderH)
	return nil
}
