package userHandlers

import (
	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/users/userUsecases"
)

type IUserHandler interface {
}

type userHandler struct {
	config      *config.IConfig
	userUsecase userUsecases.IUserUsecase
}

func NewUserHandler(config *config.IConfig, userUsecase userUsecases.IUserUsecase) IUserHandler {
	return &userHandler{config: config, userUsecase: userUsecase}
}
