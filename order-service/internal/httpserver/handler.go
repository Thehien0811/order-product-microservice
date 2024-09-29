package httpserver

import (
	orderHTTP "github.com/Thehien0811/order-product-microservice/internal/order/delivery/http"
	orderUC "github.com/Thehien0811/order-product-microservice/internal/order/usecase"
	myKafka "github.com/Thehien0811/order-product-microservice/kafka"
)

func (srv HTTPServer) mapHandlers() error {
	api := srv.gin.Group("/api/v1")
	k := myKafka.New()
	orderUC := orderUC.New(k)
	orderH := orderHTTP.New(orderUC)

	orderHTTP.MapRoutes(api.Group("/orders"), orderH)
	return nil
}
