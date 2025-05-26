package handler

import (
	"context"

	"github.com/alfariiizi/go-service/internal/delivery/http/route"
)

type CommonHandler struct{}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{}
}

type HealthOutput struct {
	Status string `json:"status"`
}

func (h *CommonHandler) GetHealth(ctx context.Context, input *struct{}) (*route.OutputResponseData[HealthOutput], error) {
	return route.GenerateOutputResponseData(HealthOutput{
		Status: "ok",
	}), nil
}

func (h *CommonHandler) GetPing(ctx context.Context, input *struct{}) (*route.OutputResponseMessage, error) {
	return route.GenerateOutputResponseMessage("pong"), nil
}
