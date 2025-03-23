package tenant

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

// CreateTenant creates a new tenant after validating input
func (s *Service) CreateTenant(tenant *Tenant) error {
	return s.data.CreateTenant(tenant)
}

// GetTenantByID retrieves a tenant by their ID
func (s *Service) GetTenantByID(id string) (*Tenant, error) {
	return s.data.GetTenantByID(id)
}

// UpdateTenant updates tenant information after validating input
func (s *Service) UpdateTenant(tenant *Tenant) error {
	return s.data.UpdateTenant(tenant)
}

// DeleteTenant deletes a tenant by ID
func (s *Service) DeleteTenant(id string) error {
	return s.data.DeleteTenant(id)
}

// GetAllTenants retrieves all tenants with pagination, search, and sorting capabilities
func (s *Service) GetAllTenants(page, pageSize int, searchQuery string, sortFields []rest.SortField) ([]Tenant, int64, error) {
	// Get the tenants from the data layer
	return s.data.GetAllTenants(page, pageSize, searchQuery, sortFields)
}
