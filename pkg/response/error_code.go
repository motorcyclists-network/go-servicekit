package response

type ErrorCode string

const (
	CodeValidationError ErrorCode = "VALIDATION_ERROR"
	CodeNotFound        ErrorCode = "NOT_FOUND"
	CodeAlreadyExists   ErrorCode = "ALREADY_EXISTS"
	CodeUnauthorized    ErrorCode = "UNAUTHORIZED"
	CodeForbidden       ErrorCode = "FORBIDDEN"
	CodeBadRequest      ErrorCode = "BAD_REQUEST"
	CodeTooManyRequests ErrorCode = "TOO_MANY_REQUESTS"
	CodeConflict        ErrorCode = "CONFLICT"

	CodeInternalError      ErrorCode = "INTERNAL_ERROR"
	CodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
	CodeDatabaseError      ErrorCode = "DATABASE_ERROR"
	CodeExternalService    ErrorCode = "EXTERNAL_SERVICE_ERROR"
)
