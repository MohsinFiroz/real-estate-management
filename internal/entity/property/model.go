package property

import (
	"gorm.io/gorm"
	"real-estate-management/pkg/id"
	"real-estate-management/pkg/validator"
	"time"
)

// Property model representing a property record in the database
type Property struct {
	ID                       string     `json:"id" gorm:"column:id"`
	OwnerID                  string     `json:"ownerId" gorm:"column:owner_id" validate:"required"`
	Address                  string     `json:"address" gorm:"column:address" validate:"required"`
	Suburb                   string     `json:"suburb" gorm:"column:suburb" validate:"required"`
	Postcode                 string     `json:"postcode" gorm:"column:postcode" validate:"required"`
	TenantID                 *string    `json:"tenantId" gorm:"column:tenant_id"`
	KeyNo                    string     `json:"keyNo" gorm:"column:key_no"`
	Rent                     float64    `json:"rent" gorm:"column:rent" validate:"required,gt=0"`
	BondAmount               float64    `json:"bondAmount" gorm:"column:bond_amount"`
	BondID                   string     `json:"bondId" gorm:"column:bond_id"`
	ManagementFee            float64    `json:"managementFee" gorm:"column:management_fee"`
	CommunicationMedium      string     `json:"communicationMedium" gorm:"column:communication_medium"`
	RoutineInspection        string     `json:"routineInspection" gorm:"column:routine_inspection"`
	ManagementAgreementStart time.Time  `json:"managementAgreementStart" gorm:"column:management_agreement_start"`
	ManagementAgreementEnd   time.Time  `json:"managementAgreementEnd" gorm:"column:management_agreement_end"`
	WaterBillAccount         string     `json:"waterBillAccount" gorm:"column:water_bill_account"`
	LastWaterBillReading     float64    `json:"lastWaterBillReading" gorm:"column:last_waterbill_reading"`
	PropertyNote             string     `json:"propertyNote" gorm:"column:property_note"`
	Other                    string     `json:"other" gorm:"column:other"`
	CreatedAt                *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt                *time.Time `json:"updatedAt" gorm:"column:updated_at"`
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

// BeforeUpdate hook to validate that the ID is a valid ULID before updating the property
func (p *Property) BeforeUpdate(tx *gorm.DB) (err error) {
	// Add updated at time
	now := time.Now()
	p.UpdatedAt = &now

	return nil
}

// SortColumnMap is sort column mappings from camelCase to snake_case
var SortColumnMap = map[string]string{
	"ownerId":                  "owner_id",
	"address":                  "address",
	"suburb":                   "suburb",
	"postcode":                 "postcode",
	"tenantId":                 "tenant_id",
	"keyNo":                    "key_no",
	"rent":                     "rent",
	"bondAmount":               "bond_amount",
	"bondId":                   "bond_id",
	"managementFee":            "management_fee",
	"communicationMedium":      "communication_medium",
	"routineInspection":        "routine_inspection",
	"managementAgreementStart": "management_agreement_start",
	"managementAgreementEnd":   "management_agreement_end",
	"waterBillAccount":         "water_bill_account",
	"lastWaterBillReading":     "last_waterbill_reading",
	"propertyNote":             "property_note",
	"other":                    "other",
	"createdAt":                "created_at",
	"updatedAt":                "updated_at",
}
