package auth_handler

import (
	"context"
	"net/http"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/enum"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/user"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

type registerPayload struct {
	Email     string `json:"email" doc:"email of user" example:"john@gmail.com" required:"true"`
	FirstName string `json:"first_name" doc:"first name of user" example:"John" required:"true"`
	LastName  string `json:"last_name" doc:"last name of user" example:"Doe" required:"true"`
	Password  string `json:"password" doc:"password of user" example:"password123" required:"true"`
	Role      string `json:"role" doc:"role of user" example:"user" required:"true"`
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

type registerInput struct {
	// JSON body for POST
	Body registerPayload `json:"body" contentType:"application/json"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type registerOutput types.OutputResponseData[authRegisterData]

type authRegisterData struct {
	User *db.User `json:"user" doc:"registered user data"`
}

type registerHandler model.HTTPHandler[registerInput, registerOutput]

type register struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func Newregister(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) registerHandler {
	h := &register{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *register) RegisterRoutes() {
	api := h.api
	method.POST(api, "/auth/register", method.Operation{
		Summary:     "Register User",
		Description: "Register a new user in the system",
		Tags:        []string{"Auth"},
		BearerAuth:  true,
		RoleAllowed: []enum.UserRole{enum.UserRoleSuperAdmin, enum.UserRoleAdmin},
	}, h.Handler)
}

func (h *register) GenerateResponse(data authRegisterData) *registerOutput {
	return (*registerOutput)(types.GenerateOutputResponseData(data))
}

func (h *register) Handler(ctx context.Context, input *registerInput) (*registerOutput, error) {
	payload := input.Body

	pashwordHash, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to hash password: "+err.Error())
	}

	var userRole user.Role
	switch payload.Role {
	case "admin":
		userRole = user.RoleADMIN
	case "superadmin":
		userRole = user.RoleSUPERADMIN
	case "user":
		userRole = user.RoleUSER
	default:
		return nil, huma.NewError(http.StatusBadRequest, "Invalid role: "+payload.Role)
	}

	userOne, err := h.client.User.Create().
		SetEmail(payload.Email).
		SetFirstName(payload.FirstName).
		SetLastName(payload.LastName).
		SetPasswordHash(*pashwordHash).
		SetRole(userRole).
		Save(ctx)
	if err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to create user: "+err.Error())
	}

	return h.GenerateResponse(authRegisterData{
		User: userOne,
	}), nil
}
