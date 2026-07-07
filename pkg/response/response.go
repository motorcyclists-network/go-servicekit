package response

import "net/http"

// APIError tek bir hatayı temsil eden yapıdır.
// Stripe ve Google Cloud API standartlarından esinlenilmiştir.
type APIError struct {
	Code    ErrorCode   `json:"code"`              // Makinenin okuyabileceği sabit hata kodu
	Message string      `json:"message"`           // İnsanın okuyabileceği açıklama
	Details interface{} `json:"details,omitempty"` // Ek detaylar (alan bazlı validation hataları vs.)
	TraceID string      `json:"traceId,omitempty"` // Distributed tracing için istek takip ID'si
}

// errorResponse hata yanıtının JSON zarfıdır (envelope).
type errorResponse struct {
	Error APIError `json:"error"`
}

// successResponse başarılı yanıtın JSON zarfıdır (envelope).
type successResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

// PaginationMeta liste endpoint'lerinde sayfalama bilgilerini taşır.
type PaginationMeta struct {
	CurrentPage int   `json:"currentPage"`
	PageSize    int   `json:"pageSize"`
	TotalItems  int64 `json:"totalItems"`
	TotalPages  int   `json:"totalPages"`
}

// --- Hata Helper Fonksiyonları ---

// newError ortak hata oluşturma mantığıdır.
func newError(code ErrorCode, message string, details interface{}) APIError {
	return APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// BadRequest 400 — genel istemci hatası
func BadRequest(message string, details interface{}) (int, APIError) {
	return http.StatusBadRequest, newError(CodeBadRequest, message, details)
}

// ValidationError 400 — istek gövdesi veya parametrelerdeki doğrulama hatası
func ValidationError(message string, details interface{}) (int, APIError) {
	return http.StatusBadRequest, newError(CodeValidationError, message, details)
}

// NotFound 404 — kaynak bulunamadı
func NotFound(message string) (int, APIError) {
	return http.StatusNotFound, newError(CodeNotFound, message, nil)
}

// AlreadyExists 409 — kaynak zaten mevcut
func AlreadyExists(message string) (int, APIError) {
	return http.StatusConflict, newError(CodeAlreadyExists, message, nil)
}

// Unauthorized 401 — kimlik doğrulaması gerekli
func Unauthorized(message string) (int, APIError) {
	return http.StatusUnauthorized, newError(CodeUnauthorized, message, nil)
}

// Forbidden 403 — yetkisiz erişim
func Forbidden(message string) (int, APIError) {
	return http.StatusForbidden, newError(CodeForbidden, message, nil)
}

// Conflict 409 — çakışma
func Conflict(message string, details interface{}) (int, APIError) {
	return http.StatusConflict, newError(CodeConflict, message, details)
}

// TooManyRequests 429 — rate limit aşıldı
func TooManyRequests(message string) (int, APIError) {
	return http.StatusTooManyRequests, newError(CodeTooManyRequests, message, nil)
}

// InternalError 500 — sunucu hatası
func InternalError(message string) (int, APIError) {
	return http.StatusInternalServerError, newError(CodeInternalError, message, nil)
}

// DatabaseError 500 — veritabanı hatası
func DatabaseError(message string) (int, APIError) {
	return http.StatusInternalServerError, newError(CodeDatabaseError, message, nil)
}

// ServiceUnavailable 503 — servis geçici olarak kullanılamıyor
func ServiceUnavailable(message string) (int, APIError) {
	return http.StatusServiceUnavailable, newError(CodeServiceUnavailable, message, nil)
}

// ExternalServiceError 502 — dış servis çağrısı başarısız
func ExternalServiceError(message string, details interface{}) (int, APIError) {
	return http.StatusBadGateway, newError(CodeExternalService, message, details)
}
