package handler

import httpctx "github.com/alfariiizi/go-service/internal/delivery/route/context"

type CommonHandler struct{}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{}
}

func (h *CommonHandler) GetHealth(ctx httpctx.HttpContext) error {
	return ctx.SendSuccessResponse(200, map[string]string{"status": "ok"})
}

func (h *CommonHandler) GetPing(ctx httpctx.HttpContext) error {
	return ctx.SendSuccessResponse(200, map[string]string{"message": "pong"})
}
