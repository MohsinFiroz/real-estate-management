package tenancy

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

// CreateTenancy creates a new tenancy after validating input
func (s *Service) CreateTenancy(tenancy CreateTenancyRequest) error {
	return s.data.CreateTenancy(tenancy)
}

// GetTenancyByID retrieves a tenancy by their ID
func (s *Service) GetTenancyByID(id string) (*Tenancy, error) {
	return s.data.GetTenancyByID(id)
}

// UpdateTenancy updates tenancy information after validating input
func (s *Service) UpdateTenancy(tenancy *Tenancy) error {
	return s.data.UpdateTenancy(tenancy)
}

// DeleteTenancy deletes a tenancy by ID
func (s *Service) DeleteTenancy(id string) error {
	return s.data.DeleteTenancy(id)
}

// GetAllTenancies retrieves all tenancies with pagination, search, and sorting capabilities
func (s *Service) GetAllTenancies(page, pageSize int, searchQuery string, sortFields []rest.SortField) ([]Tenancy, int64, error) {
	// Get the tenancies from the data layer
	return s.data.GetAllTenancies(page, pageSize, searchQuery, sortFields)
}
