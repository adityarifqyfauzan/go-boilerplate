package helper

type ApiResponse struct {
	Code       int         `json:"code"`
	Message    any         `json:"message"`
	Data       any         `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Error      error       `json:"-"`
}

func NewApiResponse(httpCode int, message any, data any) *ApiResponse {
	return &ApiResponse{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
}

func NewApiResponseWithPagination(httpCode int, message any, data any, pagination *Pagination) *ApiResponse {
	return &ApiResponse{
		Code:       httpCode,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}
