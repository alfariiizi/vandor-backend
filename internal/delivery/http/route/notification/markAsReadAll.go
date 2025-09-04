package notification_handler

import (
	"context"
	"time"

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

type markAsReadAllPayload struct{}

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

type markAsReadAllInput struct {
	// JSON body for POST
	Body markAsReadAllPayload `json:"body" contentType:"application/json"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type markAsReadAllOutput types.OutputResponseData[markAsReadAllData]

type markAsReadAllData struct {
	Message string `json:"message" doc:"Response message" example:"All notifications marked as read successfully"`
}

type markAsReadAllHandler model.HTTPHandler[markAsReadAllInput, markAsReadAllOutput]

type markAsReadAll struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewmarkAsReadAll(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) markAsReadAllHandler {
	h := &markAsReadAll{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *markAsReadAll) RegisterRoutes() {
	api := h.api
	method.POST(api, "/notifications/mark-as-read-all", method.Operation{
		Summary:     "Mark All Notifications as Read",
		Description: "Mark all notifications as read for the authenticated user",
		Tags:        []string{"Notification"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *markAsReadAll) GenerateResponse(data markAsReadAllData) *markAsReadAllOutput {
	return (*markAsReadAllOutput)(types.GenerateOutputResponseData(data))
}

func (h *markAsReadAll) Handler(ctx context.Context, input *markAsReadAllInput) (*markAsReadAllOutput, error) {
	token, err := utils.ExtractJWT(ctx)
	if err != nil {
		return nil, err
	}

	_, err = h.client.Notification.Update().
		Where(notification.UserID(token.UserID)).
		SetRead(true).
		SetReadAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(markAsReadAllData{
		Message: "All notifications marked as read successfully",
	}), nil
}
