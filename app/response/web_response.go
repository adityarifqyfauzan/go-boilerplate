package response

import "net/http"

type WebResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func New(statusCode int, message string, data any) WebResponse {
	return WebResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
}

func Success(message string, data any) WebResponse {
	return WebResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
