package utils

import (
	"fmt"
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
