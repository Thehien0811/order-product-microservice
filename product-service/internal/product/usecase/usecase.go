package usecase

import (
	"log"

	"github.com/Thehien0811/order-product-microservice/internal/product/repository"
	"golang.org/x/net/context"
)

func (uc implUseCase) CreateProduct(ctx context.Context, input CreateProductInput) (DetailProduct, error) {
	o, err := uc.repo.CreateProduct(ctx, repository.CreateProductOption{
		Name:     input.Name,
		Quantity: input.Quantity,
	})
	if err != nil {
		log.Fatal(err)
	}

	return DetailProduct{
		o.ID.Hex(),
		o.Name,
		o.Quantity,
	}, nil

}

func (uc implUseCase) DetailProduct(ctx context.Context, id string) (DetailProduct, error) {
	o, err := uc.repo.DetailProduct(ctx, id)
	if err != nil {
		log.Fatal(err)
		return DetailProduct{}, nil
	}
	return DetailProduct{
		o.ID.Hex(),
		o.Name,
		o.Quantity,
	}, nil
}
