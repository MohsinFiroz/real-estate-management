package user

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

// CreateUser creates a new user after validating input
func (s *Service) CreateUser(user *User) error {
	return s.data.CreateUser(user)
}

// GetUserByID retrieves a user by their ID
func (s *Service) GetUserByID(id string) (*User, error) {
	return s.data.GetUserByID(id)
}

// UpdateUser updates user information after validating input
func (s *Service) UpdateUser(user *User) error {
	return s.data.UpdateUser(user)
}

// DeleteUser deletes a user by ID
func (s *Service) DeleteUser(id string) error {
	return s.data.DeleteUser(id)
}

// GetAllUsers retrieves all users with pagination, search, and sorting capabilities
func (s *Service) GetAllUsers(page, pageSize int, searchQuery string, sortFields []rest.SortField, isActive *bool) ([]User, int64, error) {
	// Get the users from the data layer
	return s.data.GetAllUsers(page, pageSize, searchQuery, sortFields, isActive)
}
