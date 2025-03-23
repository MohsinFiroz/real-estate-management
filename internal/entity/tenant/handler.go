package tenant

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

// Create handles creating a new tenant
func (h *Handler) Create(c *fiber.Ctx) error {
	// Parse the request body into the tenant struct
	var tenant Tenant
	if err := c.BodyParser(&tenant); err != nil {
		return rest.BadRequest(err)
	}

	// Call service layer to create the tenant
	if err := h.service.CreateTenant(&tenant); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Tenant created successfully",
	})
}

// GetByID handles retrieving a tenant by ID
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to get tenant by ID
	tenant, err := h.service.GetTenantByID(id)
	if err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, tenant)
}

// Update handles updating a tenant
func (h *Handler) Update(c *fiber.Ctx) error {
	var tenant Tenant
	if err := c.BodyParser(&tenant); err != nil {
		return rest.BadRequest(err)
	}

	// Call service layer to update the tenant
	if err := h.service.UpdateTenant(&tenant); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Tenant updated successfully",
	})
}

// Delete handles deleting a tenant
func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to delete the tenant
	if err := h.service.DeleteTenant(id); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Tenant deleted successfully",
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

	// Call service layer to get tenants
	tenants, total, err := h.service.GetAllTenants(page, pageSize, searchQuery, sortFields)
	if err != nil {
		return rest.Error(err)
	}

	// Calculate total pages
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	// Construct response
	response := rest.SearchResponse{
		Entities:   tenants,
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
	propertyGroup := app.Group("/tenants")

	propertyGroup.Post("", h.Create)
	propertyGroup.Get("/:id", h.GetByID)
	propertyGroup.Put("/:id", h.Update)
	propertyGroup.Delete("/:id", h.Delete)
	propertyGroup.Get("", h.List)
}
