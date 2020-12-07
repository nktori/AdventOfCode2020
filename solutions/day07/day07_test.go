package day07

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day07/"

func TestProblem1(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day07_test.txt"))
	result := problem1(input)
	if result != 4 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 4)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day07_test.txt"))
	result := problem2(input)
	if result != 32 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 32)
	}

	input = utils.ParseStrings(fmt.Sprint(testPath, "day07_test2.txt"))
	result = problem2(input)
	if result != 126 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 126)
	}
}
