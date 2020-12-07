package day01

import (
	"fmt"

	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(1, "Report Repair")
	input := utils.ParseNumbers("/inputs/day01/day01.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []int) int {
	for i := range input {
		iVal := input[i]
		for j := i + 1; j < len(input); j++ {
			jVal := input[j]
			if iVal+jVal == 2020 {
				return iVal * jVal
			}
		}
	}
	panic("Couldn't find solution!")
}

func problem2(input []int) int {
	for i := range input {
		iVal := input[i]
		for j := i + 1; j < len(input)-1; j++ {
			jVal := input[j]
			for k := j + 1; k < len(input); k++ {
				kVal := input[k]
				if iVal+jVal+kVal == 2020 {
					return iVal * jVal * kVal
				}
			}
		}
	}
	panic("Couldn't find solution!")
}
