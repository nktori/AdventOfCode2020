package day12

import (
	"fmt"
	"math"

	"github.com/nktori/AdventOfCode2020/utils"
)

type instruction struct {
	action string
	value  int
}

func Solve() {
	utils.PrintDay(12, "Rain Risk")
	input := utils.ParseStrings("/inputs/day12/day12.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int {
	instructions := parseInput(input)
	ship := map[string]int{"x": 0, "y": 0, "dx": 1, "dy": 0}
	for _, inst := range instructions {
		if inst.action == "L" || inst.action == "R" {
			ship["dx"], ship["dy"] = rotate(ship["dx"], ship["dy"], inst.value, inst.action)
		} else {
			switch inst.action {
			case "N":
				ship["y"] += inst.value * -1
			case "E":
				ship["x"] += inst.value
			case "S":
				ship["y"] += inst.value
			case "W":
				ship["x"] += inst.value * -1
			case "F":
				ship["x"] += inst.value * ship["dx"]
				ship["y"] += inst.value * ship["dy"]
			}
		}
	}
	return utils.AbsInt(ship["x"]) + utils.AbsInt(ship["y"])
}

func problem2(input []string) int {
	instructions := parseInput(input)
	ship := map[rune]int{'x': 0, 'y': 0}
	waypoint := map[rune]int{'x': 10, 'y': -1}
	for _, inst := range instructions {
		if inst.action == "L" || inst.action == "R" {
			waypoint['x'], waypoint['y'] = rotate(waypoint['x'], waypoint['y'], inst.value, inst.action)
		} else {
			switch inst.action {
			case "N":
				waypoint['y'] += inst.value * -1
			case "E":
				waypoint['x'] += inst.value
			case "S":
				waypoint['y'] += inst.value
			case "W":
				waypoint['x'] += inst.value * -1
			case "F":
				ship['x'] += inst.value * waypoint['x']
				ship['y'] += inst.value * waypoint['y']
			}
		}
	}
	return utils.AbsInt(ship['x']) + utils.AbsInt(ship['y'])
}

func rotate(dx, dy, degrees int, direction string) (int, int) {
	if direction == "L" {
		degrees = 360 - degrees
	}
	rad := float64(degrees) * (math.Pi / 180.0)
	newdx := (dx * int(math.Cos(rad))) - (dy * int(math.Sin(rad)))
	dy = (dx * int(math.Sin(rad))) + (dy * int(math.Cos(rad)))
	return newdx, dy
}

func parseInput(input []string) []instruction {
	var instructions []instruction
	for _, v := range input {
		instructions = append(instructions, instruction{
			action: string(v[0]),
			value:  utils.StringToInt(v[1:]),
		})
	}
	return instructions
}
