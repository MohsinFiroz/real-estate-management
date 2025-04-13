package tenancy

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

// Create handles creating a new tenancy
func (h *Handler) Create(c *fiber.Ctx) error {
	// Parse the request body into the tenancy struct
	var tenancy Tenancy
	if err := c.BodyParser(&tenancy); err != nil {
		return rest.BadRequest(err)
	}

	// Call service layer to create the tenancy
	if err := h.service.CreateTenancy(&tenancy); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Tenancy created successfully",
	})
}

// GetByID handles retrieving a tenancy by ID
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to get tenancy by ID
	tenancy, err := h.service.GetTenancyByID(id)
	if err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, tenancy)
}

// Update handles updating a tenancy
func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var tenancy Tenancy
	if err := c.BodyParser(&tenancy); err != nil {
		return rest.BadRequest(err)
	}
	tenancy.ID = id

	// Call service layer to update the tenancy
	if err := h.service.UpdateTenancy(&tenancy); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Tenancy updated successfully",
	})
}

// Delete handles deleting a tenancy
func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to delete the tenancy
	if err := h.service.DeleteTenancy(id); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Tenancy deleted successfully",
	})
}

// List handles listing with pagination, sorting and search term
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

	// Call service layer to get tenancies
	tenancies, total, err := h.service.GetAllTenancies(page, pageSize, searchQuery, sortFields)
	if err != nil {
		return rest.Error(err)
	}

	// Calculate total pages
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	// Construct response
	response := rest.SearchResponse{
		Entities:   tenancies,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(totalPages),
	}

	// Return success response
	return rest.Success(c, response)
}

// RegisterRoutes registers all the routes for property operations
func (h *Handler) RegisterRoutes(router fiber.Router) {
	propertyGroup := router.Group("/tenancies")

	propertyGroup.Post("", h.Create)
	propertyGroup.Get("/:id", h.GetByID)
	propertyGroup.Put("/:id", h.Update)
	propertyGroup.Delete("/:id", h.Delete)
	propertyGroup.Get("", h.List)
}
