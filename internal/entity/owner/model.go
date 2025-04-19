package owner

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

// Owner model representing an owner record in the database
type Owner struct {
	ID                  string              `json:"id" gorm:"column:id"`
	Name                string              `json:"name" gorm:"column:name" validate:"required"`
	Mobile              string              `json:"mobile" gorm:"column:mobile"`
	Email               string              `json:"email" gorm:"column:email"`
	CommunicationMedium CommunicationMedium `json:"communicationMedium" gorm:"column:communication_medium" validate:"omitempty,oneof=SMS WeChat WhatsApp"`
	Insurance           string              `json:"insurance" gorm:"column:insurance"`
	AccountNumber       string              `json:"accountNumber" gorm:"column:account_number"`
	BSB                 string              `json:"bsb" gorm:"column:bsb"`
	Identification      string              `json:"identification" gorm:"column:identification"`
	Address             string              `json:"address" gorm:"column:address"`
	Notes               string              `json:"notes" gorm:"column:notes"`
	IsActive            bool                `json:"isActive" gorm:"column:is_active;default:true"`
	CreatedAt           *time.Time          `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt           *time.Time          `json:"updatedAt" gorm:"column:updated_at"`
}

// TableName specifies the table name for the Owner model
func (*Owner) TableName() string {
	return "owners"
}

// BeforeCreate hook to validate and generate ULID for the ID field before saving the owner
func (o *Owner) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate Owner
	err = validator.ValidateStruct(o)
	if err != nil {
		return err
	}

	// Create ULID
	o.ID = id.NewString()

	return nil
}

// BeforeUpdate hook to update the UpdatedAt timestamp before updating the owner
func (o *Owner) BeforeUpdate(tx *gorm.DB) (err error) {
	// Add updated at time
	now := time.Now()
	o.UpdatedAt = &now

	return nil
}

// SortColumnMap is sort column mappings from camelCase to snake_case
var SortColumnMap = map[string]string{
	"name":                "name",
	"mobile":              "mobile",
	"email":               "email",
	"communicationMedium": "communication_medium",
	"insurance":           "insurance",
	"accountNumber":       "account_number",
	"bsb":                 "bsb",
	"identification":      "identification",
	"address":             "address",
	"isActive":            "is_active",
	"createdAt":           "created_at",
	"updatedAt":           "updated_at",
}
