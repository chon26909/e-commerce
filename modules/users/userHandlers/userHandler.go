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
	signInErr         userHandlerErrCode = "users-002"
	refreshTokenErr   userHandlerErrCode = "users-003"
	signOutErr        userHandlerErrCode = "users-004"
)

type IUserHandler interface {
	SignUpCustomer(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	RefreshPassport(c *fiber.Ctx) error
}

type userHandler struct {
	config      config.IConfig
	userUsecase userUsecases.IUserUsecase
}

func NewUserHandler(config config.IConfig, userUsecase userUsecases.IUserUsecase) IUserHandler {
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

func (h *userHandler) SignIn(c *fiber.Ctx) error {

	req := new(users.UserCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signInErr), err.Error()).Res()
	}

	passport, err := h.userUsecase.GetPassport(req)
	if err != nil {
		return entities.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signInErr), err.Error()).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (h *userHandler) RefreshPassport(c *fiber.Ctx) error {

	req := new(users.UserRefreshCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(refreshTokenErr), err.Error()).Res()
	}

	passport, err := h.userUsecase.RefreshPassport(req)
	if err != nil {
		return entities.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(refreshTokenErr), err.Error()).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (h *userHandler) SignOut(c *fiber.Ctx) error {
	req := new(users.UserRemoveCreadential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signOutErr),
			err.Error(),
		).Res()
	}

	if err := h.userUsecase.DeleteOauth(req.OauthId); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signOutErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, nil).Res()
}
