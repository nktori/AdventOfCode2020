package day06

import (
	"fmt"
	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(6, "Custom Customs")
	input := utils.ParseGroups("/inputs/day06/day06.txt")
	problem1, problem2 := calculate(input)
	fmt.Println("Problem 1:", problem1)
	fmt.Println("Problem 2:", problem2)
}

func calculate(fields [][]string) (int, int) {
	unionCount := 0
	intersectionCount := 0
	for _, group := range fields {
		set := make(map[rune]int)
		for _, person := range group {
			for _, answer := range person {
				set[answer]++
				if set[answer] == len(group) {
					intersectionCount++
				}
			}
		}
		unionCount += len(set)
	}
	return unionCount, intersectionCount
}
