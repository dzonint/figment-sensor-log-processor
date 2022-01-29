package sensor_math

import "testing"

type meanAndStdDevTest struct {
	name     string
	input    []float64
	expected struct {
		mean float64
		sd   float64
	}
}

type areValuesWithinPercentageOfReferenceValueTest struct {
	name  string
	input struct {
		referenceValue float64
		percentage     float64
		values         []float64
	}
	expected bool
}

var calculateMeanAndStdDevTests []meanAndStdDevTest
var areValuesWithinPercentageOfReferenceValueTests []areValuesWithinPercentageOfReferenceValueTest

func init() {
	calculateMeanAndStdDevTests = []meanAndStdDevTest{
		{
			name:  "TestCalculateMeanAndStdDevNilValue",
			input: nil,
			expected: struct {
				mean float64
				sd   float64
			}{mean: 0, sd: 0},
		},
		{
			name:  "TestCalculateMeanAndStdDevZeroValues",
			input: []float64{0.0, 0.0},
			expected: struct {
				mean float64
				sd   float64
			}{mean: 0, sd: 0},
		},
		{
			name:  "TestCalculateMeanAndStdDevSingleValue",
			input: []float64{11.5},
			expected: struct {
				mean float64
				sd   float64
			}{mean: 11.5, sd: 11.5},
		},
		{
			name:  "TestCalculateMeanAndStdDevNegativeNumber",
			input: []float64{-11.5},
			expected: struct {
				mean float64
				sd   float64
			}{mean: -11.5, sd: -11.5},
		},
		{
			name:  "TestCalculateMeanAndStdDevNegativeNumberAndZero",
			input: []float64{-11.5, 0},
			expected: struct {
				mean float64
				sd   float64
			}{mean: -5.75, sd: 8.13},
		},
		{
			name:  "TestCalculateMeanAndStdDevTestInputWithTwoDecimals",
			input: []float64{19.96, 20.22},
			expected: struct {
				mean float64
				sd   float64
			}{mean: 20.09, sd: 0.18},
		},
		{
			name:  "TestCalculateMeanAndStdDevTest1",
			input: []float64{13.1, 33.2},
			expected: struct {
				mean float64
				sd   float64
			}{mean: 23.15, sd: 14.21},
		},
		{
			name:  "TestCalculateMeanAndStdDevNegativeNumberTest2",
			input: []float64{72.4, 76.0, 79.1, 75.6, 71.2, 71.4, 69.2, 65.2, 62.8, 61.4, 64.0, 67.5, 69.4},
			expected: struct {
				mean float64
				sd   float64
			}{mean: 69.63, sd: 5.4},
		},
		{
			name:  "TestCalculateMeanAndStdDevNegativeNumberTest3",
			input: []float64{69.5, 70.1, 71.3, 71.5, 69.8},
			expected: struct {
				mean float64
				sd   float64
			}{mean: 70.44, sd: 0.9},
		},
	}

	areValuesWithinPercentageOfReferenceValueTests = []areValuesWithinPercentageOfReferenceValueTest{
		{
			name: "areValuesWithinPercentageOfReferenceValueTestZeroPercentTrue",
			input: struct {
				referenceValue float64
				percentage     float64
				values         []float64
			}{45.0, 0, []float64{45}},
			expected: true,
		},
		{
			name: "areValuesWithinPercentageOfReferenceValueTestZeroPercentFalse",
			input: struct {
				referenceValue float64
				percentage     float64
				values         []float64
			}{45.0, 0, []float64{44.9}},
			expected: false,
		},
		{
			name: "areValuesWithinPercentageOfReferenceValueTestTrue",
			input: struct {
				referenceValue float64
				percentage     float64
				values         []float64
			}{45.0, 1, []float64{45.2, 45.3, 45.1}},
			expected: true,
		},
		{
			name: "areValuesWithinPercentageOfReferenceValueTestFalse",
			input: struct {
				referenceValue float64
				percentage     float64
				values         []float64
			}{45.0, 1, []float64{44.4, 43.9, 44.9, 43.8, 42.1}},
			expected: false,
		},
	}
}

func TestCalculateMeanAndStdDev(t *testing.T) {
	for _, test := range calculateMeanAndStdDevTests {
		t.Run(test.name, func(t *testing.T) {
			if mean, sd := CalculateMeanAndStdDev(test.input); test.expected.mean != mean || test.expected.sd != sd {
				t.Errorf("Expected values were not equal.\nexpected=[%v, %v]\nactual=[%v, %v]",
					test.expected.mean,
					test.expected.sd,
					mean,
					sd)
			}
		})
	}
}

func TestAreValuesWithinPercentageOfReferenceValue(t *testing.T) {
	for _, test := range areValuesWithinPercentageOfReferenceValueTests {
		t.Run(test.name, func(t *testing.T) {
			if output := AreValuesWithinPercentageOfReferenceValue(test.input.referenceValue, test.input.percentage, test.input.values); output != test.expected {
				t.Errorf("Expected values were not equal.\nexpected=%v,\nactual=%v", test.expected, output)
			}
		})
	}
}
