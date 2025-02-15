package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Monitor() fiber.Handler {
	config := monitor.Config{
		Title:      "Real Estate Management Server Metrics ",
		CustomHead: "Implemented by Mohsin",
	}
	return monitor.New(config)
}
