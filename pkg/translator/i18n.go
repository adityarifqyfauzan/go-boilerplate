package translator

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

// Init initializes the translation bundle by loading locale files from the given path.
//
// The locale files must be in JSON format and be named after the language they
// contain, e.g. active.en.json, active.id.json, etc.
//
// The default language is English. If a translation is not available in the
// current language, the English translation will be used as a fallback.
//
// This function panics if it encounters an error while loading the locale files.
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
