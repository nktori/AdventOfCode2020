package day05

import (
	"fmt"
	"sort"

	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(4)
	input := utils.ReadFile("/inputs/day05/day05.txt")
	fmt.Println("Problem 1: ", problem1(input))
	fmt.Println("Problem 2: ", problem2(input))
}

func problem1(input []string) int {
	max := 0
	for _, b := range input {
		max = utils.MaxInt(max, getBoardingId(b))
	}
	return max
}

func problem2(input []string) int {
	ids := make([]int, len(input))
	for i, b := range input {
		ids[i] = getBoardingId(b)
	}
	sort.Ints(ids)
	for i := 0; i < len(ids) - 1; i++ {
		if ids[i+1]-ids[i] == 2 {
			return ids[i] + 1
		}
	}
	panic("Couldn't find boarding pass")
}

func getBoardingId(pattern string) int {
	minRow, maxRow := 0, 127
	minCol, maxCol := 0, 7
	for _, rune := range pattern {
		switch rune {
		case 'F':
			maxRow = minRow + ((maxRow - minRow) / 2)
		case 'B':
			minRow = maxRow - ((maxRow - minRow) / 2)
		case 'L':
			maxCol = minCol + ((maxCol - minCol) / 2)
		case 'R':
			minCol = maxCol - ((maxCol - minCol) / 2)
		}
	}
	if minCol == maxCol && minRow == maxRow {
		return (minRow * 8) + minCol
	}
	panic("Boarding pass invalid")
}
