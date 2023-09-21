package middleware

import (
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/exception"
	"github.com/gin-gonic/gin"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				exception.ErrorHandler(ctx, err)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
