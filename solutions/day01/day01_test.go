package day01

import (
	"fmt"
	"testing"

	"github.com/nktori/AdventOfCode2020/utils"
)

var testPath = "/../../inputs/test/"

func TestProblem1(t *testing.T) {
	input := utils.ParseInput(fmt.Sprint(testPath,"day01.txt"))
	result := problem1(input)
	if result != 514579 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 514579)
	}
}

func TestProblem2(t *testing.T) {
	input := utils.ParseInput(fmt.Sprint(testPath,"day01.txt"))
	result := problem2(input)
	if result != 241861950 {
		t.Errorf("Failure - received: %d, expected: %d.", result, 241861950)
	}
}
