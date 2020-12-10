package day10

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day10/"

func TestProblem1(t *testing.T) {
	input := utils.ParseNumbers(fmt.Sprint(testPath, "day10_test.txt"))
	result := problem1(input)
	if result != 35 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 35)
	}

	input = utils.ParseNumbers(fmt.Sprint(testPath, "day10_test2.txt"))
	result = problem1(input)
	if result != 220 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 220)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseNumbers(fmt.Sprint(testPath, "day10_test.txt"))
	result := problem2(input)
	if result != 8 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 8)
	}

	input = utils.ParseNumbers(fmt.Sprint(testPath, "day10_test2.txt"))
	result = problem2(input)
	if result != 19208 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 19208)
	}
}
