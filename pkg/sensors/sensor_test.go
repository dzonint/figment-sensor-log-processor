package sensors

import (
	"reflect"
	"testing"
)

type sensorStrategyTest struct {
	name     string
	input    string
	expected struct {
		sensor DynamicSensor
		err    error
	}
}

type determineSensorTypeTest struct {
	name     string
	input    string
	expected struct {
		sensor DynamicSensor
		err    error
	}
}

type isValidSensorTypeTest struct {
	name     string
	input    string
	expected bool
}

var sensorStrategyTests []sensorStrategyTest
var determineSensorTypeTests []determineSensorTypeTest
var isValidSensorTypeTests []isValidSensorTypeTest

func init() {
	sensorStrategyTests = []sensorStrategyTest{
		{
			name:  "TestNewSensorThermometer",
			input: thermometerKeyword,
			expected: struct {
				sensor DynamicSensor
				err    error
			}{&thermometer{}, nil},
		},
		{
			name:  "TestNewSensorHumidity",
			input: humidityKeyword,
			expected: struct {
				sensor DynamicSensor
				err    error
			}{&humidity{}, nil},
		},
		{
			name:  "TestNewSensorUnknownInput",
			input: "undefined",
			expected: struct {
				sensor DynamicSensor
				err    error
			}{nil, invalidSensorTypeError},
		},
	}

	determineSensorTypeTests = []determineSensorTypeTest{
		{
			name:  "TestDetermineSensorTypeThermometer",
			input: thermometerKeyword,
			expected: struct {
				sensor DynamicSensor
				err    error
			}{&thermometer{}, nil},
		},
		{
			name:  "TestDetermineSensorTypeHumidity",
			input: humidityKeyword,
			expected: struct {
				sensor DynamicSensor
				err    error
			}{&humidity{}, nil},
		},
	}

	isValidSensorTypeTests = []isValidSensorTypeTest{
		{
			name:     "isValidSensorTypeTestThermometer",
			input:    thermometerKeyword,
			expected: true,
		},
		{
			name:     "isValidSensorTypeTestHumidity",
			input:    humidityKeyword,
			expected: true,
		},
		{
			name:     "isValidSensorTypeTestUndefined",
			input:    "undefined",
			expected: false,
		},
	}
}

func TestNewSensorStrategy(t *testing.T) {
	for _, test := range sensorStrategyTests {
		t.Run(test.name, func(t *testing.T) {
			if sensor, err := NewSensor(test.input, "", 0, 0); getTypeValueString(test.expected.sensor) != getTypeValueString(sensor.ds) || err != test.expected.err {
				t.Errorf("Expected values were not equal.\nexpected=[%v, %v]\nactual=[%v, %v]", getTypeValueString(test.expected), test.expected.err, getTypeValueString(sensor.ds), err)
			}
		})
	}
}

func TestDetermineSensor(t *testing.T) {
	for _, test := range determineSensorTypeTests {
		t.Run(test.name, func(t *testing.T) {
			if sensor, err := determineSensorType(test.input, 0, 0); getTypeValueString(test.expected.sensor) != getTypeValueString(sensor) || err != test.expected.err {
				t.Errorf("Expected values were not equal.\nexpected=[%v, %v]\nactual=[%v, %v]", getTypeValueString(test.expected), test.expected.err, getTypeValueString(sensor), err)
			}
		})
	}
}

func TestIsValidSensorType(t *testing.T) {
	for _, test := range isValidSensorTypeTests {
		t.Run(test.name, func(t *testing.T) {
			if output := isValidSensorType(test.input); output != test.expected {
				t.Errorf("Expected values were not equal.\nexpected=%v\nactual=%v", test.expected, output)
			}
		})
	}
}

func getTypeValueString(i interface{}) string {
	return reflect.ValueOf(i).String()
}
