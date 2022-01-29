package sensors

import (
	"reflect"
	"testing"
)

type sensorTest struct {
	name     string
	input    string
	expected DynamicSensor
}

var sensorTests []sensorTest

func init() {
	sensorTests = []sensorTest{
		{
			name:     "TestNewSensorThermometer",
			input:    thermometerKeyword,
			expected: &thermometer{},
		},
		{
			name:     "TestNewSensorHumidity",
			input:    humidityKeyword,
			expected: &humidity{},
		},
		{
			name:     "TestNewSensorUnknownInput",
			input:    "undefined", // Should default to thermometer.
			expected: &thermometer{},
		},
	}
}

func TestNewSensorStrategy(t *testing.T) {
	for _, test := range sensorTests {
		t.Run(test.name, func(t *testing.T) {
			if sensor := NewSensor(test.input, "", 0, 0); getTypeValueString(test.expected) != getTypeValueString(sensor.ds) {
				t.Errorf("Expected values were not equal.\nexpected=%v\nactual=%v", getTypeValueString(test.expected), getTypeValueString(sensor.ds))
			}
		})
	}
}

func getTypeValueString(i interface{}) string {
	return reflect.ValueOf(i).String()
}
