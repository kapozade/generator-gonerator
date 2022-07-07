package api

import (
	"<%=appname%>/services"

	"github.com/gofiber/fiber/v2"
)

type healthCheckController struct{}

var (
	HealthCheckController healthCheckController
)

// HealthCheck godoc
// @Summary Checks healtcheck of API
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200
// @Failure 500
// @Router  /healthcheck [get]
func (h *healthCheckController) GetHealthCheck(c *fiber.Ctx) error {
	config, err := services.ConfigurationService.LoadHealthCheckConfig()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(config)
}
