package middleware

import (
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/translator"
	"github.com/gin-gonic/gin"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "en"
		}
		localizer := translator.NewLocalizer(lang)
		c.Set(translator.LOCALIZER, localizer)
		c.Next()
	}
}
