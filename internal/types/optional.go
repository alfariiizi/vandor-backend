package types

import (
	"encoding/json"
	"reflect"
)

// Optional represents a value that can be undefined, zero, or have an actual value
type Optional[T any] struct {
	value   *T
	defined bool
}

// NewOptional creates a new Optional with a defined value
func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		value:   &value,
		defined: true,
	}
}

// NewUndefined creates a new undefined Optional
func NewUndefined[T any]() Optional[T] {
	return Optional[T]{
		value:   nil,
		defined: false,
	}
}

// IsDefined returns true if the value is defined (even if it's zero)
func (o Optional[T]) IsDefined() bool {
	return o.defined
}

// IsUndefined returns true if the value is undefined
func (o Optional[T]) IsUndefined() bool {
	return !o.defined
}

// Value returns the value and a boolean indicating if it's defined
func (o Optional[T]) ValueOrZero() T {
	if o.defined && o.value != nil {
		return *o.value
	}
	var zero T
	return zero
}

// ValueOr returns the value if defined, otherwise returns the default
func (o Optional[T]) ValueOr(defaultValue T) T {
	if o.defined && o.value != nil && !isZeroValue(*o.value) {
		return *o.value
	}
	return defaultValue
}

func (o Optional[T]) Value() *T {
	if o.defined && o.value != nil && !isZeroValue(*o.value) {
		return o.value
	}
	return nil
}

// IsZero returns true if the value is defined and is the zero value for its type
func (o Optional[T]) IsZero() bool {
	if !o.defined || o.value == nil {
		return false
	}
	return isZeroValue(*o.value)
}

// IsDefinedAndNotZero returns true if the value is defined and is not the zero value
func (o Optional[T]) IsDefinedAndNotZero() bool {
	return o.defined && o.value != nil && !isZeroValue(*o.value)
}

// IsDefinedAndNotNil returns true if the value is defined and the pointer is not nil
func (o Optional[T]) IsDefinedAndNotNil() bool {
	return o.defined && o.value != nil
}

// HasValue is an alias for IsDefinedAndNotNil for clearer semantics
func (o Optional[T]) HasValue() bool {
	return o.IsDefinedAndNotNil()
}

// HasNonZeroValue is an alias for IsDefinedAndNotZero for clearer semantics
func (o Optional[T]) HasNonZeroValue() bool {
	return o.IsDefinedAndNotZero()
}

// MarshalJSON implements JSON marshaling for Huma
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.defined {
		return []byte("null"), nil
	}
	if o.value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*o.value)
}

// isZeroValue checks if a value is the zero value for its type using reflection
func isZeroValue(v any) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0.0
	case reflect.Complex64, reflect.Complex128:
		return rv.Complex() == 0
	case reflect.String:
		return rv.String() == ""
	case reflect.Slice, reflect.Map, reflect.Chan:
		return rv.IsNil() || rv.Len() == 0
	case reflect.Array:
		// For arrays, check if all elements are zero
		for i := range rv.Len() {
			if !rv.Index(i).IsZero() {
				return false
			}
		}
		return true
	case reflect.Struct:
		// For structs, check if all fields are zero
		return rv.IsZero()
	case reflect.Ptr, reflect.Interface:
		return rv.IsNil()
	case reflect.Func:
		return rv.IsNil()
	default:
		// For other types, use reflection's IsZero if available
		return rv.IsZero()
	}
}

// UnmarshalJSON implements JSON unmarshaling for Huma
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*o = NewUndefined[T]()
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	*o = NewOptional(value)
	return nil
}
