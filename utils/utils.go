package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Johnsoct/dicthesaurus/repository"
)

// TEXT TRANSFORMATIONS
// TEXT TRANSFORMATIONS
// TEXT TRANSFORMATIONS

func BoldText(text string) string {
	return "\033[1m" + text + "\033[0m"
}

func UnderlineText(text string) string {
	return "\033[4m" + text + "\033[0m"
}

// Must be applied before any other console code because of strings.toUpper
func UppercaseText(text string) string {
	return "\033[1;4m" + strings.ToUpper(text) + "\033[0m"
}

// FORMAT STRINGS
// FORMAT STRINGS
// FORMAT STRINGS

func FormatHeader(h string) string {
	coloredText := "\033[47;30m" + UppercaseText(h) + "\033[0m"
	return fmt.Sprintf("\n\n\n\t%s\n", coloredText)
}

func FormatRow(r string) string {
	return fmt.Sprintf("\n\t%s", r)
}

func FormatTableRow(s []string) string {
	return strings.Repeat("%-15s", len(s))
}

func FormatValueBetweenTokens(text string) string {
	type Replace struct {
		submatch  string
		replaceFn func(text string) string
	}
	replaces := map[string]Replace{
		`\{a_link\|(\w+)\}`: {
			replaceFn: func(text string) string {
				return UnderlineText(text)
			},
			submatch: `\|(\w+)`,
		},
		`\{sx\|(\w+)\|*\}`: {
			replaceFn: func(text string) string {
				uppercased := UppercaseText(text)
				underlined := UnderlineText(uppercased)
				return underlined
			},
			submatch: `\|(\w+)`,
		},
	}

	replaced := text

	for regex, replace := range replaces {
		re := regexp.MustCompile(regex)
		subre := regexp.MustCompile(replace.submatch)
		submatches := 0

		replaced = re.ReplaceAllStringFunc(replaced, func(substring string) string {
			submatch := replace.replaceFn(subre.FindAllStringSubmatch(substring, -1)[0][1])

			if submatches == 0 {
				submatch = ": " + submatch
			}

			submatches++

			return submatch
		})
	}

	return replaced
}

func StripPrefixTokens(text string) string {
	stripped := text
	prefixes := []string{
		`{bc}`,
		`{sx|`,
	}

	for _, v := range prefixes {
		re := regexp.MustCompile(v)
		stripped = re.ReplaceAllString(stripped, "")
	}

	return stripped
}

func StripSuffixTokens(text string) string {
	stripped := text
	suffixes := []string{
		`\|*}`,
		`}`,
		`\|+[0-9]+[a-z]*}`,
	}

	for _, v := range suffixes {
		re := regexp.MustCompile(v)
		stripped = re.ReplaceAllString(stripped, "")
	}

	return stripped
}

func StripTokens(text string) string {
	stripped := text

	prefixStripped := StripPrefixTokens(stripped)
	suffixStripped := StripSuffixTokens(prefixStripped)

	return suffixStripped
}

// CONVERSIONS
// CONVERSIONS
// CONVERSIONS

func ConvertSliceBuiltInTypeToSliceAny[T repository.BuiltIn](s []T) []any {
	values := make([]any, len(s))

	for i, v := range s {
		values[i] = v
	}

	return values
}
