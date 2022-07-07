package server

import (
	"<%=appname%>/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartUp() {
	app := fiber.New()
	MapRoutes(app)
	config, err := services.ConfigurationService.GetApplicationConfiguration()

	var port int64
	if err != nil {
		port = 5000
	} else {
		port = config.Port
	}

	app.Listen(fmt.Sprintf(":%d", port))
}
