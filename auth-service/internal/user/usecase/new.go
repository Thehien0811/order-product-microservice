package usecase

import (
	"github.com/Thehien0811/order-product-microservice/internal/user/repository"
	"github.com/Thehien0811/order-product-microservice/kafka"
	"golang.org/x/net/context"
)

type UseCase interface {
	SignIn(ctx context.Context, input SignInInput) (DetailUser, error)
	SignUp(ctx context.Context, input SignUpInput) (string, error)
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
