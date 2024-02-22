package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToTitleCase(title string) string {
	return cases.Title(language.French, cases.Compact).String(title)
}
