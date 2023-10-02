package server

import (
	"github.com/chon26909/e-commerce/modules/middlewares/middlewaresHandlers"
	"github.com/chon26909/e-commerce/modules/middlewares/middlewaresRepositories"
	middlewaresUsecases "github.com/chon26909/e-commerce/modules/middlewares/middlewaresUseCases"
	monitorhandlers "github.com/chon26909/e-commerce/modules/monitor/MonitorHandlers"
	"github.com/chon26909/e-commerce/modules/users/userHandlers"
	"github.com/chon26909/e-commerce/modules/users/userRepositories"
	"github.com/chon26909/e-commerce/modules/users/userUsecases"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
}

type moduleFactory struct {
	router     fiber.Router
	server     *server
	middleware middlewaresHandlers.IMiddlewaresHandler
}

func NewModule(r fiber.Router, s *server, m middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		router:     r,
		server:     s,
		middleware: m,
	}
}

func NewMiddleware(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.config, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorhandlers.NewMonitorHandler(m.server.config)

	m.router.Get("/", handler.HealthCheck)
}

func (m *moduleFactory) UsersModule() {
	repository := userRepositories.NewUserRepositories(m.server.db)
	usecase := userUsecases.NewUserUsecase(m.server.config, repository)
	handler := userHandlers.NewUserHandler(m.server.config, usecase)

	router := m.router.Group("/users")

	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SignIn)
	router.Post("/refresh", handler.RefreshPassport)
	router.Post("/signout", handler.SignOut)

	router.Get("/admin-token", m.middleware.JwtAuth(), handler.GenerateAdminToken)
}
