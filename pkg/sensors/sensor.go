package sensors

import "fmt"

const (
	thermometerKeyword = "thermometer"
	humidityKeyword    = "humidity"
	statusPlaceholder  = "%v: %v\n"
)

type DynamicSensor interface {
	DetermineState(values []float64) string
}

type Sensor struct {
	ds     DynamicSensor
	name   string
	Values []float64
}

func NewSensor(sensorType string, name string, temperature float64, humidityPercentage float64) *Sensor {
	var ds DynamicSensor = &thermometer{temperature}
	if sensorType == humidityKeyword {
		ds = &humidity{humidityPercentage}
	}

	return &Sensor{
		name: name,
		ds:   ds,
	}
}

func (s *Sensor) PrintState() {
	fmt.Printf(statusPlaceholder, s.name, s.ds.DetermineState(s.Values))
}
