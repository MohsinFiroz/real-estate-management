package owner

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

// Create handles creating a new owner
func (h *Handler) Create(c *fiber.Ctx) error {
	// Parse the request body into the owner struct
	var owner Owner
	if err := c.BodyParser(&owner); err != nil {
		return rest.BadRequest(err)
	}

	// Call service layer to create the owner
	if err := h.service.CreateOwner(&owner); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Owner created successfully",
	})
}

// GetByID handles retrieving an owner by ID
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to get owner by ID
	owner, err := h.service.GetOwnerByID(id)
	if err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, owner)
}

// Update handles updating an owner
func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var owner Owner
	if err := c.BodyParser(&owner); err != nil {
		return rest.BadRequest(err)
	}
	owner.ID = id

	// Call service layer to update the owner
	if err := h.service.UpdateOwner(&owner); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Owner updated successfully",
	})
}

// Delete handles deleting an owner
func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to delete the owner
	if err := h.service.DeleteOwner(id); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "Owner deleted successfully",
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

	// Call service layer to get owners
	owners, total, err := h.service.GetAllOwners(page, pageSize, searchQuery, sortFields)
	if err != nil {
		return rest.Error(err)
	}

	// Calculate total pages
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	// Construct response
	response := rest.SearchResponse{
		Entities:   owners,
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
	propertyGroup := router.Group("/owners")

	propertyGroup.Post("", h.Create)
	propertyGroup.Get("/:id", h.GetByID)
	propertyGroup.Put("/:id", h.Update)
	propertyGroup.Delete("/:id", h.Delete)
	propertyGroup.Get("", h.List)
}
