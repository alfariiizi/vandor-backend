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

type markAsReadPayload struct {
	ID string `json:"id" doc:"ID of the notification" example:"cd230e27-5f9e-442d-a582-0abafd587130" required:"true"`
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

type markAsReadInput struct {
	// JSON body for POST
	Body markAsReadPayload `json:"body" contentType:"application/json"`
}

// NOTE:
// You can use this for output
// types.OutputResponseData[T] -> for data responses
// types.OutputResponseMessage -> for message responses
// types.OutputResponsePagination[T] -> for paginated responses

type markAsReadOutput types.OutputResponseData[markAsReadData]

type markAsReadData struct {
	Message string           `json:"message" doc:"Response message" example:"Notification marked as read successfully"`
	Data    *db.Notification `json:"data,omitempty" doc:"Notification data"`
}

type markAsReadHandler model.HTTPHandler[markAsReadInput, markAsReadOutput]

type markAsRead struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewmarkAsRead(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) markAsReadHandler {
	h := &markAsRead{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *markAsRead) RegisterRoutes() {
	api := h.api
	method.POST(api, "/notifications/mark-as-read", method.Operation{
		Summary:     "Mark Notification as Read",
		Description: "Mark a notification as read by its ID",
		Tags:        []string{"Notification"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *markAsRead) GenerateResponse(data markAsReadData) *markAsReadOutput {
	return (*markAsReadOutput)(types.GenerateOutputResponseData(data))
}

func (h *markAsRead) Handler(ctx context.Context, input *markAsReadInput) (*markAsReadOutput, error) {
	token, err := utils.ExtractJWT(ctx)
	if err != nil {
		return nil, err
	}

	payload := input.Body
	id, err := utils.IDParser(payload.ID)
	if err != nil {
		return nil, err
	}

	notif, err := h.client.Notification.UpdateOneID(*id).
		Where(notification.UserID(token.UserID)).
		SetRead(true).
		SetReadAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return h.GenerateResponse(markAsReadData{
		Message: "Notification marked as read successfully",
		Data:    notif,
	}), nil
}
