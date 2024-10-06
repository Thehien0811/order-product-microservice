package usecase

import (
	"log"

	"github.com/Thehien0811/order-product-microservice/internal/order/repository"
	"github.com/Thehien0811/order-product-microservice/pkg/curl"
	"golang.org/x/net/context"
)

func (uc implUseCase) CreateOrder(ctx context.Context, input CreateOrderInput) (DetailOrder, error) {
	_, err := curl.DetailProduct(input.ProductID)
	if err != nil {
		log.Fatal(err)
	}

	o, err := uc.repo.CreateOrder(ctx, repository.CreateOrderOption{
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	})
	if err != nil {
		log.Fatal(err)
	}

	return DetailOrder{
		o.ID.Hex(),
		o.ProductID,
		o.Quantity,
	}, nil

}

func (uc implUseCase) DetailOrder(ctx context.Context, id string) (DetailOrder, error) {
	o, err := uc.repo.DetailOrder(ctx, id)
	if err != nil {
		log.Fatal(err)
		return DetailOrder{}, nil
	}
	return DetailOrder{
		o.ID.Hex(),
		o.ProductID,
		o.Quantity,
	}, nil
}
