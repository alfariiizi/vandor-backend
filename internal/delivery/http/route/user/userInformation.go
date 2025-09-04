package user_handler

import (
	"context"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/user"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

type userInformationPayload struct {
	FirstName string `json:"first_name" doc:"First name of the user" example:"John" required:"true" minLength:"1" maxLength:"80"`
	LastName  string `json:"last_name" doc:"Last name of the user" example:"Doe" required:"true" minLength:"1" maxLength:"80"`
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

type userInformationInput struct {
	// JSON body for POST
	Body userInformationPayload `json:"body" contentType:"application/json"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type userInformationOutput types.OutputResponseMessage

type userInformationHandler model.HTTPHandler[userInformationInput, userInformationOutput]

type userInformation struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewuserInformation(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) userInformationHandler {
	h := &userInformation{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *userInformation) RegisterRoutes() {
	api := h.api
	method.PUT(api, "/users/information", method.Operation{
		Summary:     "Update User Information",
		Description: "Update the first name and last name of the authenticated user.",
		Tags:        []string{"User"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *userInformation) GenerateResponse(message string) *userInformationOutput {
	return (*userInformationOutput)(types.GenerateOutputResponseMessage(message))
}

func (h *userInformation) Handler(ctx context.Context, input *userInformationInput) (*userInformationOutput, error) {
	token, err := utils.ExtractJWT(ctx)
	if err != nil {
		return nil, err
	}

	_, err = h.client.User.Update().
		Where(user.ID(token.UserID)).
		SetFirstName(input.Body.FirstName).
		SetLastName(input.Body.LastName).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(
		"Success update user information",
	), nil
}
