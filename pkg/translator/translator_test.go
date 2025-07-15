package translator

import "testing"

func TestTranslator(t *testing.T) {
	bundle := Init("../../locales")
	if bundle == nil {
		panic("failed to init i18n")
	}

	localizer := NewLocalizer("en")

	translator := NewTranslator(localizer)
	msg := translator.T("hello", map[string]any{
		"Name": "Aditya",
	})

	if msg != "Hello Aditya" {
		t.Errorf("expected %s, got %s", "Hello Aditya", msg)
	}
}

func TestFieldName(t *testing.T) {
	bundle := Init("../../locales")
	if bundle == nil {
		panic("failed to init i18n")
	}

	localizer := NewLocalizer("en")

	translator := NewTranslator(localizer)
	msg := translator.FieldName("name")

	if msg != "Name" {
		t.Errorf("expected %s, got %s", "Name", msg)
	}
}
