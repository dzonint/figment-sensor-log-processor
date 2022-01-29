package sensors

import "testing"

type humidityTest struct {
	name     string
	humidity humidity
	input    []float64
	expected string
}

var determineStateHumidityTests []humidityTest

func init() {
	determineStateHumidityTests = []humidityTest{
		{
			name:     "TestDetermineStateHumidityOK",
			humidity: humidity{45.0},
			input:    []float64{45.2, 45.3, 45.1},
			expected: okValue,
		},
		{
			name:     "TestDetermineStateHumidityDiscard",
			humidity: humidity{45.0},
			input:    []float64{44.4, 43.9, 44.9, 43.8, 42.1},
			expected: discardValue,
		},
	}
}

func TestDetermineStateHumidity(t *testing.T) {
	for _, test := range determineStateHumidityTests {
		t.Run(test.name, func(t *testing.T) {
			if output := test.humidity.DetermineState(test.input); test.expected != output {
				t.Errorf("Expected values were not equal.\nexpected=%v\nactual=%v", test.expected, output)
			}
		})
	}
}
