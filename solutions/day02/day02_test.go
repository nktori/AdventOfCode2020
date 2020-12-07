package day02

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day02/"

func TestProblem1(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day02_test.txt"))
	result := problem1(input)
	if result != 2 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 2)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseStrings(fmt.Sprint(testPath, "day02_test.txt"))
	result := problem2(input)
	if result != 1 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 1)
	}
}
