package usecase

import (
	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/product/repository"
	"github.com/Thehien0811/order-product-microservice/kafka"
	"golang.org/x/net/context"
)

type UseCase interface {
	CreateProduct(ctx context.Context, input CreateProductInput) (DetailProduct, error)
	DetailProduct(ctx context.Context, id string) (DetailProduct, error)
	UpdateProduct(ctx context.Context, product model.Product) (bool, error)
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
