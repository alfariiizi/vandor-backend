package httpport

import "context"

type HttpServer interface {
	Start() error
	Stop(context.Context) error
	RegisterHandler() error
}
