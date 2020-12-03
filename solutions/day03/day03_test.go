package day03

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/test/"

func TestProblem1(t *testing.T) {
	input := utils.ReadFile(fmt.Sprint(testPath,"day03_test.txt"))
	result := problem1(input)
	if result != 7 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 7)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ReadFile(fmt.Sprint(testPath,"day03_test.txt"))
	result := problem2(input)
	if result != 336 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 336)
	}
}
