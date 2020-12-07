package day06

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day06/"

func TestProblem1(t *testing.T) {
	input := utils.ParseGroups(fmt.Sprint(testPath, "day06_test.txt"))
	result, _ := calculate(input)
	if result != 11 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 11)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseGroups(fmt.Sprint(testPath, "day06_test.txt"))
	_, result := calculate(input)
	if result != 6 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 6)
	}
}
