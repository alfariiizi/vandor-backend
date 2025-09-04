package notification_handler

import (
	"context"
	"math"

	"entgo.io/ent/dialect/sql"
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

type ListNotificationFilterInput struct {
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListNotificationInput struct {
	model.HTTPInputParamsPagination
	ListNotificationFilterInput
}

type ListNotificationOutput types.OutputResponsePagination[ListNotificationData]

type ListNotificationData []*db.Notification

type ListNotificationHandler model.HTTPHandler[ListNotificationInput, ListNotificationOutput]

type listNotification struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListNotification(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListNotificationHandler {
	h := &listNotification{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listNotification) RegisterRoutes() {
	api := h.api
	method.GET(api, "/notifications", method.Operation{
		Summary:     "List Notifications",
		Description: "Retrieve a list of notifications",
		Tags:        []string{"Notification"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listNotification) Handler(ctx context.Context, input *ListNotificationInput) (*ListNotificationOutput, error) {
	token, err := utils.ExtractJWT(ctx)
	if err != nil {
		return nil, err
	}

	page := 1
	limit := 10
	if input.Page > 0 {
		page = input.Page
	}
	if input.Limit > 0 {
		limit = input.Limit
	}

	query := h.FilterQuery(h.client.Notification.Query(), input.ListNotificationFilterInput)

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	items, err := query.
		Where(notification.UserID(token.UserID)).
		Order(notification.ByCreatedAt(sql.OrderDesc())).
		Offset((page - 1) * limit).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var totalPages int
	if limit == 0 {
		totalPages = 1
	}
	totalPages = int(math.Ceil(float64(total) / float64(limit)))
	if totalPages == 0 {
		totalPages = 1
	}

	data := make(ListNotificationData, len(items))
	copy(data, items)
	res := types.GenerateOutputPaginationData(
		data, page, limit, totalPages, total,
	)

	return (*ListNotificationOutput)(&res), nil
}

func (h *listNotification) FilterQuery(query *db.NotificationQuery, filter ListNotificationFilterInput) *db.NotificationQuery {
	return query.Where()
}
