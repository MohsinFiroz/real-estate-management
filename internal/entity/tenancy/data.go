package tenancy

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

// CreateTenancy creates a new tenancy in the database
func (d *Data) CreateTenancy(tenancy *Tenancy) error {
	return d.db.Create(tenancy).Error
}

// GetTenancyByID retrieves a tenancy by ID
func (d *Data) GetTenancyByID(id string) (*Tenancy, error) {
	var tenancy Tenancy
	if err := d.db.First(&tenancy, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tenancy, nil
}

// UpdateTenancy updates tenancy information in the database
func (d *Data) UpdateTenancy(tenancy *Tenancy) error {
	return d.db.Model(tenancy).Updates(tenancy).Error
}

// DeleteTenancy deletes a tenancy by ID
func (d *Data) DeleteTenancy(id string) error {
	return d.db.Delete(&Tenancy{}, "id = ?", id).Error
}

// GetAllTenancies retrieves all tenancies with pagination, search, and sorting capabilities
func (d *Data) GetAllTenancies(page, pageSize int, searchQuery string, sortFields []rest.SortField) ([]Tenancy, int64, error) {
	var tenancies []Tenancy
	var total int64

	query := d.db.Model(&Tenancy{})

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

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&tenancies).Error; err != nil {
		return nil, 0, err
	}

	return tenancies, total, nil
}
