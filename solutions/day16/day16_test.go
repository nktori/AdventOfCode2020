package day16

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day16/"

func TestProblem1(t *testing.T) {
	input := utils.ParseAsGroups(fmt.Sprint(testPath, "day16_test.txt"))
	result := problem1(input)
	if result != 71 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 71)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseAsGroups(fmt.Sprint(testPath, "day16_test2.txt"))
	result := problem2(input)
	if result != 11 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 11)
	}
	problem2(input)
}
