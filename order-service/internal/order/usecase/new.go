package usecase

import (
	"golang.org/x/net/context"
)

type UseCase interface {
	CreateOrder(ctx context.Context, input CreateOrderInput) (DetailOrder, error)
	DetailOrder(ctx context.Context, id string) (*DetailOrder, error)
}

type implUseCase struct {
}

var _ UseCase = implUseCase{}

func New() UseCase {
	return implUseCase{}
}
