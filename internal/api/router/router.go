package router

import (
	"errors"
	"real-estate-management/internal/entity/property"
	"real-estate-management/internal/entity/user"
	"real-estate-management/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *fiber.App {
	// Configure Fiber with custom error handling
	app := fiber.New(fiber.Config{
		// Enable custom error handling
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Handle errors thrown during routing
			code := fiber.StatusInternalServerError

			// Check if it's a Fiber error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Return JSON error response with details
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		},
		// Display detailed error messages
		EnablePrintRoutes: true,
	})

	app.Use(middleware.Recover())
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

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

	// Monitor route
	api.Get("/metrics", middleware.Monitor())

	// User routes
	users := api.Group("/users")
	users.Post("/", userHandler.Create)
	users.Get("/", userHandler.List)
	users.Get("/:id", userHandler.GetByID)
	users.Put("/:id", userHandler.Update)
	users.Delete("/:id", middleware.Auth(), userHandler.Delete)

	// Property routes
	properties := api.Group("/properties")
	properties.Post("/", middleware.Auth(), propertyHandler.Create)
	properties.Get("/", propertyHandler.GetAll)
	properties.Get("/:id", propertyHandler.GetByID)

	return app
}
