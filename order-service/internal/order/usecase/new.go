package usecase

import (
	"github.com/Thehien0811/order-product-microservice/internal/order/repository"
	"github.com/Thehien0811/order-product-microservice/kafka"
	"golang.org/x/net/context"
)

type UseCase interface {
	CreateOrder(ctx context.Context, input CreateOrderInput) (DetailOrder, error)
	DetailOrder(ctx context.Context, id string) (*DetailOrder, error)
}

type implUseCase struct {
	//kafka kafka.UseCase
	repo repository.Repository
}

var _ UseCase = implUseCase{}

func New(
	k kafka.UseCase,
	repo repository.Repository,
) UseCase {
	return implUseCase{
		//kafka: k,
		repo: repo,
	}
}
