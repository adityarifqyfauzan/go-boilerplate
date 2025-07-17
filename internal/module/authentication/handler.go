package authentication

import (
	"net/http"

	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type handler struct {
	service Service
}

func NewHandler(
	service Service,
) handler {
	return handler{
		service: service,
	}
}

func (h *handler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.NewApiResponse(http.StatusBadRequest, "Invalid request body", nil))
		return
	}

	validate := validator.New(c.Value("localizer").(*i18n.Localizer))
	errors := validate.Validate(request)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, helper.NewApiResponse(http.StatusBadRequest, validate.FirstError(errors), nil))
		return
	}

	response := h.service.Login(c, request)
	c.JSON(response.Code, response)
}

func (h *handler) Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.NewApiResponse(http.StatusBadRequest, "Invalid request body", nil))
		return
	}

	validate := validator.New(c.Value("localizer").(*i18n.Localizer))
	errors := validate.Validate(request)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, helper.NewApiResponse(http.StatusBadRequest, validate.FirstError(errors), nil))
		return
	}

	response := h.service.Register(c, request)
	c.JSON(response.Code, response)
}

func (h *handler) ForgotPassword(c *gin.Context) {
	// TODO: Implement forgot password functionality
	c.JSON(http.StatusNotImplemented, helper.NewApiResponse(http.StatusNotImplemented, "Forgot password not implemented yet", nil))
}

// refresh token
func (h *handler) RefreshToken(c *gin.Context) {
	var request RefreshTokenRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.NewApiResponse(http.StatusBadRequest, "Invalid request body", nil))
		return
	}

	validate := validator.New(c.Value("localizer").(*i18n.Localizer))
	errors := validate.Validate(request)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, helper.NewApiResponse(http.StatusBadRequest, validate.FirstError(errors), nil))
		return
	}

	response := h.service.RefreshToken(c, request.RefreshToken)
	c.JSON(response.Code, response)
}

func (h *handler) Me(c *gin.Context) {
	response := h.service.Me(c)
	c.JSON(response.Code, response)
}
