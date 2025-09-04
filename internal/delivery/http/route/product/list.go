package analyze_handler

import (
	"context"
	"math"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/product"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

type ListProductFilterInput struct {
	Query string `query:"q" doc:"Search query for product name or description" example:"laptop"`
}

type ListProductInput struct {
	model.HTTPInputParamsPagination
	ListProductFilterInput
}

type ListProductOutput types.OutputResponsePagination[ListProductData]

type ListProductData []*db.Product

type ListProductHandler model.HTTPHandler[ListProductInput, ListProductOutput]

type listProduct struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListProduct(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListProductHandler {
	h := &listProduct{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listProduct) RegisterRoutes() {
	api := h.api
	method.GET(api, "/products", method.Operation{
		Summary:     "List of Products",
		Description: "Get a list of products with pagination and filtering options.",
		Tags:        []string{"Product"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listProduct) Handler(ctx context.Context, input *ListProductInput) (*ListProductOutput, error) {
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

	query := h.FilterQuery(h.client.Product.Query(), input.ListProductFilterInput)

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	items, err := query.
		Where(product.UserID(token.UserID)).
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

	data := make(ListProductData, len(items))
	copy(data, items)
	res := types.GenerateOutputPaginationData(
		data, page, limit, totalPages, total,
	)

	return (*ListProductOutput)(&res), nil
}

func (h *listProduct) FilterQuery(query *db.ProductQuery, filter ListProductFilterInput) *db.ProductQuery {
	return query
}
