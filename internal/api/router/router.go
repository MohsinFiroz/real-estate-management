package router

import (
	"real-estate-management/internal/entity/property"
	"real-estate-management/internal/entity/user"
	"real-estate-management/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *fiber.App {
	app := fiber.New()

	// Initialize repositories
	propertyRepo := property.NewData(db)
	userRepo := user.NewData(db)

	// Initialize services
	propertyService := property.NewService(propertyRepo)
	userService := user.NewService(userRepo)

	// Initialize handlers
	propertyHandler := property.NewHandler(propertyService)
	userHandler := user.NewHandler(userService)

	// API routes
	api := app.Group("/v1")

	// User routes
	users := api.Group("/users")
	users.Post("/", userHandler.Create)
	users.Get("/", middleware.Auth(), userHandler.GetAll)
	users.Get("/:id", middleware.Auth(), userHandler.GetByID)

	// Property routes
	properties := api.Group("/properties")
	properties.Post("/", middleware.Auth(), propertyHandler.Create)
	properties.Get("/", propertyHandler.GetAll)
	properties.Get("/:id", propertyHandler.GetByID)

	return app
}
