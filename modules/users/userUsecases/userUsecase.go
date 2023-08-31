package userUsecases

import (
	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/users/userRepositories"
)

type IUserUsecase interface {
}

type userUsecase struct {
	config         *config.IConfig
	userRepository userRepositories.IUserRepository
}

func NewUserUsecase(config *config.IConfig, userRepository userRepositories.IUserRepository) IUserUsecase {
	return &userUsecase{config: config, userRepository: userRepository}
}
