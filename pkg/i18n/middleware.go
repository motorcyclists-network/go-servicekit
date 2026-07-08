package i18n

import "github.com/gin-gonic/gin"

const ContextKeyLanguage = "Accept-Language"

func LanguageMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.GetHeader("Accept-Language")

		if lang == "" {
			lang = "en"
		}

		ctx.Set(ContextKeyLanguage, lang)
		ctx.Next()
	}
}
