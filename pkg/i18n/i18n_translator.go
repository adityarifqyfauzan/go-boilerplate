package i18n

import (
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
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
	msg, err := t.localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "fields." + strings.ToLower(field),
	})
	if err != nil {
		return field
	}
	return msg
}
