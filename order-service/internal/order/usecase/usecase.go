package usecase

import (
	"errors"

	"golang.org/x/net/context"
)

var orders = make(map[string]DetailOrder)

func (uc implUseCase) CreateOrder(ctx context.Context, input CreateOrderInput) (DetailOrder, error) {
	order := DetailOrder{
		input.ID,
		input.Name,
		input.Quantity,
	}
	orders[input.ID] = order
	return order, nil
}

func (uc implUseCase) DetailOrder(ctx context.Context, id string) (*DetailOrder, error) {

	order, exists := orders[id]
	if !exists {
		return nil, errors.New("Not found order")
	}
	return &order, nil
}
