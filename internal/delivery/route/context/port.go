package httpctx

type HttpContext interface {
	// GetPath returns the path of the HTTP request.
	GetPath() string
	// GetMethod returns the HTTP method of the request.
	GetMethod() string

	// GetHeaders returns the headers of the HTTP request.
	GetHeaders() map[string]string
	// GetHeader returns a specific header from the HTTP request.
	GetHeader(key string, required bool) (string, error)

	// GetBody returns the body of the HTTP request.
	BindBody(v any) error

	// GetParams returns all parameters from the HTTP request.
	GetParams() map[string]string
	// GetParam returns a specific parameter from the HTTP request.
	GetParam(key string, required bool) (string, error)

	// GetQueryParams returns the query parameters of the HTTP request.
	GetQueryParams() map[string]string
	// GetQueryParam returns a specific query parameter from the HTTP request.
	GetQueryParam(key string, required bool) (string, error)

	SendSuccessResponse(statusCode int, data any) error
	SendSuccessMessageResponse(statusCode int, message string) error
	SendErrorResponse(statusCode int, message string, error error) error
}
