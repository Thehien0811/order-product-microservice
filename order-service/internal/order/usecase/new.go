package usecase

import "github.com/gin-gonic/gin"

type UseCase interface {
	CreateOrder(ctx gin.Context, input CreateOrderInput) (DetailOrder, error)
	DetailOrder(ctx gin.Context, id string) (DetailOrder, error)
}

type implUseCase struct {
}

var _ UseCase = implUseCase{}

func New() UseCase {
	return implUseCase{}
}
