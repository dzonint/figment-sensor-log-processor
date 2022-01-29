package sensors

import (
	sensor_math "figment-sensor-log-processor/pkg/sensor-math"
	"math"
)

const (
	meanDeviationThreshold            = 0.5
	stdDeviationThresholdUltraPrecise = 3
	stdDeviationThresholdVeryPrecise  = 5
	ultraPreciseValue                 = "ultra precise"
	veryPreciseValue                  = "very precise"
	preciseValue                      = "precise"
)

type Thermometer struct {
	temperature float64
}

func (t *Thermometer) DetermineState(values []float64) string {
	mean, sd := sensor_math.CalculateMeanAndStdDev(values)
	if math.Abs(t.temperature-mean) <= meanDeviationThreshold && sd < stdDeviationThresholdUltraPrecise {
		return ultraPreciseValue
	} else if math.Abs(t.temperature-mean) <= meanDeviationThreshold && sd < stdDeviationThresholdVeryPrecise {
		return veryPreciseValue
	}
	return preciseValue
}
