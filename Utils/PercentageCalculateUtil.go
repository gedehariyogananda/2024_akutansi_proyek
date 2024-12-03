package Utils

import "math"

func CalculatePercentageInit(prevValue float64, currValue float64) float64 {
	if prevValue == 0 {
		return 0
	}

	percentage := ((currValue - prevValue) * 100) / prevValue

	return math.Round(percentage*100) / 100 // keeped 2 digits percentage
}
