package sensors

import (
	sensor_math "figment-sensor-log-processor/pkg/sensor-math"
)

const (
	allowedDeviationPercentage = 1
	okValue                    = "OK"
	discardValue               = "discard"
)

type humidity struct {
	humidity float64
}

func (h *humidity) DetermineState(values []float64) string {
	if output := sensor_math.AreValuesWithinPercentageOfReferenceValue(h.humidity, allowedDeviationPercentage, values); output {
		return okValue
	}
	return discardValue
}
