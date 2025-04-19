package tenant

import (
	"real-estate-management/pkg/id"
	"real-estate-management/pkg/validator"
	"time"

	"gorm.io/gorm"
)

type CommunicationMedium string

const (
	SMS      CommunicationMedium = "SMS"
	WeChat   CommunicationMedium = "WeChat"
	WhatsApp CommunicationMedium = "WhatsApp"
)

// Tenant model representing a tenant record in the database
type Tenant struct {
	ID                  string              `json:"id" gorm:"column:id"`
	Name                string              `json:"name" gorm:"column:name" validate:"required"`
	Mobile              string              `json:"mobile" gorm:"column:mobile"`
	Email               string              `json:"email" gorm:"column:email"`
	CommunicationMedium CommunicationMedium `json:"communicationMedium" gorm:"column:communication_medium" validate:"omitempty,oneof=SMS WeChat WhatsApp"`
	IsActive            bool                `json:"isActive" gorm:"column:is_active;default:true"`
	CreatedAt           *time.Time          `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt           *time.Time          `json:"updatedAt" gorm:"column:updated_at"`
}

// TableName specifies the table name for the Tenant model
func (*Tenant) TableName() string {
	return "tenants"
}

// BeforeCreate hook to validate and generate ULID for the ID field before saving the tenant
func (t *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate Tenant
	err = validator.ValidateStruct(t)
	if err != nil {
		return err
	}

	// Create ULID
	t.ID = id.NewString()

	return nil
}

// BeforeUpdate hook to update the UpdatedAt timestamp before updating the tenant
func (t *Tenant) BeforeUpdate(tx *gorm.DB) (err error) {
	// Add updated at time
	now := time.Now()
	t.UpdatedAt = &now

	return nil
}

// SortColumnMap is sort column mappings from camelCase to snake_case
var SortColumnMap = map[string]string{
	"name":                "name",
	"mobile":              "mobile",
	"email":               "email",
	"communicationMedium": "communication_medium",
	"isActive":            "is_active",
	"createdAt":           "created_at",
	"updatedAt":           "updated_at",
}
