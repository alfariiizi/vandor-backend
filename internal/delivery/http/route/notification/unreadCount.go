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

type unreadCountInput struct{}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type unreadCountOutput types.OutputResponseData[unreadCountData]

type unreadCountData struct {
	Count int `json:"count" doc:"Number of unread notifications" example:"5"`
}

type unreadCountHandler model.HTTPHandler[unreadCountInput, unreadCountOutput]

type unreadCount struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewunreadCount(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) unreadCountHandler {
	h := &unreadCount{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *unreadCount) RegisterRoutes() {
	api := h.api
	method.GET(api, "/notifications/unread-count", method.Operation{
		Summary:     "Get Unread Notification Count",
		Description: "Retrieve the count of unread notifications for the authenticated user",
		Tags:        []string{"Notification"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *unreadCount) GenerateResponse(data unreadCountData) *unreadCountOutput {
	return (*unreadCountOutput)(types.GenerateOutputResponseData(data))
}

func (h *unreadCount) Handler(ctx context.Context, input *unreadCountInput) (*unreadCountOutput, error) {
	token, err := utils.ExtractJWT(ctx)
	if err != nil {
		return nil, err
	}
	count, err := h.client.Notification.Query().
		Where(notification.UserID(token.UserID)).
		Where(notification.Read(false)).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(unreadCountData{
		Count: count,
	}), nil
}
