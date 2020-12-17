package day17

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day17/"

func TestProblem1(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day17_test.txt"))
	result := problem1(input)
	if result != 112 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 112)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day17_test.txt"))
	result := problem2(input)
	if result != 848 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 848)
	}
}

