package middlewaresHandlers

import (
	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/entities"
	middlewaresUsecases "github.com/chon26909/e-commerce/modules/middlewares/middlewaresUseCases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type middlewareHandlersErrCode string

const (
	routerCheckError middlewareHandlersErrCode = "M-0001"
)

type IMiddlewaresHandler interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
}

type middlewaresHandler struct {
	config            *config.IConfig
	middlewareUsecase middlewaresUsecases.IMiddlewaresUsecase
}

func MiddlewaresHandler(c *config.IConfig, u middlewaresUsecases.IMiddlewaresUsecase) IMiddlewaresHandler {
	return &middlewaresHandler{config: c, middlewareUsecase: u}
}

func (h *middlewaresHandler) Cors() fiber.Handler {

	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE,PATCH",
		AllowHeaders:     "*",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})

}

func (h *middlewaresHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return entities.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(routerCheckError),
			"router not found",
		).Res()
	}
}
