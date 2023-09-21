package handler

import (
	"github.com/adityarifqyfauzan/go-boilerplate/app/request"
	"github.com/adityarifqyfauzan/go-boilerplate/app/usecase"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/logger"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/utils"
	"github.com/gin-gonic/gin"
)

type HelloWorldHandler interface {
	Hello(ctx *gin.Context)
	Exception(ctx *gin.Context)
}

type helloWorldHandler struct {
	uc usecase.HelloWorldUsecase
}

func NewHelloWorldHandler(uc usecase.HelloWorldUsecase) HelloWorldHandler {
	return &helloWorldHandler{
		uc: uc,
	}
}

func (h *helloWorldHandler) Hello(ctx *gin.Context) {

	// get request body
	var req request.HelloWorldRequest
	req.Name = ctx.Param("name")

	// call usecase
	resp := h.uc.SayHello(req)
	utils.WriteToResponseBody(ctx, resp)
}

func (h *helloWorldHandler) Exception(ctx *gin.Context) {

	var req request.ExceptionRequest
	if err := ctx.Bind(&req); err != nil {
		logger.Console.Info("Bad Request")
		panic("error")
	}

	resp := h.uc.ExceptionUsecase(req)
	utils.WriteToResponseBody(ctx, resp)
}
