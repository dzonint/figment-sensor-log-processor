package sensors

import (
	sensor_math "figment-sensor-log-processor/pkg/sensor-math"
	"math"
)

type sensor struct {
	sensorType  string
	name        string
	readings    []reading
	values      []float64
	temperature float64
	humidity    float64
}

type reading struct {
	time  string
	name  string
	value float64
}

type Thermometer struct {
	sensor
}

func (t *Thermometer) DeterminePrecision() string {
	mean, sd := sensor_math.CalculateMeanAndStdDev(t.values)

	if math.Abs(t.temperature-mean) <= 0.5 && sd < 3 {
		return "ultra precise"
	} else if math.Abs(t.temperature-mean) <= 0.5 && sd < 5 {
		return "very precise"
	}
	return "precise"
}

type Humidity struct {
	sensor
}

func (h *Humidity) DetermineStatus() string {
	if output := sensor_math.AreValuesWithinPercentageOfReferenceValue(h.humidity, 1, h.values); output {
		return "OK"
	}
	return "discard"
}
