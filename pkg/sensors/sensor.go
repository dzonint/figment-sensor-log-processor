package sensors

import (
	"errors"
	"fmt"
)

const (
	thermometerKeyword = "thermometer"
	humidityKeyword    = "humidity"
	statusPlaceholder  = "%v: %v\n"
)

var allowedSensorTypes = []string{thermometerKeyword, humidityKeyword}
var invalidSensorTypeError = errors.New("sensor type not found among supported sensors")
var undefinedSensorTypeError = errors.New("sensor type is within supported sensors but not implemented")

type DynamicSensor interface {
	DetermineState(values []float64) string
}

type Sensor struct {
	ds     DynamicSensor
	name   string
	Values []float64
}

func NewSensor(sensorType string, name string, temperature float64, humidityPercentage float64) (sensor *Sensor, err error) {
	var ds DynamicSensor
	if ds, err = determineSensorType(sensorType, temperature, humidityPercentage); err != nil {
		return &Sensor{}, err
	}

	return &Sensor{
		name: name,
		ds:   ds,
	}, nil
}

func determineSensorType(sensorType string, temperature float64, humidityPercentage float64) (DynamicSensor, error) {
	if !isValidSensorType(sensorType) {
		return nil, invalidSensorTypeError
	}

	switch sensorType {
	case thermometerKeyword:
		return &thermometer{temperature}, nil
	case humidityKeyword:
		return &humidity{humidityPercentage}, nil
	default: // Should never reach this part (undefined sensor will get filtered out by isValidSensorType method).
		return nil, undefinedSensorTypeError
	}
}

func isValidSensorType(sensorType string) bool {
	for _, allowedSensorType := range allowedSensorTypes {
		if sensorType == allowedSensorType {
			return true
		}
	}
	return false
}

func (s *Sensor) PrintState() {
	fmt.Printf(statusPlaceholder, s.name, s.ds.DetermineState(s.Values))
}
