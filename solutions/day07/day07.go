package day07

import (
	"fmt"
	"strings"

	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(7, "Handy Haversacks")
	input := utils.ParseStrings("/inputs/day07/day07.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int {
	bags := parseInput(input)
	count := 0
	for _, contents := range bags {
		if containsBag(bags, contents, "shiny gold") {
			count++
		}
	}
	return count
}

func problem2(input []string) int {
	return countInternalBags(parseInput(input), "shiny gold")
}

func containsBag(allBags map[string]map[string]int, containedBags map[string]int, targetColour string) bool {
	contains := false
	if len(containedBags) != 0 {
		for bagColour := range containedBags {
			if bagColour == targetColour {
				return true
			}
			contains = contains || containsBag(allBags, allBags[bagColour], targetColour)
		}
	}
	return contains
}

func countInternalBags(allBags map[string]map[string]int, bagColour string) int {
	count := 0
	internalBags := allBags[bagColour]
	if len(internalBags) != 0 {
		for internalBagColour, internalBagCount := range internalBags {
			count += internalBagCount
			count += countInternalBags(allBags, internalBagColour) * internalBagCount
		}
	}
	return count
}

func parseInput(input []string) map[string]map[string]int {
	bags := map[string]map[string]int{}
	for _, line := range input {
		line = strings.ReplaceAll(strings.ReplaceAll(line, "bags", ""), "bag", "")
		split := strings.Split(line, "contain")
		containsMap := make(map[string]int)
		for _, internalBags := range strings.Split(strings.Trim(split[1], "."), ",") {
			internalBags = strings.TrimSpace(internalBags)
			if internalBags != "no other" {
				for _, internalBag := range strings.Split(internalBags, ",") {
					containsMap[internalBag[2:]] = utils.StringToInt(string(internalBag[0]))
				}
			}
			bags[strings.TrimSpace(split[0])] = containsMap
		}
	}
	return bags
}
