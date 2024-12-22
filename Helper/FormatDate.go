package Helper

import "time"

func FormatDate(date string) (time.Time, error) {
	formateddate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return formateddate, err
	}

	return formateddate, nil
}
