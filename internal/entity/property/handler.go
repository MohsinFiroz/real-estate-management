package property

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"real-estate-management/pkg/rest"
)

// Handler struct to handle HTTP requests for property operations
type Handler struct {
	service *Service
}

// NewHandler initializes the handler with a service layer
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Create handles the request to create a new property
func (h *Handler) Create(c *fiber.Ctx) error {
	var property Property
	if err := c.BodyParser(&property); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateProperty(&property); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(property)
}

// GetByID handles the request to get a property by ID
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	property, err := h.service.GetPropertyByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	return c.Status(fiber.StatusOK).JSON(property)
}

// Update handles the request to update a property
func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var property Property
	if err := c.BodyParser(&property); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	property.ID = id
	if err := h.service.UpdateProperty(&property); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(property)
}

// Delete handles the request to delete a property
func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.DeleteProperty(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Property deleted successfully"})
}

// List handles the request to get all properties with pagination, search, and sorting
func (h *Handler) List(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
	searchQuery := c.Query("searchQuery", "")
	sortBy := c.Query("sortBy", "createdAt")

	// Parse sorting fields
	sortFields, err := rest.ParseSortFields(SortColumnMap, sortBy)
	if err != nil {
		return rest.Error(err)
	}

	// Parse owner filter if provided
	var ownerID *string
	if owner := c.Query("ownerID"); owner != "" {
		ownerID = &owner
	}

	properties, total, err := h.service.GetAllProperties(page, pageSize, searchQuery, sortFields, ownerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Calculate total pages
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	// Construct response
	response := rest.SearchResponse{
		Entities:   properties,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(totalPages),
	}

	// Return success response
	return rest.Success(c, response)
}

// RegisterRoutes registers all the routes for property operations
func (h *Handler) RegisterRoutes(app *fiber.App) {
	propertyGroup := app.Group("/properties")

	propertyGroup.Post("", h.Create)
	propertyGroup.Get("/:id", h.GetByID)
	propertyGroup.Put("/:id", h.Update)
	propertyGroup.Delete("/:id", h.Delete)
	propertyGroup.Get("", h.List)
}
