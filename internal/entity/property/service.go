package property

import (
	"real-estate-management/pkg/rest"
)

// Service struct to hold the data layer for performing operations
type Service struct {
	data *Data
}

//new wntity

// NewService initializes the service layer with a data layer
func NewService(data *Data) *Service {
	return &Service{data: data}
}

// CreateProperty creates a new property after validating input
func (s *Service) CreateProperty(property *Property) error {
	return s.data.CreateProperty(property)
}

// GetPropertyByID retrieves a property by its ID
func (s *Service) GetPropertyByID(id string) (*Property, error) {
	return s.data.GetPropertyByID(id)
}

// UpdateProperty updates property information after validating input
func (s *Service) UpdateProperty(property *Property) error {
	return s.data.UpdateProperty(property)
}

// DeleteProperty deletes a property by ID
func (s *Service) DeleteProperty(id string) error {
	return s.data.DeleteProperty(id)
}

// GetAllProperties retrieves all properties with pagination, search, and sorting capabilities
func (s *Service) GetAllProperties(page, pageSize int, searchQuery string, sortFields []rest.SortField) ([]Property, int64, error) {
	// Get the properties from the data layer
	return s.data.GetAllProperties(page, pageSize, searchQuery, sortFields)
}
