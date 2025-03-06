package user

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

// Create handles creating a new user
func (h *Handler) Create(c *fiber.Ctx) error {
	// Parse the request body into the user struct
	var user User
	if err := c.BodyParser(&user); err != nil {
		return rest.BadRequest(err)
	}

	// Call service layer to create the user
	if err := h.service.CreateUser(&user); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "User created successfully",
	})
}

// GetByID handles retrieving a user by ID
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to get user by ID
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, user)
}

// Update handles updating a user
func (h *Handler) Update(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return rest.BadRequest(err)
	}

	// Call service layer to update the user
	if err := h.service.UpdateUser(&user); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "User updated successfully",
	})
}

// Delete handles deleting a user
func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	// Call service layer to delete the user
	if err := h.service.DeleteUser(id); err != nil {
		return rest.Error(err)
	}

	return rest.Success(c, fiber.Map{
		"message": "User deleted successfully",
	})
}

// List handles listing with pagination, sorting and search term
func (h *Handler) List(c *fiber.Ctx) error {
	// Get query parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
	searchQuery := c.Query("searchQuery", "")
	sortBy := c.Query("sortBy", "createdAt")

	// Parse is_active parameter
	isActiveParam := c.Query("isActive", "")
	var isActive *bool
	if isActiveParam != "" {
		isActiveVal := isActiveParam == "true"
		isActive = &isActiveVal
	}

	// Parse sorting fields
	sortFields, err := rest.ParseSortFields(SortColumnMap, sortBy)
	if err != nil {
		return rest.Error(err)
	}

	// Call service layer to get users
	users, total, err := h.service.GetAllUsers(page, pageSize, searchQuery, sortFields, isActive)
	if err != nil {
		return rest.Error(err)
	}

	// Calculate total pages
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	// Construct response
	response := rest.SearchResponse{
		Entities:   users,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(totalPages),
	}

	// Return success response
	return rest.Success(c, response)
}
