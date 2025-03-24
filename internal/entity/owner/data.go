package owner

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

// CreateOwner creates a new owner in the database
func (d *Data) CreateOwner(owner *Owner) error {
	return d.db.Create(owner).Error
}

// GetOwnerByID retrieves an owner by ID
func (d *Data) GetOwnerByID(id string) (*Owner, error) {
	var owner Owner
	if err := d.db.First(&owner, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &owner, nil
}

// UpdateOwner updates owner information in the database
func (d *Data) UpdateOwner(owner *Owner) error {
	return d.db.Model(owner).Updates(owner).Error
}

// DeleteOwner deletes an owner by ID
func (d *Data) DeleteOwner(id string) error {
	return d.db.Delete(&Owner{}, "id = ?", id).Error
}

// GetAllOwners retrieves all owners with pagination, search, and sorting capabilities
func (d *Data) GetAllOwners(page, pageSize int, searchQuery string, sortFields []rest.SortField) ([]Owner, int64, error) {
	var owners []Owner
	var total int64

	query := d.db.Model(&Owner{})

	if searchQuery != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ? OR mobile ILIKE ?",
			"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	for _, sortField := range sortFields {
		query = query.Order(fmt.Sprintf("%s %s", sortField.Field, sortField.Order))
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&owners).Error; err != nil {
		return nil, 0, err
	}

	return owners, total, nil
}
