package http

import (
	"github.com/Thehien0811/order-product-microservice/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type handler struct {
	uc usecase.UseCase
}

func New(uc usecase.UseCase) Handler {
	return handler{
		uc: uc,
	}
}
