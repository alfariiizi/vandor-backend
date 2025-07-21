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

type LoginAdminPayload LoginPayload

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

type LoginAdminInput struct {
	// JSON body for POST
	Body LoginAdminPayload `json:"body" contentType:"application/json"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type LoginAdminOutput types.OutputResponseData[LoginAdminData]

type LoginAdminData *auth_service.LoginOutput

type LoginAdminHandler model.HTTPHandler[LoginAdminInput, LoginAdminOutput]

type loginAdmin struct {
	api     huma.API
	service *service.Services
}

func NewLoginAdmin(
	api *api.HttpApi,
	service *service.Services,
) LoginAdminHandler {
	h := &loginAdmin{
		api:     api.BaseAPI,
		service: service,
	}
	h.RegisterRoutes()
	return h
}

func (h *loginAdmin) RegisterRoutes() {
	api := h.api
	method.POST(api, "/admin/login", method.Operation{
		Summary:     "Login Admin",
		Description: "Login as an admin user",
		Tags:        []string{"Auth"},
		BearerAuth:  false,
	}, h.Handler)
}

func (h *loginAdmin) GenerateResponse(data LoginAdminData) *LoginAdminOutput {
	return (*LoginAdminOutput)(types.GenerateOutputResponseData(data))
}

func (h *loginAdmin) Handler(ctx context.Context, input *LoginAdminInput) (*LoginAdminOutput, error) {
	payload := input.Body
	res, err := h.service.Auth.Login.Execute(ctx, auth_service.LoginInput{
		Email:    payload.Email,
		Password: payload.Password,
		IsAdmin:  false,
	})
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(LoginAdminData(res)), nil
}
