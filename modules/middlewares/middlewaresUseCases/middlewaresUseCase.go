package middlewaresUsecases

import (
	"github.com/chon26909/e-commerce/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface {
	FindAccessToken(userId, accessToken string) bool
}

type middlewaresUsecase struct {
	middleawresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(r middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{middleawresRepository: r}
}

func (u *middlewaresUsecase) FindAccessToken(userId, accessToken string) bool {
	return u.middleawresRepository.FindAccessToken(userId, accessToken)
}
