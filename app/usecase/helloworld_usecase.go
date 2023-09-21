package usecase

import (
	"fmt"
	"net/http"

	"github.com/adityarifqyfauzan/go-boilerplate/app/request"
	"github.com/adityarifqyfauzan/go-boilerplate/app/response"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/exception"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/logger"
)

type HelloWorldUsecase interface {
	SayHello(req request.HelloWorldRequest) response.WebResponse
	ExceptionUsecase(req request.ExceptionRequest) response.WebResponse
}

type helloWorldUsecase struct {
}

func NewHelloWorldUsecase() HelloWorldUsecase {
	return &helloWorldUsecase{}
}

// define all message here
// variable naming usecaseName + function name + message
var (
	exceptionUsecaseNotFoundErr      = "not found exception"
	exceptionUsecaseUnprocessableErr = "unprocessable entity"
	exceptionUsecaseSuccess          = "berhasil apa?"
)

func (u *helloWorldUsecase) SayHello(req request.HelloWorldRequest) response.WebResponse {
	name := req.Name

	if name == "" {
		name = "World!"
	}

	logger.Console.Info("say hello invoked")
	return response.New(http.StatusOK, fmt.Sprintf("Hello, %s", name), nil)
}

func (u *helloWorldUsecase) ExceptionUsecase(req request.ExceptionRequest) response.WebResponse {

	// not found error handling
	if req.ExceptionType == "not_found" {
		logger.Console.Error(exceptionUsecaseNotFoundErr)
		panic(exception.NewNotFoundException(exceptionUsecaseNotFoundErr))
	}

	if req.ExceptionType == "unprocessable" {
		logger.Console.Error(exceptionUsecaseUnprocessableErr)
		panic(exception.NewUnprocessableEntityException(exceptionUsecaseUnprocessableErr))
	}

	return response.Success("No Error Exception", nil)
}
