package handler

import (
	httpport "github.com/alfariiizi/go-echo-fx-template/internal/delivery/port/http"
)

func HelloWorldHandler(ctx httpport.HttpContext) error {
	// Get the path of the HTTP request
	path := ctx.GetPath()

	// Get the method of the HTTP request
	method := ctx.GetMethod()

	// Get the headers of the HTTP request
	headers := ctx.GetHeaders()

	// Get the query parameters of the HTTP request
	queryParams := ctx.GetQueryParams()

	// Get the body of the HTTP request
	// body := ctx.GetBody()

	// Create a response object
	response := map[string]any{
		"path":        path,
		"method":      method,
		"headers":     headers,
		"queryParams": queryParams,
		// "body":        string(body),
	}

	// Send a success response with status code 200 and the response object
	return ctx.SendSuccessResponse(200, response)
}

func HelloWorldHandlerWithParams(ctx httpport.HttpContext) error {
	// Get the path of the HTTP request
	path := ctx.GetPath()

	// Get the method of the HTTP request
	method := ctx.GetMethod()

	// Get the headers of the HTTP request
	headers := ctx.GetHeaders()

	// _, err := ctx.GetHeader("Authorization", true)
	// fmt.Println("Authorization Header:", err)
	// if err != nil {
	// 	return err
	// }

	// body, ok := ctx.GetBody().(string)
	// if !ok {
	// 	return ctx.SendErrorResponse(400, "Invalid body format", nil)
	// }

	var body struct {
		Status string `json:"status" validate:"required"`
	}
	if err := ctx.BindBody(&body); err != nil {
		return err
	}

	// Get the query parameters of the HTTP request
	queryParams := ctx.GetQueryParams()

	// Get a specific parameter from the HTTP request
	id, err := ctx.GetParam("id", true)
	if err != nil {
		return err
	}

	// Create a response object
	response := map[string]any{
		"path":        path,
		"method":      method,
		"headers":     headers,
		"queryParams": queryParams,
		"param":       id,
		"body":        body,
	}

	// Send a success response with status code 200 and the response object
	return ctx.SendSuccessResponse(200, response)
}
