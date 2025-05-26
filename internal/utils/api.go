package utils

import (
	"fmt"
	"time"
)

type BaseResponse struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Signature string `json:"signature"`
}

type ResponseData[T any] struct {
	Data T `json:"data"`
	BaseResponse
}

type ResponseMessage struct {
	Message string `json:"message"`
	BaseResponse
}

type OutputResponseData[Body any] struct {
	Body ResponseData[Body]
}

type OutputResponseMessage struct {
	Body ResponseMessage
}

func GenerateBaseResponse(status string, code int) BaseResponse {
	return BaseResponse{
		Status:    status,
		Code:      code,
		Signature: fmt.Sprintf("©%d go-services", time.Now().Year()),
	}
}

func GenerateResponseData[T any](status string, code int, data T) ResponseData[T] {
	return ResponseData[T]{
		Data:         data,
		BaseResponse: GenerateBaseResponse(status, code),
	}
}

func GenerateResponseMessage(status string, code int, message string) ResponseMessage {
	return ResponseMessage{
		Message:      message,
		BaseResponse: GenerateBaseResponse(status, code),
	}
}

func GenerateOutputResponseData[Body any](body Body) *OutputResponseData[Body] {
	return &OutputResponseData[Body]{
		Body: GenerateResponseData("success", 200, body),
	}
}

func GenerateOutputResponseMessage(message string) *OutputResponseMessage {
	return &OutputResponseMessage{
		Body: GenerateResponseMessage("success", 200, message),
	}
}
