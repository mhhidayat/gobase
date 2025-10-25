package response

type errorResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ErrorResponse(message string) *errorResponse[any] {
	return &errorResponse[any]{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
}
