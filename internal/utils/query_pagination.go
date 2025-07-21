package utils

import (
	"context"
	"fmt"

	"github.com/alfariiizi/vandor/internal/core/model"
)

func PaginateQuery[T any](
	ctx context.Context,
	countFn func(context.Context) (int, error),
	allFn func(context.Context, int, int) ([]T, error),
	pagination model.PaginationInput,
) ([]T, model.PaginationMeta, error) {
	page := pagination.Page.ValueOr(1)
	limit := pagination.Limit.ValueOr(10)
	offset := (page - 1) * limit

	totalCount, err := countFn(ctx)
	if err != nil {
		return nil, model.PaginationMeta{}, fmt.Errorf("count failed: %w", err)
	}

	items, err := allFn(ctx, offset, limit)
	if err != nil {
		return nil, model.PaginationMeta{}, fmt.Errorf("query failed: %w", err)
	}

	meta := model.PaginationMeta{
		Page:       page,
		PerPage:    limit,
		TotalItems: totalCount,
		TotalPages: (totalCount + limit - 1) / limit,
	}

	return items, meta, nil
}
