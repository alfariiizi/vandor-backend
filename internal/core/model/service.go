package model

import (
	"context"
)

type Service[I any, O any] interface {
	Execute(ctx context.Context, input I) (*O, error)
}
