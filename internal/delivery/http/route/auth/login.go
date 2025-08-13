package auth_handler

import (
	"context"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	auth_service "github.com/alfariiizi/vandor/internal/core/service/auth"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/danielgtaylor/huma/v2"
)

type LoginPayload struct {
	Email    string `json:"email" doc:"Email of the user" required:"true"`
	Password string `json:"password" doc:"Password of the user" example:"password123" required:"true"`
}

// NOTE:
// Hint Tags for input parameters
// @ref: https://huma.rocks/features/request-inputs
//
// Tag       | Description                           | Example
// -------------------------------------------------------------------
// path      | Name of the path parameter            | path:"thing-id"
// query     | Name of the query string parameter    | query:"q"
// header    | Name of the header parameter          | header:"Authorization"
// cookie    | Name of the cookie parameter          | cookie:"session"
// required  | Mark a query/header param as required | required:"true"

type LoginInput struct {
	// JSON body for POST
	Body LoginPayload `json:"body" contentType:"application/json"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type LoginOutput types.OutputResponseData[LoginData]

type LoginData *auth_service.LoginOutput

type LoginHandler model.HTTPHandler[LoginInput, LoginOutput]

type login struct {
	api     huma.API
	service *service.Services
}

func NewLogin(
	api *api.HttpApi,
	service *service.Services,
) LoginHandler {
	h := &login{
		api:     api.BaseAPI,
		service: service,
	}
	h.RegisterRoutes()
	return h
}

func (h *login) RegisterRoutes() {
	api := h.api
	method.POST(api, "/auth/login", method.Operation{
		Summary:     "Login",
		Description: "Login handler",
		Tags:        []string{"Auth"},
		BearerAuth:  false,
	}, h.Handler)
}

func (h *login) Handler(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	payload := input.Body
	res, err := h.service.Auth.Login.Execute(ctx, auth_service.LoginInput{
		Email:    payload.Email,
		Password: payload.Password,
		// IsAdmin:  false,
	})
	if err != nil {
		return nil, err
	}

	return (*LoginOutput)(types.GenerateOutputResponseData(LoginData(res))), nil
}
