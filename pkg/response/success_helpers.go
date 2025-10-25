package response

type successResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func SuccessResponse(message string, data any) *successResponse[any] {
	return &successResponse[any]{
		Status:  "success",
		Message: message,
		Data:    nil,
	}
}
