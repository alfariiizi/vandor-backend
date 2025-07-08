package model

import (
	"context"
)

type Usecase[I any, O any] interface {
	Execute(ctx context.Context, input I) (*O, error)
}
