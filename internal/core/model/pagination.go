package model

import "github.com/alfariiizi/vandor/internal/types"

type PaginationInput struct {
	Page  types.Optional[int] `json:"page"`
	Limit types.Optional[int] `json:"limit"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type PaginationOutput[I any] struct {
	Meta PaginationMeta `json:"meta"`
	Data I              `json:"data"`
}
