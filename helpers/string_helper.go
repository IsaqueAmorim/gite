package stringHelper

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Normalize(str string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, str)

	return result

}

func RemoveSpecialCharacters(str string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")

}

func AddSeparators(str string) string {
	return strings.ReplaceAll(str, " ", "-")
}

func RemoveSequeneOfSeparators(str string) string {

	return regexp.MustCompile(`(--)+`).ReplaceAllString(str, "-")
}
