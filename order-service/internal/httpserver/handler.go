package httpserver

import (
	orderHTTP "github.com/Thehien0811/order-product-microservice/internal/order/delivery/http"
	orderUC "github.com/Thehien0811/order-product-microservice/internal/order/usecase"
	myKafka "github.com/Thehien0811/order-product-microservice/kafka"
	orderRepo "github.com/Thehien0811/order-product-microservice/internal/order/repository/mongo"
)

func (srv HTTPServer) mapHandlers() error {
	api := srv.gin.Group("/api/v1")
	k := myKafka.New()
	orderRepo := orderRepo.New(*srv.db)
	orderUC := orderUC.New(k, orderRepo)
	orderH := orderHTTP.New(orderUC)
	orderHTTP.MapRoutes(api.Group("/orders"), orderH)
	return nil
}
