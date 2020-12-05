package day05

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/day05/"

func TestProblem1(t *testing.T) {
	input := utils.ReadFile(fmt.Sprint(testPath, "day05_test.txt"))
	result := problem1(input)
	if result != 820 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 820)
	}
}
