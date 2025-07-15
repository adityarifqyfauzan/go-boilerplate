package validator

import (
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/translator"
	goValidator "github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Validator struct {
	localizer *i18n.Localizer
}

// New returns a new Validator instance.
//
// localizer is an instance of github.com/nicksnyder/go-i18n/v2/i18n.Localizer.
// It is used to translate the validation errors.
func New(localizer *i18n.Localizer) *Validator {
	return &Validator{
		localizer: localizer,
	}
}

// Validate checks the provided data against validation rules and returns a map of errors.
// Each key in the map is a field name, and the value is a localized error message.
// The function uses go-playground/validator for validation and i18n for error message translation.
// If no validation errors are found, an empty map is returned.

func (v *Validator) Validate(data any) map[string]any {
	validator := goValidator.New()

	errors := make(map[string]any, 0)
	err := validator.Struct(data)
	if err == nil {
		return errors
	}

	translator := translator.NewTranslator(v.localizer)

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

// FirstError returns the first error message found in the map of errors, or an empty string if the map is empty.
func (v *Validator) FirstError(errors map[string]any) string {
	for _, err := range errors {
		return err.(string)
	}
	return ""
}
