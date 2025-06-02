package validator

import (
	ii18n "github.com/adityarifqyfauzan/go-boilerplate/pkg/i18n"
	goValidator "github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Validator struct {
	localizer *i18n.Localizer
}

func New(localizer *i18n.Localizer) *Validator {
	return &Validator{
		localizer: localizer,
	}
}

func (v *Validator) Validate(data any) map[string]any {
	validator := goValidator.New()

	errors := make(map[string]any, 0)
	err := validator.Struct(data)
	if err == nil {
		return errors
	}

	translator := ii18n.NewTranslator(v.localizer)

	for _, err := range err.(goValidator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		messageParam := map[string]any{
			"Field": translator.FieldName(field),
		}

		if tag == "min" || tag == "max" || tag == "gte" || tag == "lte" {
			messageParam["Param"] = err.Param()
		}

		errors[field] = translator.T("validation."+tag, messageParam)
	}

	return errors
}
