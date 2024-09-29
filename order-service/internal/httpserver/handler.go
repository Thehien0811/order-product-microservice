package httpserver

import (
	orderHTTP "github.com/Thehien0811/order-product-microservice/internal/order/delivery/http"
	orderUC "github.com/Thehien0811/order-product-microservice/internal/order/usecase"
)

func (srv HTTPServer) mapHandlers() error {
	api := srv.gin.Group("/api/v1")
	orderUC := orderUC.New()
	orderH := orderHTTP.New(orderUC)

	orderHTTP.MapRoutes(api.Group("/orders"), orderH)
	return nil
}
