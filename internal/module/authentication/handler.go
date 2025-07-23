package authentication

import (
	"net/http"

	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	tr := otel.Tracer("authentication-handler")
	ctx, span := tr.Start(c, "LoginHandler")
	defer span.End()

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

	span.AddEvent("Register User", trace.WithAttributes(
		attribute.String("email", request.Email),
	))

	response := h.service.Login(ctx, request)
	c.JSON(response.Code, response)
}

func (h *handler) Register(c *gin.Context) {
	tr := otel.Tracer("authentication-handler")
	ctx, span := tr.Start(c, "RegisterHandler")
	defer span.End()

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

	span.AddEvent("Register User", trace.WithAttributes(
		attribute.String("email", request.Email),
	))

	response := h.service.Register(ctx, request)
	c.JSON(response.Code, response)
}

func (h *handler) ForgotPassword(c *gin.Context) {
	// TODO: Implement forgot password functionality
	c.JSON(http.StatusNotImplemented, helper.NewApiResponse(http.StatusNotImplemented, "Forgot password not implemented yet", nil))
}

// refresh token
func (h *handler) RefreshToken(c *gin.Context) {
	tr := otel.Tracer("authentication-handler")
	ctx, span := tr.Start(c, "RefreshTokenHandler")
	defer span.End()

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

	response := h.service.RefreshToken(ctx, request.RefreshToken)
	c.JSON(response.Code, response)
}

func (h *handler) Me(c *gin.Context) {
	tr := otel.Tracer("authentication-handler")
	ctx, span := tr.Start(c, "MeHandler")
	defer span.End()

	response := h.service.Me(ctx)
	c.JSON(response.Code, response)
}
