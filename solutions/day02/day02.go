package day02

import (
	"fmt"
	"strings"

	"github.com/nktori/AdventOfCode2020/utils"
)

type rule struct {
	min int
	max int
	letter string
}

func Solve()  {
	utils.PrintDay(2)
	input := utils.ReadFile("/inputs/day02.txt")
	fmt.Println("Problem 1: ", problem1(input))
	fmt.Println("Problem 2: ", problem2(input))
}

func problem1(input []string) int {
	inputMap := parseInput(input)
	correct := 0
	for key, value := range inputMap {
		count := strings.Count(key, value.letter)
		if count >= value.min && count <= value.max {
			correct++
		}
	}
	return correct
}

func problem2(input []string) int {
	inputMap := parseInput(input)
	correct := 0
	for key, value := range inputMap {
		if (string(key[value.min - 1]) == value.letter) != (string(key[value.max - 1]) == value.letter) {
			correct++
		}
	}
	return correct
}

func parseInput(input []string) map[string]rule {
	parsedInputs := make(map[string]rule)
	for _, s := range input {
		split := strings.Split(s, ":")
		rules := strings.Split(split[0], " ")
		counts := strings.Split(rules[0], "-")
		parsedInputs[strings.TrimSpace(split[1])] = rule {
			min: utils.StringToInt(counts[0]),
			max: utils.StringToInt(counts[1]),
			letter: rules[1],
		}
	}
	return parsedInputs
}