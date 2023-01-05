package formatter

import (
	"go-campaign-app/model/web"
)

func WriteToResponseBody(code int, st string, msg string, data interface{}) web.APIResponse {
	meta := web.Meta{
		Code:    code,
		Status:  st,
		Message: msg,
	}

	resp := web.APIResponse{
		Meta: meta,
		Data: data,
	}

	return resp
}
