package Helper

func CalculateComparisonYesterday(today float64, yesterday float64) float64 {
	if yesterday == 0 {
		if today == 0 {
			return 0
		} else {
			return 100
		}
	}

	return ((today - yesterday) / yesterday) * 100
}
