package model

import (
	"context"

	"github.com/alfariiizi/go-service/internal/types"
)

type Service[I any, O any] interface {
	Execute(ctx context.Context, input I) types.Result[O]
}
