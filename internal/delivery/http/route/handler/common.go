package handler

import (
	"context"

	"github.com/alfariiizi/go-service/internal/utils"
)

type CommonHandler struct{}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{}
}

type HealthOutput struct {
	Status string `json:"status"`
}

func (h *CommonHandler) GetHealth(ctx context.Context, input *struct{}) (*utils.OutputResponseData[HealthOutput], error) {
	return utils.GenerateOutputResponseData(HealthOutput{
		Status: "ok",
	}), nil
}

func (h *CommonHandler) GetPing(ctx context.Context, input *struct{}) (*utils.OutputResponseMessage, error) {
	return utils.GenerateOutputResponseMessage("pong"), nil
}
