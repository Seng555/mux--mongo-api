// model/response.go
package model

import "net/http"

// Response represents a JSON response structure
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	// Add additional fields as needed
	ErrorCode int    `json:"errorCode,omitempty"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(status int, errorMsg string) *Response {
	return &Response{
		Status:    status,
		Message:   "Error",
		ErrorMsg:  errorMsg,
		ErrorCode: 1001, // Assign a custom error code as needed
	}
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}
