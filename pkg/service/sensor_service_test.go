package service

import (
	"go.uber.org/zap"
	"os"
	"testing"
)

type determineTemperatureAndHumidityTest struct {
	name     string
	input    string
	expected struct {
		temp     float64
		humidity float64
		err      error
	}
}

type processLogTest struct {
	name     string
	input    string
	expected error
}

type isFileEmptyTest struct {
	name     string
	input    string
	expected bool
}

var ss SensorService
var determineTemperatureAndHumidityTests []determineTemperatureAndHumidityTest
var processLogTests []processLogTest
var isFileEmptyTests []isFileEmptyTest

func init() {
	ss = NewSensorService(zap.L())
	determineTemperatureAndHumidityTests = []determineTemperatureAndHumidityTest{
		{
			name:  "TestDetermineTemperatureAndHumidityNotEnoughArguments",
			input: "reference 70.0",
			expected: struct {
				temp     float64
				humidity float64
				err      error
			}{0, 0, notEnoughArgumentsError},
		},
		{
			name:  "TestDetermineTemperatureAndHumidityMissingReferenceKeyword",
			input: "not-reference 70.0 45.0",
			expected: struct {
				temp     float64
				humidity float64
				err      error
			}{0, 0, missingReferenceKeywordError},
		},
		{
			name:  "TestDetermineTemperatureAndHumidityFailedToParseTemperature",
			input: "reference temp 45.0",
			expected: struct {
				temp     float64
				humidity float64
				err      error
			}{0, 0, failedToParseTemperatureError},
		},
		{
			name:  "TestDetermineTemperatureAndHumidityFailedToParseHumidity",
			input: "reference 70.0 humidity",
			expected: struct {
				temp     float64
				humidity float64
				err      error
			}{0, 0, failedToParseHumidityError},
		},
		{
			name:  "TestDetermineTemperatureAndHumidityNoError",
			input: "reference 70.0 45.0",
			expected: struct {
				temp     float64
				humidity float64
				err      error
			}{70.0, 45.0, nil},
		},
	}

	processLogTests = []processLogTest{
		{
			name:     "TestProcessLogInvalidFilepath",
			input:    "testdata/.log-does-not-exist",
			expected: failedOpeningLogError,
		},
		{
			name:     "TestProcessLogInvalidReference",
			input:    "testdata/.log-with-invalid-reference",
			expected: determineTemperatureAndHumidityError,
		},
		{
			name:     "TestProcessLogEmptyFile",
			input:    "testdata/.log-empty",
			expected: nil,
		},
		{
			name:     "TestProcessLogInvalidValue",
			input:    "testdata/.log-with-invalid-value",
			expected: nil, // This test should still pass.
		},
		{
			name:     "TestProcessLogNoError",
			input:    "testdata/.log",
			expected: nil,
		},
	}

	isFileEmptyTests = []isFileEmptyTest{
		{
			name:     "isFileEmptyTrue",
			input:    "testdata/.log-empty",
			expected: true,
		},
		{
			name:     "isFileEmptyFalse",
			input:    "testdata/.log",
			expected: false,
		},
	}
}

func TestDetermineTemperatureAndHumidity(t *testing.T) {
	for _, test := range determineTemperatureAndHumidityTests {
		t.Run(test.name, func(t *testing.T) {
			if temp, humidity, err := ss.determineTemperatureAndHumidity(test.input); test.expected.temp != temp || test.expected.humidity != humidity || test.expected.err != err {
				t.Errorf("Expected values were not equal.\nexpected=[%v, %v, %v]\nactual=[%v, %v, %v]",
					test.expected.temp,
					test.expected.humidity,
					test.expected.err,
					temp,
					humidity,
					err)
			}
		})
	}
}

func TestProcessLog(t *testing.T) {
	for _, test := range processLogTests {
		t.Run(test.name, func(t *testing.T) {
			if err := ss.ProcessLog(test.input); test.expected != err {
				t.Errorf("Expected values were not equal.\nexpected=%v\nactual=%v", test.expected, err)
			}
		})
	}
}

func TestIsFileEmpty(t *testing.T) {
	for _, test := range isFileEmptyTests {
		t.Run(test.name, func(t *testing.T) {
			file, _ := os.Open(test.input)
			if output := ss.isFileEmpty(file); test.expected != output {
				t.Errorf("Expected values were not equal.\nexpected=%v\nactual=%v", test.expected, output)
			}
		})
	}
}
