package main

import (
	"adventOfCode/utils"
	"testing"
)

func TestFindFirstDigit(t *testing.T) {
	inputs := []utils.TestCase[string, int]{
		{Case: "1abc2", Expected: 1},
		{Case: "pqr3stu8vwx", Expected: 3},
		{Case: "a1b2c3d4e5f", Expected: 1},
		{Case: "treb7uchet", Expected: 7},
	}
	for _, input := range inputs {
		t.Run(input.Case, func(t *testing.T) {
			result, _ := getFirstDigit(input.Case)

			if result != input.Expected {
				t.Errorf("from %+v got %d", input, result)
			}
		})
	}
}

func TestPanicFindFirstDigit(t *testing.T) {
	line := "absajberakb"
	want := NoDigitFound
	if _, err := getFirstDigit(line); err != want {
		t.Errorf("should have panicked with '%v', instead got '%v'", want, err)
	}
}

func TestGetCalibrationValue(t *testing.T) {
	inputs := []utils.TestCase[string, int]{
		{Case: "1abc2", Expected: 12},
		{Case: "pqr3stu8vwx", Expected: 38},
		{Case: "a1b2c3d4e5f", Expected: 15},
		{Case: "treb7uchet", Expected: 77},
	}

	for _, input := range inputs {
		t.Run(input.Case, func(t *testing.T) {
			result, _ := getCalibrationValue(input.Case)
			if result != input.Expected {
				t.Errorf("expected %d got %d", input.Expected, result)
			}
		})
	}
}

func TestReplaceSubstringNumbers(t *testing.T) {
	testCases := []utils.TestCase[string, string]{
		{Case: "two1nine", Expected: "219"},
		{Case: "eightwothree", Expected: "8wo3"},
		{Case: "abcone2threexyz", Expected: "abc123xyz"},
		{Case: "xtwone3four", Expected: "x2ne34"},
		{Case: "4nineeightseven2", Expected: "49872"},
		{Case: "zoneight234", Expected: "z1ight234"},
		{Case: "7pqrstsixteen", Expected: "7pqrst6teen"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			result := replaceSubstringNumbers(testCase.Case)
			if result != testCase.Expected {
				t.Errorf("expected %s got %s", testCase.Expected, result)
			}
		})
	}
}

func TestCalibrationWithNumbersReplacement(t *testing.T) {
	testCases := []utils.TestCase[string, int]{
		{Case: "two1nine", Expected: 29},
		{Case: "eightwothree", Expected: 83},
		{Case: "abcone2threexyz", Expected: 13},
		{Case: "xtwone3four", Expected: 24},
		{Case: "4nineeightseven2", Expected: 42},
		{Case: "zoneight234", Expected: 14},
		{Case: "7pqrstsixteen", Expected: 76},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			caseReplaced := replaceSubstringNumbers(testCase.Case)
			if result, err := getCalibrationValue(caseReplaced); err != nil {
				t.Errorf("got error %s", err)
			} else if result != testCase.Expected {
				t.Errorf("got %d expected %d", result, testCase.Expected)
			}
		})
	}
}