package utils

type ApiError struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
	Code      string `json:"code"`
}
