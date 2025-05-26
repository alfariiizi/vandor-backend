package httpctx

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type httpContext struct {
	context echo.Context
}

func NewHttpContext(c echo.Context) HttpContext {
	return &httpContext{
		context: c,
	}
}

func (h *httpContext) GetPath() string {
	return h.context.Path()
}

func (h *httpContext) GetMethod() string {
	return h.context.Request().Method
}

func (h *httpContext) GetHeaders() map[string]string {
	headers := make(map[string]string)
	for key, values := range h.context.Request().Header {
		headers[key] = values[0]
	}
	return headers
}

func (h *httpContext) GetHeader(key string, required bool) (string, error) {
	value := h.context.Request().Header.Get(key)
	if value == "" && required {
		return "", errors.New("header " + key + " is required")
	}
	return value, nil
}

func (h *httpContext) GetParams() map[string]string {
	params := make(map[string]string)
	for _, name := range h.context.ParamNames() {
		params[name] = h.context.Param(name)
	}
	return params
}

func (h *httpContext) GetParam(key string, required bool) (string, error) {
	value := h.context.Param(key)
	if value == "" && required {
		h.SendErrorResponse(400, "Parameter "+key+" is required", nil)
		return "", errors.New("parameter " + key + " is required")
	}
	return value, nil
}

func (h *httpContext) GetQueryParams() map[string]string {
	queryParams := make(map[string]string)
	for key, values := range h.context.QueryParams() {
		queryParams[key] = values[0]
	}
	return queryParams
}

func (h *httpContext) GetQueryParam(key string, required bool) (string, error) {
	value := h.context.QueryParam(key)
	if value == "" && required {
		h.SendErrorResponse(400, "Query parameter "+key+" is required", nil)
		return "", errors.New("query parameter " + key + " is required")
	}
	return value, nil
}

func (h *httpContext) BindBody(v any) error {
	if err := h.context.Bind(v); err != nil {
		return errors.New("failed to bind body")
	}
	if err := h.context.Validate(v); err != nil {
		return errors.New("validation failed")
	}
	return nil
}

func (h *httpContext) SendSuccessResponse(statusCode int, data any) error {
	return h.context.JSON(statusCode, SuccessResponse{
		Data:         data,
		baseResponse: generateBaseResponse("success", statusCode),
	})
}

func (h *httpContext) SendSuccessMessageResponse(statusCode int, message string) error {
	return h.context.JSON(statusCode, SuccessMessageResponse{
		Message:      message,
		baseResponse: generateBaseResponse("success", statusCode),
	})
}

func (h *httpContext) SendErrorResponse(statusCode int, message string, error error) error {
	return h.context.JSON(statusCode, ErrorResponse{
		Message:      message,
		Error:        error,
		baseResponse: generateBaseResponse("error", statusCode),
	})
}
