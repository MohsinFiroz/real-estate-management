package tenancy

import (
	"real-estate-management/pkg/id"
	"real-estate-management/pkg/validator"
	"time"

	"gorm.io/gorm"
)

type RentFrequency string

const (
	Weekly      RentFrequency = "Weekly"
	Fortnightly RentFrequency = "Fortnightly"
	Monthly     RentFrequency = "Monthly"
)

type Status string

const (
	Active     Status = "Active"
	Inactive   Status = "Inactive"
	BondRefund Status = "BondRefund"
	SACAT      Status = "SACAT"
)

// Tenancy model representing a tenancy record in the database
type Tenancy struct {
	ID              string        `json:"id" gorm:"column:id"`
	Status          Status        `json:"status" gorm:"column:status" validate:"oneof=Active Inactive BondRefund SACAT"`
	PropertyID      string        `json:"propertyID" gorm:"column:property_id" validate:"required"`
	PrimaryTenantID string        `json:"primaryTenantID" gorm:"column:primary_tenant_id"`
	Rent            int           `json:"rent" gorm:"column:rent"`
	BondAmount      int           `json:"bondAmount" gorm:"column:bond_amount"`
	BondID          string        `json:"bondID" gorm:"column:bond_id"`
	RentFrequency   RentFrequency `json:"rentFrequency" gorm:"column:rent_frequency" validate:"required,oneof=Weekly Fortnightly Monthly"`
	StartDate       string        `json:"startDate" gorm:"column:start_date"` // 30/02/2025 format
	EndDate         string        `json:"endDate" gorm:"column:end_date"`     // 30/02/2025 format
	Notes           string        `json:"notes" gorm:"column:notes"`
	Notices         string        `json:"notices" gorm:"column:notices"`
	CreatedAt       *time.Time    `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt       *time.Time    `json:"updatedAt" gorm:"column:updated_at"`
}

// TableName specifies the table name for the Tenancy model
func (*Tenancy) TableName() string {
	return "tenancies"
}

// BeforeCreate hook to validate and generate ULID for the ID field before saving the tenancy
func (t *Tenancy) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate Tenancy
	err = validator.ValidateStruct(t)
	if err != nil {
		return err
	}

	// Create ULID
	t.ID = id.NewString()

	return nil
}

// BeforeUpdate hook to update the UpdatedAt timestamp before updating the tenancy
func (t *Tenancy) BeforeUpdate(tx *gorm.DB) (err error) {
	// Add updated at time
	now := time.Now()
	t.UpdatedAt = &now

	return nil
}

// SortColumnMap is sort column mappings from camelCase to snake_case
var SortColumnMap = map[string]string{
	"id":              "id",
	"propertyID":      "property_id",
	"primaryTenantID": "primary_tenant_id",
	"rent":            "rent",
	"bondAmount":      "bond_amount",
	"bondID":          "bond_id",
	"rentFrequency":   "rent_frequency",
	"startDate":       "start_date",
	"endDate":         "end_date",
	"notes":           "notes",
	"notice":          "notices",
	"createdAt":       "created_at",
	"updatedAt":       "updated_at",
}
