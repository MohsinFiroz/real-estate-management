package owner

import (
	"real-estate-management/pkg/rest"
)

// Service struct to hold the data layer for performing operations
type Service struct {
	data *Data
}

// NewService initializes the service layer with a data layer
func NewService(data *Data) *Service {
	return &Service{data: data}
}

// CreateOwner creates a new owner after validating input
func (s *Service) CreateOwner(owner *Owner) error {
	return s.data.CreateOwner(owner)
}

// GetOwnerByID retrieves an owner by their ID
func (s *Service) GetOwnerByID(id string) (*Owner, error) {
	return s.data.GetOwnerByID(id)
}

// UpdateOwner updates owner information after validating input
func (s *Service) UpdateOwner(owner *Owner) error {
	return s.data.UpdateOwner(owner)
}

// DeleteOwner deletes an owner by ID
func (s *Service) DeleteOwner(id string) error {
	return s.data.DeleteOwner(id)
}

// GetAllOwners retrieves all owners with pagination, search, and sorting capabilities
func (s *Service) GetAllOwners(page, pageSize int, searchQuery string, sortFields []rest.SortField) ([]Owner, int64, error) {
	// Get the owners from the data layer
	return s.data.GetAllOwners(page, pageSize, searchQuery, sortFields)
}
