package Utils

import (
	"strconv"
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

type Seperation struct {
	Year  *int
	Month *int
	Day   *int
}

func SeperateDate(date string) (options *Seperation) {
	dateParse := strings.Split(date, "-")
	return &Seperation{
		Year:  func() *int { year, _ := strconv.Atoi(dateParse[0]); return &year }(),
		Month: func() *int { month, _ := strconv.Atoi(dateParse[1]); return &month }(),
		Day:   func() *int { day, _ := strconv.Atoi(dateParse[2]); return &day }(),
	}
}
