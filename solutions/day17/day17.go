package day17

import (
	"fmt"

	"github.com/nktori/AdventOfCode2020/utils"
)

type point struct {
	x, y, z, w int
}

type gridBoundaries struct {
	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW int
}

func Solve() {
	utils.PrintDay(17, "Conway Cubes")
	input := utils.ParseStrings("/inputs/day17/day17.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int {
	seenCoords, grid := parseInput(input)
	for i := 0; i < 6; i++ {
		seenCoords = cycle(seenCoords, &grid, false)
	}
	return countActive(seenCoords)
}

func problem2(input []string) int {
	seenCoords, grid := parseInput(input)
	for i := 0; i < 6; i++ {
		seenCoords = cycle(seenCoords, &grid, true)
	}
	return countActive(seenCoords)
}

func cycle(seenCoords map[point]rune, grid *gridBoundaries, fourthDim bool) map[point]rune {
	newCoords := make(map[point]rune)
	for x := grid.minX - 1; x <= grid.maxX+1; x++ {
		for y := grid.minY - 1; y <= grid.maxY+1; y++ {
			for z := grid.minZ - 1; z <= grid.maxZ+1; z++ {
				if fourthDim {
					for w := grid.minW - 1; w <= grid.maxW+1; w++ {
						newCoords[point{x, y, z, w}] = getNewValue(x, y, z, w, seenCoords, fourthDim)
					}
				} else {
					newCoords[point{x, y, z, 0}] = getNewValue(x, y, z, 0, seenCoords, fourthDim)
				}
			}
		}
	}
	grid.minX--
	grid.maxX++
	grid.minY--
	grid.maxY++
	grid.minZ--
	grid.maxZ++
	grid.minW--
	grid.maxW++
	return newCoords
}

func getNewValue(x, y, z, w int, seenCoords map[point]rune, fourthDim bool) rune {
	isActive := false
	if coord, ok := seenCoords[point{x, y, z, w}]; ok && coord == '#' {
		isActive = true
	}
	neighbours := countActiveNeighbours(x, y, z, w, seenCoords, fourthDim)
	if isActive && (neighbours == 2 || neighbours == 3) {
		return '#'
	} else if !isActive && neighbours == 3 {
		return '#'
	} else {
		return '.'
	}
}

func countActiveNeighbours(x, y, z, w int, grid map[point]rune, fourthDim bool) int {
	activeNeighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if fourthDim {
					for l := -1; l <= 1; l++ {
						if i == 0 && j == 0 && k == 0 && l == 0 {
							continue
						}
						if checkIsActive(x+i, y+j, z+k, w+l, grid) {
							activeNeighbours++
						}
					}
				} else {
					if i == 0 && j == 0 && k == 0 {
						continue
					}
					if checkIsActive(x+i, y+j, z+k, w, grid) {
						activeNeighbours++
					}
				}
			}
		}
	}
	return activeNeighbours
}

func checkIsActive(x, y, z, w int, grid map[point]rune) bool {
	if v, ok := grid[point{x, y, z, w}]; ok && v == '#' {
		return true
	}
	return false
}

func countActive(coords map[point]rune) int {
	count := 0
	for _, v := range coords {
		if v == '#' {
			count++
		}
	}
	return count
}

func parseInput(input []string) (map[point]rune, gridBoundaries) {
	coords := make(map[point]rune)
	grid := gridBoundaries{}
	for y, line := range input {
		for x, v := range line {
			coords[point{x, y, 0, 0}] = v
			grid.maxX = x
		}
		grid.maxY = y
	}
	return coords, grid
}
