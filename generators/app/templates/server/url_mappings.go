package server

import (
	"<%=appname%>/api"
	_ "<%=appname%>/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"
	"github.com/gofiber/swagger"
)

func MapRoutes(app *fiber.App) {
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/swagger",
		},
		StatusCode: 301,
	}))

	app.Get("/healthcheck", api.HealthCheckController.GetHealthCheck)
	app.Get("/readiness", api.ReadinessController.GetReadiness)

	app.Get("/swagger/*", swagger.HandlerDefault)
}
