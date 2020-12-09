package day09

import (
	"fmt"
	"sort"

	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(9, "Encoding Error")
	input := utils.ParseNumbers("/inputs/day09/day09.txt")
	prob1 := problem1(input, 25)
	fmt.Println("Problem 1:", prob1)
	fmt.Println("Problem 2:", problem2(input, prob1))
}

func problem1(input []int, preambleSize int) int {
	for i := preambleSize; i < len(input); i++ {
		tmp := make([]int, preambleSize)
		copy(tmp, input[i-preambleSize:i])
		if !canSum(tmp, input[i]) {
			return input[i]
		}
	}
	panic("Could not find solution!")
}

func problem2(input []int, invalidTarget int) int {
	i, j, sum := 0, 1, input[0]
	for true {
		if sum == invalidTarget {
			break
		} else if sum < invalidTarget {
			sum += input[j]
			j++
		} else if sum > invalidTarget {
			sum -= input[i]
			i++
		}
	}
	min, max := utils.MinMax(input[i:j])
	return min + max
}

func canSum(numbers []int, target int) bool {
	sort.Ints(numbers)
	for _, first := range numbers {
		second := target - first
		j := sort.SearchInts(numbers, second)
		if j < len(numbers) && second == numbers[j] {
			return true
		}
	}
	return false
}
