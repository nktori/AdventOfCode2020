package day03

import (
	"fmt"

	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(3, "Toboggan Trajectory")
	input := utils.ParseStrings("/inputs/day03/day03.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int {
	return countTrees(input, 3, 1)
}

func problem2(input []string) int {
	return countTrees(input, 1, 1) *
		countTrees(input, 3, 1) *
		countTrees(input, 5, 1) *
		countTrees(input, 7, 1) *
		countTrees(input, 1, 2)
}

func countTrees(input []string, stepRight int, stepDown int) int {
	count := 0
	columnIndex := 0
	for i := 0; i < len(input); i += stepDown {
		if columnIndex >= len(input[i]) {
			columnIndex = columnIndex - len(input[i])
		}
		if string(input[i][columnIndex]) == "#" {
			count++
		}
		columnIndex += stepRight
	}
	return count
}
