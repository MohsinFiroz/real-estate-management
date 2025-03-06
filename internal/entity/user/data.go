package user

import (
	"fmt"
	"gorm.io/gorm"
	"real-estate-management/pkg/rest"
)

// Data struct to hold the GORM DB instance for performing database operations
type Data struct {
	db *gorm.DB
}

// NewData initializes the Data layer with a GORM DB instance
func NewData(db *gorm.DB) *Data {
	return &Data{db: db}
}

// CreateUser creates a new user in the database
func (d *Data) CreateUser(user *User) error {
	return d.db.Create(user).Error
}

// GetUserByID retrieves a user by ID
func (d *Data) GetUserByID(id string) (*User, error) {
	var user User
	if err := d.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates user information in the database
func (d *Data) UpdateUser(user *User) error {
	return d.db.Save(user).Error
}

// DeleteUser deletes a user by ID
func (d *Data) DeleteUser(id string) error {
	return d.db.Delete(&User{}, "id = ?", id).Error
}

func (d *Data) GetAllUsers(page, pageSize int, searchQuery string, sortFields []rest.SortField, isActive *bool) ([]User, int64, error) {
	var users []User
	var total int64

	query := d.db.Model(&User{})

	if searchQuery != "" {
		query = query.Where("email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ?",
			"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	if isActive != nil {
		query = query.Where("is_active = ?", *isActive)
	}

	for _, sortField := range sortFields {
		query = query.Order(fmt.Sprintf("%s %s", sortField.Field, sortField.Order))
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
