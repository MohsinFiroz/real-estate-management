package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS() fiber.Handler {
	config := cors.Config{
		//AllowOrigins: "https://rem.softcelia.com",
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}
	return cors.New(config)
}
