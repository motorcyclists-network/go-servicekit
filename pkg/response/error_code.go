package response

// ErrorCode tüm mikroservislerde kullanılan standart hata kodlarını tanımlar.
// Frontend ve diğer servisler bu kodlara göre davranış belirler.
type ErrorCode string

const (
	// Client Hataları (4xx)
	CodeValidationError  ErrorCode = "VALIDATION_ERROR"   // İstek gövdesi veya parametreler hatalı
	CodeNotFound         ErrorCode = "NOT_FOUND"           // İstenen kaynak bulunamadı
	CodeAlreadyExists    ErrorCode = "ALREADY_EXISTS"      // Kaynak zaten mevcut (unique constraint vs.)
	CodeUnauthorized     ErrorCode = "UNAUTHORIZED"        // Kimlik doğrulaması gerekli
	CodeForbidden        ErrorCode = "FORBIDDEN"           // Yetkisiz erişim
	CodeBadRequest       ErrorCode = "BAD_REQUEST"         // Genel istemci hatası
	CodeTooManyRequests  ErrorCode = "TOO_MANY_REQUESTS"   // Rate limit aşıldı
	CodeConflict         ErrorCode = "CONFLICT"            // Çakışma (concurrent update vs.)

	// Server Hataları (5xx)
	CodeInternalError    ErrorCode = "INTERNAL_ERROR"      // Bilinmeyen sunucu hatası
	CodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE" // Servis geçici olarak kullanılamıyor
	CodeDatabaseError    ErrorCode = "DATABASE_ERROR"      // Veritabanı işlemi başarısız
	CodeExternalService  ErrorCode = "EXTERNAL_SERVICE_ERROR" // Dış servis çağrısı başarısız
)
