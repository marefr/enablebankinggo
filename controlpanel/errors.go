package controlpanel

import "errors"

// ErrorResponse represents an error response from the API.
type ErrorResponse struct {
	ErrorObj struct {
		Code    int              `json:"code"`
		Message string           `json:"message"`
		Errors  []map[string]any `json:"errors"`
	} `json:"error"`
}

// Error implements the error interface for ErrorResponse.
func (e ErrorResponse) Error() string {
	return e.ErrorObj.Message
}

// IsErrorResponse checks if the provided error is of type [ErrorResponse] and
// returns it along with a boolean indicating the result.
func IsErrorResponse(err error) (*ErrorResponse, bool) {
	var errorResp *ErrorResponse
	if errors.As(err, &errorResp) {
		return errorResp, true
	}

	return nil, false
}
