package middlewaresUsecases

import (
	"github.com/chon26909/e-commerce/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middleawresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(r middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{middleawresRepository: r}
}
