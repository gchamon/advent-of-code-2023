package main

import (
	"adventOfCode/utils"
	"fmt"
	"testing"
)

func TestGameStringParse(t *testing.T) {
	testCases := []utils.TestCase[string, CubeGame]{
		{
			Case: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: CubeGame{Id: 1, CubesSubsets: []CubesSubset{
				{Red: 4, Blue: 3, Green: 0},
				{Red: 1, Blue: 6, Green: 2},
				{Red: 0, Blue: 0, Green: 2},
			}},
		},
		{
			Case: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: CubeGame{Id: 2, CubesSubsets: []CubesSubset{
				{Red: 0, Blue: 1, Green: 2},
				{Red: 1, Blue: 4, Green: 3},
				{Red: 0, Blue: 1, Green: 1},
			}},
		},
		{
			Case: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: CubeGame{Id: 3, CubesSubsets: []CubesSubset{
				{Red: 20, Blue: 6, Green: 8},
				{Red: 4, Blue: 5, Green: 13},
				{Red: 1, Blue: 0, Green: 5},
			}},
		},
		{
			Case: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: CubeGame{Id: 4, CubesSubsets: []CubesSubset{
				{Red: 3, Blue: 6, Green: 1},
				{Red: 6, Blue: 0, Green: 3},
				{Red: 14, Blue: 15, Green: 3},
			}},
		},
		{
			Case: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: CubeGame{Id: 5, CubesSubsets: []CubesSubset{
				{Red: 6, Blue: 1, Green: 3},
				{Red: 1, Blue: 2, Green: 2},
			}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			utils.AssertDeepEqual(t, parseGame(testCase.Case), testCase.Expected)
		})
	}
}

func TestGamePossible(t *testing.T) {
	testCases := []utils.TestCase[string, bool]{
		{Case: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", Expected: true},
		{Case: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", Expected: true},
		{Case: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", Expected: false},
		{Case: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", Expected: false},
		{Case: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", Expected: true},
	}

	sumPossibleIds := 0
	expectedSumPossibleIds := 8
	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			game := parseGame(testCase.Case)
			result := game.IsPossible()
			utils.AssertBool(t, result, testCase.Expected)
			if result == true {
				sumPossibleIds += game.Id
			}
		})
	}
	utils.AssertInt(t, expectedSumPossibleIds, sumPossibleIds)
}

func TestMinimumGamePossible(t *testing.T) {
	testCases := []utils.TestCase[string, CubesSubset]{
		{
			Case: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: CubesSubset{
				Red: 4, Green: 2, Blue: 6,
			},
		},
		{
			Case: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: CubesSubset{
				Red: 1, Green: 3, Blue: 4,
			},
		},
		{
			Case: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: CubesSubset{
				Red: 20, Green: 13, Blue: 6,
			},
		},
		{
			Case: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: CubesSubset{
				Red: 14, Green: 3, Blue: 15,
			},
		},
		{
			Case: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: CubesSubset{
				Red: 6, Green: 3, Blue: 2,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			utils.AssertDeepEqual(t, parseGame(testCase.Case).GetMinimumSubset(), testCase.Expected)
		})
	}
}

func TestCalulatePower(t *testing.T) {
	testCases := []utils.TestCase[CubesSubset, int]{
		{
			Case: CubesSubset{
				Red: 4, Green: 2, Blue: 6,
			},
			Expected: 48,
		},
		{
			Case: CubesSubset{
				Red: 1, Green: 3, Blue: 4,
			},
			Expected: 12,
		},
		{
			Case: CubesSubset{
				Red: 20, Green: 13, Blue: 6,
			},
			Expected: 1560,
		},
		{
			Case: CubesSubset{
				Red: 14, Green: 3, Blue: 15,
			},
			Expected: 630,
		},
		{
			Case: CubesSubset{
				Red: 6, Green: 3, Blue: 2,
			},
			Expected: 36,
		},
	}
	expectedTotalPowerSum := 2286
	totalPowerSum := 0
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%+v", testCase.Case), func(t *testing.T) {
			result := testCase.Case.CalculatePower()
			utils.AssertInt(t, result, testCase.Expected)
			totalPowerSum += result
		})
	}
	utils.AssertInt(t, totalPowerSum, expectedTotalPowerSum)
}
