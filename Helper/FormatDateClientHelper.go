package Helper

import "time"

func FormatDateClient(date string) string {
	if date == "" {
		return time.Now().Format("2006-01-02 15:04:05")
	}

	parseDate, _ := time.Parse("2006-01-02", date)

	return parseDate.Format("2006-01-02")
}
