package i18n

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func Init(path string) *i18n.Bundle {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	files, _ := filepath.Glob(filepath.Join(path, "*.json"))
	for _, file := range files {
		_, err := bundle.LoadMessageFile(file)
		if err != nil {
			panic(fmt.Sprintf("failed to load locale file %s: %v", file, err))
		}
	}

	return bundle
}

func GetBundle() *i18n.Bundle {
	return bundle
}

func NewLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, lang)
}
