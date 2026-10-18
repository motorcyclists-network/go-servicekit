package response

import (
	"github.com/gin-gonic/gin"
	"github.com/motorcyclists-network/go-servicekit/pkg/i18n"
)

func SendError(ctx *gin.Context, status int, apiErr APIError) {
	apiErr.TraceID = ctx.GetString("X-Request-ID")

	userLang := ctx.GetString(i18n.ContextKeyLanguage)
	apiErr.Message = i18n.Translate(userLang, string(apiErr.Code), apiErr.Message)

	ctx.JSON(status, ErrorResponse{
		Error: apiErr,
	})
}

func SendSuccess(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, SuccessResponse{
		Data: data,
	})
}

func SendSuccessWithMeta(ctx *gin.Context, status int, data interface{}, meta PaginationMeta) {
	ctx.JSON(status, SuccessResponse{
		Data: data,
		Meta: meta,
	})
}
