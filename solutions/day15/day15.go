package day15

import (
	"fmt"

	"github.com/nktori/AdventOfCode2020/utils"
)

var input = []int{2,0,6,12,1,3}

func Solve() {
	utils.PrintDay(15, "Rambunctious Recitation")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []int) int {
	return compute(input, 2020)
}

func problem2(input []int) int {
	return compute(input, 30000000)
}

func compute(input []int, limit int) int {
	var lastSpoken int
	numbersSpoken := make(map[int][]int)
	counter := 1
	for _, v := range input {
		numbersSpoken[v] = append(numbersSpoken[v], counter)
		lastSpoken = v
		counter++
	}

	for counter <= limit {
		whenSaid := numbersSpoken[lastSpoken]
		if len(whenSaid) <= 1 {
			numbersSpoken[0] = append(numbersSpoken[0], counter)
			lastSpoken = 0
			counter++
			continue
		}
		timesSaid := len(whenSaid)
		nextNumber := whenSaid[timesSaid-1] - whenSaid[timesSaid-2]
		numbersSpoken[nextNumber] = append(numbersSpoken[nextNumber], counter)
		lastSpoken = nextNumber
		counter++
	}
	return lastSpoken
}