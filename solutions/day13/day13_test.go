package day13

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day13/"

func TestProblem1(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day13_test.txt"))
	result := problem1(input)
	if result != 295 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 295)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day13_test2.txt"))
	result := problem2(input)
	if result != 3417 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 3417)
	}

	input = utils.ParseStrings(fmt.Sprint(testPath, "day13_test3.txt"))
	result = problem2(input)
	if result != 754018 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 754018)
	}

	input = utils.ParseStrings(fmt.Sprint(testPath, "day13_test4.txt"))
	result = problem2(input)
	if result !=779210 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 779210)
	}

	input = utils.ParseStrings(fmt.Sprint(testPath, "day13_test5.txt"))
	result = problem2(input)
	if result != 1261476 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 1261476)
	}

	input = utils.ParseStrings(fmt.Sprint(testPath, "day13_test6.txt"))
	result = problem2(input)
	if result != 1202161486 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 1202161486)
	}
}
