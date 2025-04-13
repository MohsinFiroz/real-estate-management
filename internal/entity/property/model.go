package property

import (
	"gorm.io/gorm"
	"real-estate-management/pkg/id"
	"real-estate-management/pkg/validator"
	"time"
)

// Property model representing a property record in the database
type Property struct {
	ID                   string     `json:"id" gorm:"column:id"`
	OwnerID              string     `json:"ownerID" gorm:"column:owner_id" validate:"required"`
	Address              string     `json:"address" gorm:"column:address" validate:"required"`
	Suburb               string     `json:"suburb" gorm:"column:suburb" validate:"required"`
	Postcode             string     `json:"postcode" gorm:"column:postcode" validate:"required"`
	KeyNo                string     `json:"keyNo" gorm:"column:key_no"`
	IsActive             bool       `json:"isActive" gorm:"column:is_active"`
	ManagementFee        float64    `json:"managementFee" gorm:"column:management_fee"`
	ManagementStartDate  string     `json:"managementStartDate" gorm:"column:management_start_date"`
	ManagementEndDate    string     `json:"managementEndDate" gorm:"column:management_end_date"`
	WaterBillAccount     string     `json:"waterBillAccount" gorm:"column:water_bill_account"`
	LastWaterBillReading float64    `json:"lastWaterBillReading" gorm:"column:last_water_bill_reading"`
	Notes                string     `json:"notes" gorm:"column:notes"`
	Other                string     `json:"other" gorm:"column:other"`
	CreatedAt            *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt            *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

// TableName specifies the table name for the Property model
func (*Property) TableName() string {
	return "properties"
}

// BeforeCreate hook to validate and generate ULID for the ID field before saving the property
func (p *Property) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate Property
	err = validator.ValidateStruct(p)
	if err != nil {
		return err
	}

	// Create ULID
	p.ID = id.NewString()

	return nil
}

// BeforeUpdate hook to update the UpdatedAt timestamp before updating the property
func (p *Property) BeforeUpdate(tx *gorm.DB) (err error) {
	// Add updated at time
	now := time.Now()
	p.UpdatedAt = &now

	return nil
}

// SortColumnMap is sort column mappings from camelCase to snake_case
var SortColumnMap = map[string]string{
	"ownerID":              "owner_id",
	"address":              "address",
	"suburb":               "suburb",
	"postcode":             "postcode",
	"keyNo":                "key_no",
	"isActive":             "is_active",
	"managementFee":        "management_fee",
	"managementStartDate":  "management_start_date",
	"managementEndDate":    "management_end_date",
	"waterBillAccount":     "water_bill_account",
	"lastWaterBillReading": "last_water_bill_reading",
	"createdAt":            "created_at",
	"updatedAt":            "updated_at",
}
