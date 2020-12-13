package day12

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day12/"

func TestProblem1(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day12_test.txt"))
	result := problem1(input)
	if result != 25 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 25)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day12_test.txt"))
	result := problem2(input)
	if result != 286 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 286)
	}
}
