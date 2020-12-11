package day11

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day11/"

func TestProblem1(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day11_test.txt"))
	result := problem1(input)
	if result != 37 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 37)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day11_test.txt"))
	result := problem2(input)
	if result != 26 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 26)
	}
}
