package exception

import (
	"net/http"

	"github.com/adityarifqyfauzan/go-boilerplate/app/response"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context, err any) {
	if notFoundError(ctx, err) {
		return
	}
	if credentialError(ctx, err) {
		return
	}
	if unprocessableEntity(ctx, err) {
		return
	}
	internalServerError(ctx, err)
}

func unprocessableEntity(ctx *gin.Context, err any) bool {
	exception, ok := err.(UnprocessableEntityException)
	if !ok {
		return false
	}

	webResponse := response.New(http.StatusUnprocessableEntity, exception.Error, nil)
	utils.WriteToResponseBody(ctx, webResponse)

	return true
}

func notFoundError(ctx *gin.Context, err any) bool {
	exception, ok := err.(NotFoundException)
	if !ok {
		return false
	}

	webResponse := response.New(http.StatusNotFound, exception.Error, nil)
	utils.WriteToResponseBody(ctx, webResponse)

	return true
}

func credentialError(ctx *gin.Context, err any) bool {
	exception, ok := err.(CredentialException)
	if !ok {
		return false
	}

	webResponse := response.New(http.StatusUnauthorized, exception.Error, nil)
	utils.WriteToResponseBody(ctx, webResponse)

	return true
}

func internalServerError(ctx *gin.Context, err any) bool {
	webResponse := response.New(http.StatusInternalServerError, "Internal Server Error: Please try in few minutes", nil)
	utils.WriteToResponseBody(ctx, webResponse)
	return true
}
