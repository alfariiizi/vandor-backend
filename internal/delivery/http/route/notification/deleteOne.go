package notification_handler

import (
	"context"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/notification"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

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

type deleteOneInput struct {
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type deleteOneOutput types.OutputResponseData[deleteOneData]

type deleteOneData struct {
	Message string `json:"message" doc:"Response message" example:"Notification deleted successfully"`
}

type deleteOneHandler model.HTTPHandler[deleteOneInput, deleteOneOutput]

type deleteOne struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewdeleteOne(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) deleteOneHandler {
	h := &deleteOne{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteOne) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/notifications/{id}", method.Operation{
		Summary:     "Delete One Notification",
		Description: "Delete a notification by ID",
		Tags:        []string{"Notification"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteOne) GenerateResponse(data deleteOneData) *deleteOneOutput {
	return (*deleteOneOutput)(types.GenerateOutputResponseData(data))
}

func (h *deleteOne) Handler(ctx context.Context, input *deleteOneInput) (*deleteOneOutput, error) {
	token, err := utils.ExtractJWT(ctx)
	if err != nil {
		return nil, err
	}

	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}

	err = h.client.Notification.DeleteOneID(*id).
		Where(notification.UserID(token.UserID)).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(deleteOneData{
		Message: "Notification deleted successfully",
	}), nil
}
