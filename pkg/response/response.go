package response

import "net/http"

type APIError struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
	TraceID string      `json:"traceId,omitempty"`
}

type ErrorResponse struct {
	Error APIError `json:"error"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"currentPage"`
	PageSize    int   `json:"pageSize"`
	TotalItems  int64 `json:"totalItems"`
	TotalPages  int   `json:"totalPages"`
}

func newError(code ErrorCode, message string, details interface{}) APIError {
	return APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func BadRequest(message string, details interface{}) (int, APIError) {
	return http.StatusBadRequest, newError(CodeBadRequest, message, details)
}

func ValidationError(message string, details interface{}) (int, APIError) {
	return http.StatusBadRequest, newError(CodeValidationError, message, details)
}

func NotFound(message string) (int, APIError) {
	return http.StatusNotFound, newError(CodeNotFound, message, nil)
}

func AlreadyExists(message string) (int, APIError) {
	return http.StatusConflict, newError(CodeAlreadyExists, message, nil)
}

func Unauthorized(message string) (int, APIError) {
	return http.StatusUnauthorized, newError(CodeUnauthorized, message, nil)
}

func Forbidden(message string) (int, APIError) {
	return http.StatusForbidden, newError(CodeForbidden, message, nil)
}

func Conflict(message string, details interface{}) (int, APIError) {
	return http.StatusConflict, newError(CodeConflict, message, details)
}

func TooManyRequests(message string) (int, APIError) {
	return http.StatusTooManyRequests, newError(CodeTooManyRequests, message, nil)
}

func InternalError(message string) (int, APIError) {
	return http.StatusInternalServerError, newError(CodeInternalError, message, nil)
}

func DatabaseError(message string) (int, APIError) {
	return http.StatusInternalServerError, newError(CodeDatabaseError, message, nil)
}

func ServiceUnavailable(message string) (int, APIError) {
	return http.StatusServiceUnavailable, newError(CodeServiceUnavailable, message, nil)
}

func ExternalServiceError(message string, details interface{}) (int, APIError) {
	return http.StatusBadGateway, newError(CodeExternalService, message, details)
}
