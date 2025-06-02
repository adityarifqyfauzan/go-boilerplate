package i18n

type Translator interface {
	T(messageID string, data map[string]interface{}) string
	FieldName(field string) string
}
