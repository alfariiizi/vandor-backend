package utils

import (
	"fmt"
	"reflect"
)

// Required checks if the provided named arguments are non-zero / non-empty.
// Returns a descriptive error like: `argument "Name" is required and empty`.
func Required(fields map[string]any) error {
	for name, val := range fields {
		if val == nil {
			return fmt.Errorf(`argument "%s" is required and nil`, name)
		}
		v := reflect.ValueOf(val)

		// Handle pointer values
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return fmt.Errorf(`argument "%s" is required and nil`, name)
			}
			v = v.Elem()
		}

		// Check for zero value
		if v.IsZero() {
			return fmt.Errorf(`argument "%s" is required and empty`, name)
		}
	}
	return nil
}
