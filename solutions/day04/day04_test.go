package day04

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day04/"

func TestProblem1(t *testing.T) {
	input := utils.ParseGroups(fmt.Sprint(testPath, "day04_test.txt"))
	result := problem1(input)
	if result != 2 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 2)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseGroups(fmt.Sprint(testPath, "day04_test2.txt"))
	result := problem2(input)
	if result != 4 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 4)
	}
}
