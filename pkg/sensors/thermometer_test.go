package sensors

import "testing"

type thermometerTest struct {
	name        string
	thermometer thermometer
	input       []float64
	expected    string
}

var determineStateThermometerTests []thermometerTest

func init() {
	determineStateThermometerTests = []thermometerTest{
		{
			name:        "TestDetermineStateThermometerPrecise",
			thermometer: thermometer{70},
			input:       []float64{72.4, 76.0, 79.1, 75.6, 71.2, 71.4, 69.2, 65.2, 62.8, 61.4, 64.0, 67.5, 69.4},
			expected:    preciseValue,
		},
		{
			name:        "TestDetermineStateThermometerVeryPrecise",
			thermometer: thermometer{70},
			input:       []float64{68.5, 71.3, 73.5, 65.8},
			expected:    veryPreciseValue,
		},
		{
			name:        "TestDetermineStateThermometerUltraPrecise",
			thermometer: thermometer{70},
			input:       []float64{69.5, 70.1, 71.3, 71.5, 69.8},
			expected:    ultraPreciseValue,
		},
	}
}

func TestDetermineStateThermometer(t *testing.T) {
	for _, test := range determineStateThermometerTests {
		t.Run(test.name, func(t *testing.T) {
			if output := test.thermometer.DetermineState(test.input); test.expected != output {
				t.Errorf("Expected values were not equal.\nexpected=%v\nactual=%v", test.expected, output)
			}
		})
	}
}
