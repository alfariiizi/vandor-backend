package route

import (
	"github.com/danielgtaylor/huma/v2"
)

type HttpRoute struct {
	api huma.API
}

func NewHttpRoute(api huma.API) *HttpRoute {
	return &HttpRoute{
		api: api,
	}
}

// type Docs struct {
//
// }
//
// func (r *HttpRoute) GET[I any, O any](path string, summary string, description string, tags []string) {
// 	huma.Register(r.api, huma.Operation{
// 		OperationID: "GET" + path,
// 		Method:     http.MethodGet,
// 		Path:      path,
// 		Summary:   summary,
// 		Description: description,
// 		Tags:      tags,
// 	}, )
// }
