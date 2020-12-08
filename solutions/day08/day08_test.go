package day08

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day08/"

func TestProblem1(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day08_test.txt"))
	result := problem1(input)
	if result != 5 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 5)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day08_test.txt"))
	result := problem2(input)
	if result != 8 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 8)
	}
}
