package sensors

import (
	sensor_math "figment-sensor-log-processor/pkg/sensor-math"
)

const (
	okValue      = "OK"
	discardValue = "discard"
)

type Humidity struct {
	humidity float64
}

func (h *Humidity) DetermineState(values []float64) string {
	if output := sensor_math.AreValuesWithinPercentageOfReferenceValue(h.humidity, 1, values); output {
		return okValue
	}
	return discardValue
}
