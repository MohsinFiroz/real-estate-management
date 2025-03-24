package property

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

// CreateProperty creates a new property in the database
func (d *Data) CreateProperty(property *Property) error {
	return d.db.Create(property).Error
}

// GetPropertyByID retrieves a property by ID
func (d *Data) GetPropertyByID(id string) (*Property, error) {
	var property Property
	if err := d.db.First(&property, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &property, nil
}

// UpdateProperty updates property information in the database
func (d *Data) UpdateProperty(property *Property) error {
	return d.db.Model(property).Updates(property).Error
}

// DeleteProperty deletes a property by ID
func (d *Data) DeleteProperty(id string) error {
	return d.db.Delete(&Property{}, "id = ?", id).Error
}

// GetAllProperties retrieves all properties with pagination, search, and sorting capabilities
func (d *Data) GetAllProperties(page, pageSize int, searchQuery string, sortFields []rest.SortField, ownerID *string) ([]Property, int64, error) {
	var properties []Property
	var total int64

	query := d.db.Model(&Property{})

	if searchQuery != "" {
		query = query.Where("address ILIKE ? OR suburb ILIKE ? OR postcode ILIKE ?",
			"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	if ownerID != nil {
		query = query.Where("owner_id = ?", *ownerID)
	}

	for _, sortField := range sortFields {
		query = query.Order(fmt.Sprintf("%s %s", sortField.Field, sortField.Order))
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&properties).Error; err != nil {
		return nil, 0, err
	}

	return properties, total, nil
}
