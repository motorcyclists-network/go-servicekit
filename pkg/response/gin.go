package response

import "github.com/gin-gonic/gin"

func SendError(ctx *gin.Context, status int, apiErr APIError) {
	apiErr.TraceID = ctx.GetString("X-Request-ID")

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
