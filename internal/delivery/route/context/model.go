package httpctx

import (
	"fmt"
	"time"
)

type baseResponse struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Signature string `json:"signature"`
}

type SuccessResponse struct {
	Data any `json:"data"`
	baseResponse
}

type SuccessMessageResponse struct {
	Message any `json:"message"`
	baseResponse
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
	baseResponse
}

func generateBaseResponse(status string, code int) baseResponse {
	return baseResponse{
		Status:    status,
		Code:      code,
		Signature: fmt.Sprintf(" ©%d Alfarizi", time.Now().Year()),
	}
}
