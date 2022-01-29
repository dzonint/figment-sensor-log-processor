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

func NewSensor(sensorType string, name string, temperature float64, humidity float64) *Sensor {
	var ds DynamicSensor = &Thermometer{temperature}
	if sensorType == humidityKeyword {
		ds = &Humidity{humidity}
	}

	return &Sensor{
		name: name,
		ds:   ds,
	}
}

func (s *Sensor) PrintState() {
	fmt.Printf(statusPlaceholder, s.name, s.ds.DetermineState(s.Values))
}
