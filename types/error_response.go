package types

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Name     uint32         `json:"code"`
	Message  string         `json:"message"`
}

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
