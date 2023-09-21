package utils

import (
	"github.com/adityarifqyfauzan/go-boilerplate/app/response"
	"github.com/gin-gonic/gin"
)

func WriteToResponseBody(ctx *gin.Context, webResponse response.WebResponse) {
	ctx.JSON(webResponse.Code, webResponse)
}
