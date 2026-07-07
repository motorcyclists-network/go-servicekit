package response

import "github.com/gin-gonic/gin"

// SendError hata yanıtı döner.
// Middleware'de set edilen X-Request-ID varsa otomatik olarak traceId'ye eklenir.
func SendError(ctx *gin.Context, status int, apiErr APIError) {
	apiErr.TraceID = ctx.GetString("X-Request-ID")

	ctx.JSON(status, errorResponse{
		Error: apiErr,
	})
}

// SendSuccess başarılı yanıt döner — tek kaynak (single resource) için.
func SendSuccess(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, successResponse{
		Data: data,
	})
}

// SendSuccessWithMeta başarılı yanıt döner — liste + sayfalama (pagination) için.
func SendSuccessWithMeta(ctx *gin.Context, status int, data interface{}, meta PaginationMeta) {
	ctx.JSON(status, successResponse{
		Data: data,
		Meta: meta,
	})
}
