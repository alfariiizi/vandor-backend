package types

import (
	"fmt"
	"time"

	"github.com/alfariiizi/vandor/config"
)

type Meta struct {
	Status    string `json:"status" doc:"Response status" example:"success"`
	Code      int    `json:"code" doc:"Response status code" example:"200"`
	Signature string `json:"signature" doc:"Response signature"`
	Version   string `json:"version" doc:"Application version" example:"1.0.0"`
}

type MetaPagination struct {
	Pagination Pagination `json:"pagination" doc:"Pagination information"`
	Status     string     `json:"status" doc:"Response status" example:"success"`
	Code       int        `json:"code" doc:"Response status code" example:"200"`
	Signature  string     `json:"signature" doc:"Response signature"`
	Version    string     `json:"version" doc:"Application version" example:"1.0.0"`
}

type Pagination struct {
	Page       int `json:"page" doc:"Current page number" example:"1"`
	PerPage    int `json:"per_page" doc:"Number of items per page" example:"10"`
	TotalPages int `json:"total_pages" doc:"Total number of pages" example:"5"`
	TotalItems int `json:"total_items" doc:"Total number of items" example:"42"`
}

// --- Response Types ---

type Response[T any] struct {
	Data T    `json:"data" doc:"Response data"`
	Meta Meta `json:"meta" doc:"Response metadata"`
}

type ResponseMessage struct {
	Message string `json:"message" doc:"Response message"`
	Meta    Meta   `json:"meta" doc:"Response metadata"`
}

type ResponsePagination[Data any] struct {
	Data Data           `json:"data" doc:"Response data"`
	Meta MetaPagination `json:"meta" doc:"Response metadata with pagination"`
}

// --- Output Types ---

type OutputResponseData[Body any] struct {
	Body Response[Body]
}

type OutputResponseMessage struct {
	Body ResponseMessage
}

type OutputResponsePagination[Data any] struct {
	Body ResponsePagination[Data]
}

// --- Response Generation Functions ---

func GenerateBaseResponse(status string, code int) Meta {
	cfg := config.GetConfig()
	return Meta{
		Status:    status,
		Code:      code,
		Signature: fmt.Sprintf("©%d %s", time.Now().Year(), cfg.App.SignatureResponse),
		Version:   cfg.App.Version,
	}
}

func GenerateResponseData[T any](status string, code int, data T) Response[T] {
	return Response[T]{
		Data: data,
		Meta: GenerateBaseResponse(status, code),
	}
}

func GenerateResponseMessage(status string, code int, message string) ResponseMessage {
	return ResponseMessage{
		Message: message,
		Meta:    GenerateBaseResponse(status, code),
	}
}

func GenerateOutputResponseData[Data any](data Data) *OutputResponseData[Data] {
	return &OutputResponseData[Data]{
		Body: GenerateResponseData("success", 200, data),
	}
}

func GenerateOutputResponseMessage(message string) *OutputResponseMessage {
	return &OutputResponseMessage{
		Body: GenerateResponseMessage("success", 200, message),
	}
}

func GenerateOutputPaginationData[Data any](data Data, page int, perPage int, totalPage int, totalItems int) OutputResponsePagination[Data] {
	return OutputResponsePagination[Data]{
		Body: ResponsePagination[Data]{
			Data: data,
			Meta: MetaPagination{
				Pagination: Pagination{
					Page:       page,
					PerPage:    perPage,
					TotalPages: totalPage,
					TotalItems: totalItems,
				},
				Status:    "success",
				Code:      200,
				Signature: fmt.Sprintf("©%d %s", time.Now().Year(), config.GetConfig().App.SignatureResponse),
				Version:   config.GetConfig().App.Version,
			},
		},
	}
}
