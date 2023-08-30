package server

import (
	"github.com/chon26909/e-commerce/modules/middlewares/middlewaresHandlers"
	"github.com/chon26909/e-commerce/modules/middlewares/middlewaresRepositories"
	middlewaresUsecases "github.com/chon26909/e-commerce/modules/middlewares/middlewaresUseCases"
	monitorhandlers "github.com/chon26909/e-commerce/modules/monitor/MonitorHandlers"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
}

func NewModule(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
	}
}

func NewMiddleware(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(&s.config, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorhandlers.NewMonitorHandler(m.server.config)

	m.router.Get("/", handler.HealthCheck)
}
