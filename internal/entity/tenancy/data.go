package tenancy

import (
	"fmt"
	"real-estate-management/internal/entity/tenant"
	"real-estate-management/pkg/rest"

	"gorm.io/gorm"
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
func (d *Data) CreateTenancy(req CreateTenancyRequest) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		// Create tenancy first
		if err := tx.Create(req.Tenancy).Error; err != nil {
			return err
		}

		// Insert into tenancy_tenants table
		var tenancyTenants []TenancyTenant
		for _, tenantID := range req.TenantIDs {
			tenancyTenants = append(tenancyTenants, TenancyTenant{
				TenancyID: req.Tenancy.ID,
				TenantID:  tenantID,
			})
		}
		if len(tenancyTenants) > 0 {
			if err := tx.Create(&tenancyTenants).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// GetTenancyByID retrieves a tenancy by ID
func (d *Data) GetTenancyByID(id string) (*Tenancy, error) {
	var tenancy Tenancy
	if err := d.db.First(&tenancy, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Fetch primary tenant
	var primaryTenant tenant.Tenant
	if err := d.db.First(&primaryTenant, "id = ?", tenancy.PrimaryTenantID).Error; err == nil {
		tenancy.PrimaryTenant = &primaryTenant
	}

	// Fetch all tenants linked to this tenancy
	var tenants []tenant.Tenant
	if err := d.db.
		Table("tenants").
		Select("tenants.*").
		Joins("INNER JOIN tenancy_tenants ON tenants.id = tenancy_tenants.tenant_id").
		Where("tenancy_tenants.tenancy_id = ?", tenancy.ID).
		Find(&tenants).Error; err != nil {
		return nil, err
	}
	tenancy.Tenants = tenants

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
