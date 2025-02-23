package id

import (
	"github.com/oklog/ulid/v2"
)

// New generates a ULID using Make()
func New() ulid.ULID {
	return ulid.Make()
}

// NewString generates a ULID string
func NewString() string {
	return New().String()
}
