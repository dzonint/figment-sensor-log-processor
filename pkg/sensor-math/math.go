package sensor_math

import (
	"math"
)

func CalculateMeanAndStdDev(values []float64) (mean, sd float64) {
	if values == nil {
		return 0, 0
	}

	var sum, numberOfValues float64
	for _, value := range values {
		sum += value
	}
	if sum == 0 {
		return 0, 0
	}
	if numberOfValues = float64(len(values)); numberOfValues == 1.0 {
		return values[0], 0
	}
	mean = sum / numberOfValues

	// Get sum of squares of differences of each value and mean.
	for _, value := range values {
		sd += math.Pow(value-mean, 2)
	}
	// Get square root of division of the sum with n-1
	sd = math.Sqrt(sd / (numberOfValues - 1))

	// Round to 2 decimals.
	return math.Round(mean*100) / 100, math.Round(sd*100) / 100
}

func AreValuesWithinPercentageOfReferenceValue(referenceValue, percentage float64, values []float64) bool {
	allowedDeviationValue := (percentage / 100) * referenceValue
	for _, value := range values {
		if math.Abs(referenceValue-value) > allowedDeviationValue {
			return false
		}
	}
	return true
}
