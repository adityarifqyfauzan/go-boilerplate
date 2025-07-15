package translator

import (
	"testing"
)

func TestI18n(t *testing.T) {
	bundle := Init("../../locales")
	if bundle == nil {
		panic("failed to init i18n")
	}

	if GetBundle() == nil {
		t.Errorf("expected %s, got %s", "bundle is not nil", "bundle is nil")
	}

	localizer := NewLocalizer("en")
	translator := NewTranslator(localizer)
	msg := translator.T("hello", map[string]any{
		"Name": "Aditya",
	})

	if msg != "Hello Aditya" {
		t.Errorf("expected %s, got %s", "Hello Aditya", msg)
	}

	localizerID := NewLocalizer("id")

	translatorID := NewTranslator(localizerID)
	msgId := translatorID.T("hello", map[string]any{
		"Name": "Aditya",
	})

	if msgId != "Halo Aditya" {
		t.Errorf("expected %s, got %s", "Halo Aditya", msgId)
	}

	localizerJA := NewLocalizer("ja")

	translatorJA := NewTranslator(localizerJA)
	msgJA := translatorJA.T("hello", map[string]any{
		"Name": "Aditya",
	})

	if msgJA != "こんにちは、Aditya" {
		t.Errorf("expected %s, got %s", "こんにちは、Aditya", msgJA)
	}
}

func TestPlural(t *testing.T) {
	bundle := Init("../../locales")
	if bundle == nil {
		panic("failed to init i18n")
	}

	if GetBundle() == nil {
		t.Errorf("expected %s, got %s", "bundle is not nil", "bundle is nil")
	}

	localizer := NewLocalizer("en")

	translator := NewTranslator(localizer)
	msg := translator.T("plural.message", map[string]any{
		"Count": 1,
	})

	if msg != "There is one message" {
		t.Errorf("expected %s, got %s", "There is one message", msg)
	}

	msg = translator.T("plural.message", map[string]any{
		"Count": 2,
	})

	if msg != "There are 2 messages" {
		t.Errorf("expected %s, got %s", "There are 2 messages", msg)
	}

	// except indonesia and jp
	// it's not have plural term
	// so it will return the same message
	localizerID := NewLocalizer("id")

	translatorID := NewTranslator(localizerID)
	msgId := translatorID.T("plural.message", map[string]any{
		"Count": 1,
	})

	if msgId != "Ada 1 pesan" {
		t.Errorf("expected %s, got %s", "Ada 1 pesan", msgId)
	}

	msgId = translatorID.T("plural.message", map[string]any{
		"Count": 2,
	})

	if msgId != "Ada 2 pesan" {
		t.Errorf("expected %s, got %s", "Ada 2 pesan", msgId)
	}

	localizeJA := NewLocalizer("ja")

	translatorJA := NewTranslator(localizeJA)
	msgJA := translatorJA.T("plural.message", map[string]any{
		"Count": 1,
	})

	if msgJA != "1件のメッセージ" {
		t.Errorf("expected %s, got %s", "1件のメッセージ", msgJA)
	}

	msgJA = translatorJA.T("plural.message", map[string]any{
		"Count": 2,
	})

	if msgJA != "2件のメッセージ" {
		t.Errorf("expected %s, got %s", "2件のメッセージ", msgJA)
	}
}
