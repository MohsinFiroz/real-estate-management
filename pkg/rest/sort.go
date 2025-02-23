package rest

import (
	"github.com/pkg/errors"
	"strings"
)

const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

// SortField represents the field and its direction for sorting
type SortField struct {
	Field     string
	Direction string
}

// ParseSortFields parses the sortBy string into a list of SortField structs
func ParseSortFields(allowedColumnMap map[string]string, sortBy string) ([]SortField, error) {
	var sortFields []SortField

	// Trim the input and handle empty case
	sortBy = strings.TrimSpace(sortBy)
	if sortBy == "" {
		return []SortField{
			{Field: "created_at", Direction: SortDesc},
		}, nil
	}

	// Parse the sortBy string
	for _, field := range strings.Split(sortBy, ",") {
		// Split field by ":" to extract field name and direction
		fieldParts := strings.Split(strings.TrimSpace(field), ":")
		fieldName := fieldParts[0]
		direction := SortDesc // default direction is DESC

		// If direction is provided, validate it
		if len(fieldParts) == 2 {
			direction = fieldParts[1]
			if direction != SortAsc && direction != SortDesc {
				return nil, errors.Errorf("invalid direction '%s' for field '%s'", direction, fieldName)
			}
		}

		// Map field name from camelCase to snake_case using pre-defined map
		if snakeField, exists := allowedColumnMap[fieldName]; exists {
			sortFields = append(sortFields, SortField{Field: snakeField, Direction: direction})
		} else {
			return nil, errors.Errorf("invalid sort field '%s'", fieldName)
		}
	}

	return sortFields, nil
}
