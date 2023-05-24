package utils

import "strings"

// ловеркейс, убрать повторяющиеся подряд символы, пробелы по краям
func NormalizeString(rawString string) string {
	trimmedAndLowered := strings.Trim(strings.ToLower(rawString), " ")

	var resultBuilder strings.Builder
	var previousRune rune = -1
	for _, symbol := range trimmedAndLowered {
		if symbol != previousRune {
			resultBuilder.WriteRune(symbol)
			previousRune = symbol
		}
	}

	return resultBuilder.String()
}
