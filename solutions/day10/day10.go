package day10

import (
	"fmt"
	"sort"

	"github.com/nktori/AdventOfCode2020/utils"
	"gonum.org/v1/gonum/stat/combin"
)

func Solve() {
	utils.PrintDay(10, "Adapter Array")

	input := utils.ParseNumbers("/inputs/day10/day10.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []int) int {
	sort.Ints(input)
	oneCount, threeCount, prev := 0, 0, 0
	for _, v := range input {
		if v-prev == 1 {
			oneCount++
		} else if v-prev == 3 {
			threeCount++
		}
		prev = v
	}
	return oneCount * (threeCount + 1)
}

func problem2(input []int) int {
	sort.Ints(input)
	combinationsLengths := getValidCombinations(getOneDiffs(input))
	answer := 1
	for _, v := range combinationsLengths {
		answer *= v
	}
	return answer
}

func getValidCombinations(input [][]int) []int {
	var combinationsLengths []int
	for _, set := range input {
		size := len(set)
		var combinations [][]int
		for i := size; i > 0; i-- {
			combinations = append(combinations, combin.Combinations(size, i)...)
		}
		var validCombinations [][]int
		for _, combination := range combinations {
			if combination[0] == 0 && combination[len(combination)-1] == size-1 {
				valid := true
				prev := 0
				for _, v := range combination {
					if v-prev > 3 {
						valid = false
					}
					prev = v
				}
				if valid {
					validCombinations = append(validCombinations, combination)
				}
			}
		}
		combinationsLengths = append(combinationsLengths, len(validCombinations))
	}
	for _, v := range combinationsLengths {
		if v < 0 {
			fmt.Println(combinationsLengths)
		}
	}
	return combinationsLengths
}

func getOneDiffs(input []int) [][]int {
	var oneDiffs [][]int
	var oneDiff []int
	prev := 0
	for _, v := range input {
		if v-prev == 1 {
			oneDiff = append(oneDiff, prev)
		} else if len(oneDiff) != 0 {
			oneDiff = append(oneDiff, prev)
			oneDiffs = append(oneDiffs, oneDiff)
			oneDiff = nil
		}
		prev = v
	}
	if len(oneDiff) != 0 {
		oneDiff = append(oneDiff, prev)
		oneDiffs = append(oneDiffs, oneDiff)
	}
	return oneDiffs
}
