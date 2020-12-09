package day09

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day09/"

func TestProblem1(t *testing.T) {
	input := utils.ParseNumbers(fmt.Sprint(testPath, "day09_test.txt"))
	result := problem1(input, 5)
	if result != 127 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 127)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseNumbers(fmt.Sprint(testPath, "day09_test.txt"))
	result := problem2(input, 127)
	if result != 62 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 62)
	}
}
