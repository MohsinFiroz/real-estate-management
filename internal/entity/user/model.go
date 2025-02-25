package user

import (
	"gorm.io/gorm"
	"real-estate-management/pkg/id"
	"real-estate-management/pkg/validator"
	"time"
)

// Role defines the user role type
type Role string

// User role constants
const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// User model representing a user record in the database
type User struct {
	ID        string     `json:"id" gorm:"column:id"`
	Email     string     `json:"email" gorm:"column:email" validate:"required,email"`
	Password  string     `json:"password" gorm:"column:password" validate:"required,min=8"`
	FirstName string     `json:"firstName" gorm:"column:first_name" validate:"required"`
	LastName  string     `json:"lastName" gorm:"column:last_name"`
	Phone     string     `json:"phone" gorm:"column:phone" validate:"omitempty,e164" `
	Role      Role       `json:"role" gorm:"column:role" validate:"required,oneof=admin user"`
	IsActive  bool       `json:"isActive" gorm:"column:is_active"`
	LastLogin *time.Time `json:"lastLogin" gorm:"column:last_login"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

// TableName specifies the table name for the User model
func (*User) TableName() string {
	return "users"
}

// BeforeCreate hook to validate and generate ULID for the ID field before saving the user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate User
	err = validator.ValidateStruct(u)
	if err != nil {
		return err
	}

	// Create ULID
	u.ID = id.NewString()

	return nil
}

// BeforeUpdate hook to validate that the ID is a valid ULID before updating the user
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Validate User fields
	err = validator.ValidateStruct(u)
	if err != nil {
		return err
	}

	// Add updated at time
	now := time.Now()
	u.UpdatedAt = &now

	return nil
}

// SortColumnMap is sort column mappings from camelCase to snake_case
var SortColumnMap = map[string]string{
	"firstName": "first_name",
	"lastName":  "last_name",
	"email":     "email",
	"createdAt": "created_at",
	"updatedAt": "updated_at",
}
