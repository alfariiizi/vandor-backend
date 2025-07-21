package cmd_enum_types

// EnumValue represents a single enum value with its metadata
type EnumValue struct {
	Name        string // Original key from YAML (e.g., "PENDING")
	ConstName   string // Go constant name (e.g., "OrderStatusPending")
	FieldName   string // Struct field name (e.g., "Pending")
	Description string // Human readable description
}

// TemplateData holds all data needed for template generation
type TemplateData struct {
	Type        string
	Package     string
	Values      []EnumValue
	Description string
}
