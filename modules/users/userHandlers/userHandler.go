package userHandlers

import (
	"fmt"

	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/entities"
	"github.com/chon26909/e-commerce/modules/users"
	"github.com/chon26909/e-commerce/modules/users/userUsecases"
	"github.com/gofiber/fiber/v2"
)

type userHandlerErrCode string

const (
	signUpCustomerErr userHandlerErrCode = "users-001"
)

type IUserHandler interface {
	SignUpCustomer(c *fiber.Ctx) error
}

type userHandler struct {
	config      *config.IConfig
	userUsecase userUsecases.IUserUsecase
}

func NewUserHandler(config *config.IConfig, userUsecase userUsecases.IUserUsecase) IUserHandler {
	return &userHandler{config: config, userUsecase: userUsecase}
}

func (h *userHandler) SignUpCustomer(c *fiber.Ctx) error {

	req := new(users.UserRegisterRequest)
	if err := c.BodyParser(&req); err != nil {
		return entities.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signUpCustomerErr), err.Error()).Res()
	}

	fmt.Println(*req)

	// email validattion
	if !req.IsEmail() {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signUpCustomerErr),
			"email pattern is invalid",
		).Res()
	}

	// insert
	result, err := h.userUsecase.InsertCustomer(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signUpCustomerErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, result).Res()
}
