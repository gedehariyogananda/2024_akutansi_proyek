package Utils

import (
	"strings"
	"time"
)

func FormatUsernameClient(input string) string {
	trimmedInput := strings.TrimSpace(input)
	withUnderscore := strings.ReplaceAll(trimmedInput, " ", "_")
	formattedString := strings.ToLower(withUnderscore)
	return formattedString
}

func ParseDateStringToDate(input string, format *string) time.Time {
	if format == nil {
		defaultFormat := "2006-01-02"
		format = &defaultFormat
	}

	parseDate, _ := time.Parse(*format, input)

	return parseDate
}
