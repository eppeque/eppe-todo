package models

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewError(message string) *ErrorResponse {
	return &ErrorResponse{message}
}
