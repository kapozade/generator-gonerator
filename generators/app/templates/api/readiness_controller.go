package api

import (
	"<%=appname%>/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type readinesController struct{}

var (
	ReadinessController readinesController
)

// Readiness godoc
// @Summary Gets API Readiness
// @Tags Readiness
// @Accept json
// @Produce json
// @Success 200
// @Failure 500
// @Router /readiness [get]
func (r *readinesController) GetReadiness(c *fiber.Ctx) error {
	currentTime := time.Now()
	return c.Status(fiber.StatusOK).JSON(&models.Readiness{
		Now: currentTime.Format("2006-01-02 15:04:05.000"),
	})
}
