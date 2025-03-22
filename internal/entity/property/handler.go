package property

import (
	"github.com/gofiber/fiber/v2"
	"real-estate-management/pkg/rest"
	"strconv"
)

// Handler struct
type Handler struct {
	service *Service
}

// NewHandler initializes a new Handler with the service layer
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Create handles creating a new property
func (h *Handler) Create(c *fiber.Ctx) error {
	// Parse the request body into the property struct
	var property Property
	if err := c.BodyParser(&property); err != nil {
		return rest.BadRequest(err)
	}

	// Call service layer to create the property
	if err := h.service.CreateProperty(&property); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Property created successfully",
	})
}

// GetByID handles retrieving a property by ID
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to get property by ID
	property, err := h.service.GetPropertyByID(id)
	if err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, property)
}

// Update handles updating a property
func (h *Handler) Update(c *fiber.Ctx) error {
	var property Property
	if err := c.BodyParser(&property); err != nil {
		return rest.BadRequest(err)
	}

	// Call service layer to update the property
	if err := h.service.UpdateProperty(&property); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Property updated successfully",
	})
}

// Delete handles deleting a property
func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to delete the property
	if err := h.service.DeleteProperty(id); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Property deleted successfully",
	})
}

// List handles listing properties with pagination, sorting, and search term
func (h *Handler) List(c *fiber.Ctx) error {
	// Get query parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
	searchQuery := c.Query("searchQuery", "")
	sortBy := c.Query("sortBy", "createdAt")

	// Parse sorting fields
	sortFields, err := rest.ParseSortFields(SortColumnMap, sortBy)
	if err != nil {
		return rest.Error(err)
	}

	// Call service layer to get properties
	properties, total, err := h.service.GetAllProperties(page, pageSize, searchQuery, sortFields)
	if err != nil {
		return rest.Error(err)
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
