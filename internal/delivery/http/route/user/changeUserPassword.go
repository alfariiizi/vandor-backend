package user_handler

import (
	"context"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
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

type changeUserPasswordPayload struct {
	CurrentPassword string `json:"current_password" doc:"Current password of the user" example:"currentpassword123" required:"true"`
	NewPassword     string `json:"new_password" doc:"New password of the user" example:"newpassword123" required:"true"`
	ConfirmPassword string `json:"confirm_password" doc:"Confirm new password of the user" example:"newpassword123" required:"true"`
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

type changeUserPasswordInput struct {
	// JSON body for POST
	Body changeUserPasswordPayload `json:"body" contentType:"application/json"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type changeUserPasswordOutput types.OutputResponseMessage

type changeUserPasswordHandler model.HTTPHandler[changeUserPasswordInput, changeUserPasswordOutput]

type changeUserPassword struct {
	api     huma.API
	service *service.Services
	client  *db.Client
	domain  *domain_entries.Domain
}

func NewchangeUserPassword(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
	domain *domain_entries.Domain,
) changeUserPasswordHandler {
	h := &changeUserPassword{
		api:     api.BaseAPI,
		service: service,
		client:  client,
		domain:  domain,
	}
	h.RegisterRoutes()
	return h
}

func (h *changeUserPassword) RegisterRoutes() {
	api := h.api
	method.POST(api, "/users/change-password", method.Operation{
		Summary:     "Change User Password",
		Description: "Change the password of the authenticated user.",
		Tags:        []string{"User"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *changeUserPassword) GenerateResponse(message string) *changeUserPasswordOutput {
	return (*changeUserPasswordOutput)(types.GenerateOutputResponseMessage(message))
}

func (h *changeUserPassword) Handler(ctx context.Context, input *changeUserPasswordInput) (*changeUserPasswordOutput, error) {
	token, err := utils.ExtractJWT(ctx)
	if err != nil {
		return nil, err
	}

	payload := input.Body

	if payload.NewPassword != payload.ConfirmPassword {
		return nil, huma.Error400BadRequest("New password and confirm password do not match")
	}

	userOne, err := h.domain.User.One(h.client.User.Query().Where(user.ID(token.UserID)).Only(ctx))
	if err != nil {
		return nil, err
	}
	if !userOne.IsPasswordMatches(payload.CurrentPassword) {
		return nil, huma.Error400BadRequest("Current password is incorrect")
	}

	passwordHash, err := utils.HashPassword(payload.NewPassword)
	if err != nil {
		return nil, err
	}

	_, err = h.client.User.Update().
		Where(user.ID(token.UserID)).
		SetPasswordHash(*passwordHash).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(
		"Success change user password",
	), nil
}
