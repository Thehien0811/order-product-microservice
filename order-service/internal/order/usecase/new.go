package usecase

import (
	"golang.org/x/net/context"
	"github.com/Thehien0811/order-product-microservice/kafka"
)

type UseCase interface {
	CreateOrder(ctx context.Context, input CreateOrderInput) (DetailOrder, error)
	DetailOrder(ctx context.Context, id string) (*DetailOrder, error)
}

type implUseCase struct {
	kafka kafka.UseCase
}

var _ UseCase = implUseCase{}

func New(
	k kafka.UseCase,
) UseCase {
	return implUseCase{
		kafka: k,
	}
}
