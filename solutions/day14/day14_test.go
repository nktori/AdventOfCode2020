package day14

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day14/"

func TestProblem1(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day14_test.txt"))
	result := problem1(input)
	if result != 165 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 165)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day14_test2.txt"))
	result := problem2(input)
	if result != 208 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 208)
	}
}

