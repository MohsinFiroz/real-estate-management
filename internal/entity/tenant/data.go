package tenant

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

// CreateTenant creates a new tenant in the database
func (d *Data) CreateTenant(tenant *Tenant) error {
	return d.db.Create(tenant).Error
}

// GetTenantByID retrieves a tenant by ID
func (d *Data) GetTenantByID(id string) (*Tenant, error) {
	var tenant Tenant
	if err := d.db.First(&tenant, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

// UpdateTenant updates tenant information in the database
func (d *Data) UpdateTenant(tenant *Tenant) error {
	return d.db.Save(tenant).Error
}

// DeleteTenant deletes a tenant by ID
func (d *Data) DeleteTenant(id string) error {
	return d.db.Delete(&Tenant{}, "id = ?", id).Error
}

// GetAllTenants retrieves all tenants with pagination, search, and sorting capabilities
func (d *Data) GetAllTenants(page, pageSize int, searchQuery string, sortFields []rest.SortField) ([]Tenant, int64, error) {
	var tenants []Tenant
	var total int64

	query := d.db.Model(&Tenant{})

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

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&tenants).Error; err != nil {
		return nil, 0, err
	}

	return tenants, total, nil
}
