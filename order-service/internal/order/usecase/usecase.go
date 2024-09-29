package usecase

import "github.com/gin-gonic/gin"

func (uc implUseCase) CreateOrder(ctx gin.Context, input CreateOrderInput) (DetailOrder, error) {
	return DetailOrder{
		input.ID,
		input.Name,
		input.Quantity,
	}, nil
}

func (uc implUseCase) DetailOrder(ctx gin.Context, id string) (DetailOrder, error) {
	return DetailOrder{
		ID: id,
		Name: "A",
		Quantity: 1,
	}, nil
}