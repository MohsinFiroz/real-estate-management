package validator

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

// The init function will be automatically called when this package is imported
func init() {
	if v == nil {
		v = validator.New()
	}
}

// ValidateStruct validates the struct using the global validator
func ValidateStruct(s interface{}) error {
	return v.Struct(s)
}
