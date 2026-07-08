package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

var defaultLang language.Tag

func Init(localesPath string, defaultLanguage language.Tag) error {
	defaultLang = defaultLanguage

	bundle = i18n.NewBundle(defaultLanguage)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	files, err := os.ReadDir(localesPath)
	if err != nil {
		return fmt.Errorf("i18n: could not read locales directory (%s): %w", localesPath, err)
	}

	loaded := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.ToLower(filepath.Ext(file.Name())) != ".json" {
			continue
		}

		fullPath := filepath.Join(localesPath, file.Name())
		if _, err := bundle.LoadMessageFile(fullPath); err != nil {
			return fmt.Errorf("i18n: could not load message file (%s): %w", fullPath, err)
		}
		loaded++
	}

	if loaded == 0 {
		return fmt.Errorf("i18n: no .json language files found in directory '%s'", localesPath)
	}

	return nil
}

func Translate(lang string, messageID string, fallback string) string {
	if bundle == nil {
		return fallback
	}

	localizer := i18n.NewLocalizer(bundle, lang, defaultLang.String())

	translated, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})

	if err != nil || translated == "" {
		return fallback
	}

	return translated
}
