package http

import "context"

type HttpServer interface {
	Start() error
	Stop(context.Context) error
}
