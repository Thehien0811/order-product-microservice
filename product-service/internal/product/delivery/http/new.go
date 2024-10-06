package http

import (
	"github.com/Thehien0811/order-product-microservice/internal/product/usecase"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
}

type handler struct {
	uc usecase.UseCase
}

func New(uc usecase.UseCase) Handler {
	return handler{
		uc: uc,
	}
}
