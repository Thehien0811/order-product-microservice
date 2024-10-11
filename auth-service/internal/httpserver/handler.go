package httpserver

import (
	userHTTP "github.com/Thehien0811/order-product-microservice/internal/user/delivery/http"
	userUC "github.com/Thehien0811/order-product-microservice/internal/user/usecase"
	myKafka "github.com/Thehien0811/order-product-microservice/kafka"
	userRepo "github.com/Thehien0811/order-product-microservice/internal/user/repository/mongo"
)

func (srv HTTPServer) mapHandlers() error {
	api := srv.gin.Group("/api/v1")
	k := myKafka.New()
	userRepo := userRepo.New(*srv.db)
	userUC := userUC.New(k, userRepo)
	orderH := userHTTP.New(userUC)
	userHTTP.MapRoutes(api.Group("/auth"), orderH)
	return nil
}
