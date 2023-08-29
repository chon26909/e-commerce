package monitorhandlers

import (
	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/entities"
	"github.com/chon26909/e-commerce/modules/monitor"
	"github.com/gofiber/fiber/v2"
)

type IMonitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	config config.IConfig
}

func NewMonitorHandler(config config.IConfig) IMonitorHandler {
	return &monitorHandler{
		config: config,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {

	res := &monitor.Monitor{
		Name:    h.config.App().Name(),
		Version: h.config.App().Version(),
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
