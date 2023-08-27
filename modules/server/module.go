package server

import (
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

func (m *moduleFactory) MonitorModule() {
	handler := monitorhandlers.NewMonitorHandler(m.server.config)

	m.router.Get("/", handler.HealthCheck)
}
