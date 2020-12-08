package day08

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nktori/AdventOfCode2020/utils"
)

type instruction struct {
	operation string
	argument  int
}

func Solve() {
	utils.PrintDay(8, "Handheld Halting")
	input := utils.ParseStrings("/inputs/day08/day08.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int {
	accumulator, _ := runInstructions(parseInput(input))
	return accumulator
}

func problem2(input []string) int {
	instructions := parseInput(input)
	accumulator := 0
	for i, instruct := range instructions {
		match, _ := regexp.MatchString("^(jmp|nop)$", instruct.operation)
		if match {
			tmp := make([]instruction, len(instructions))
			copy(tmp, instructions)
			if instruct.operation == "jmp" {
				tmp[i].operation = "nop"
			} else {
				tmp[i].operation = "jmp"
			}
			count, looped := runInstructions(tmp)
			if !looped {
				accumulator = count
				break
			}
		}
	}
	return accumulator
}

func runInstructions(instructions []instruction) (int, bool) {
	visited := make(map[int]bool)
	index := 0
	accumulator := 0
	looped := false
	for !looped {
		if visited[index] {
			looped = true
			break
		}
		if index >= len(instructions) {
			break
		}
		visited[index] = true
		instruct := instructions[index]
		switch instruct.operation {
		case "acc":
			accumulator += instruct.argument
			index++
		case "jmp":
			index += instruct.argument
		case "nop":
			index++
		default:
			panic(fmt.Sprint("Could not parse operation [", instruct.operation, "]"))
		}
	}
	return accumulator, looped
}

func parseInput(input []string) []instruction {
	var instructions []instruction
	for _, line := range input {
		split := strings.Split(line, " ")
		instructions = append(instructions, instruction{
			operation: split[0],
			argument:  utils.StringToInt(split[1]),
		})
	}
	return instructions
}
