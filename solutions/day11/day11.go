package day11

import (
	"fmt"

	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(11, "Seating System")

	input := utils.ParseStrings("/inputs/day11/day11.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int {
	seatPlan := parseInput(input)
	for {
		newPlan, changed := simulateSitting(seatPlan, false, 4)
		if !changed {
			break
		}
		seatPlan = newPlan
	}
	return countOccupied(seatPlan)
}

func problem2(input []string) int {
	seatPlan := parseInput(input)
	for {
		newPlan, changed := simulateSitting(seatPlan, true, 5)
		if !changed {
			break
		}
		seatPlan = newPlan
	}
	return countOccupied(seatPlan)
}

func simulateSitting(seatPlan [][]rune, extendView bool, moveThreshold int) ([][]rune, bool) {
	updated := make([][]rune, len(seatPlan))
	for i := range seatPlan {
		updated[i] = make([]rune, len(seatPlan[i]))
		copy(updated[i], seatPlan[i])
	}
	hasChanged := false
	for i, row := range seatPlan {
		for j, seat := range row {
			seats := visibleSeats(seatPlan, i, j, extendView)
			canSit, shouldMove := canSitShouldMove(seats, moveThreshold)
			if seat == 'L' && canSit {
				updated[i][j] = '#'
				hasChanged = true
			}
			if seat == '#' && shouldMove {
				updated[i][j] = 'L'
				hasChanged = true
			}
		}
	}
	return updated, hasChanged
}

func visibleSeats(seatPlan [][]rune, i, j int, extend bool) []rune {
	directions := [][]int{{-1, 0}, {-1, -1}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	var seats []rune
directionLoop:
	for _, direction := range directions {
		k := 1
		for {
			checkI := i + (direction[0] * k)
			checkJ := j + (direction[1] * k)
			if checkI < 0 || checkI > len(seatPlan)-1 || checkJ < 0 || checkJ > len(seatPlan[0])-1 {
				continue directionLoop
			}
			if seatPlan[checkI][checkJ] == '#' || seatPlan[checkI][checkJ] == 'L' {
				seats = append(seats, seatPlan[checkI][checkJ])
				continue directionLoop
			}
			if !extend {
				break
			}
			k++
		}
	}
	return seats
}

func canSitShouldMove(seats []rune, moveThreshold int) (bool, bool) {
	unavailable := false
	occupiedCount := 0
	for _, v := range seats {
		if v == '#' {
			unavailable = true
			occupiedCount++
		}
	}
	return !unavailable, occupiedCount >= moveThreshold
}

func countOccupied(seatPlan [][]rune) int {
	occupiedSeats := 0
	for _, row := range seatPlan {
		for _, seat := range row {
			if seat == '#' {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func parseInput(input []string) [][]rune {
	var seatPlan [][]rune
	for _, line := range input {
		var row []rune
		for _, c := range line {
			row = append(row, c)
		}
		seatPlan = append(seatPlan, row)
	}
	return seatPlan
}
