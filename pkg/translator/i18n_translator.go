package translator

import (
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const (
	LOCALIZER = "localizer"
)

type GoI18nTranslator struct {
	localizer *i18n.Localizer
}

// NewTranslator returns a new Translator instance
func NewTranslator(localizer *i18n.Localizer) Translator {
	return &GoI18nTranslator{
		localizer: localizer,
	}
}

func (t *GoI18nTranslator) T(messageID string, data map[string]interface{}) string {
	msg, err := t.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
		PluralCount:  data["Count"],
	})
	if err != nil {
		return messageID
	}
	return msg
}

func (t *GoI18nTranslator) FieldName(field string) string {
	// the field format is CamelCase, so I wanna convert it to snake_case
	var snakeCaseField strings.Builder
	for i, r := range field {
		if i > 0 && r >= 'A' && r <= 'Z' {
			snakeCaseField.WriteByte('_')
		}
		snakeCaseField.WriteByte(byte(strings.ToLower(string(r))[0]))
	}
	msg, err := t.localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "fields." + snakeCaseField.String(),
	})
	if err != nil {
		return field
	}
	return msg
}
