package day13

import (
	"fmt"
	"math"
	"strings"

	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(13, "Shuttle Search")
	input := utils.ParseStrings("/inputs/day13/day13.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int {
	earliestTime := utils.StringToInt(input[0])
	var validBuses []int
	for _, bus := range strings.Split(input[1], ",") {
		if bus != "x" {
			validBuses = append(validBuses, utils.StringToInt(bus))
		}
	}
	timesToId := make(map[int]int)
	for _, bus := range validBuses {
		time := 0
		for time < earliestTime {
			time += bus
		}
		timesToId[time] = bus
	}
	minTime := math.MaxInt64
	for k := range timesToId {
		if k < minTime {
			minTime = k
		}
	}
	return (minTime - earliestTime) * timesToId[minTime]
}

func problem2(input []string) int {
	busList := strings.Split(input[1], ",")
	var buses [][]int
	for i, bus := range busList {
		if bus == "x" {
			continue
		}
		buses = append(buses, []int{i, utils.StringToInt(bus)})
	}
	increment := buses[0][1]
	busTimes := []int{increment}
	time := 0
	for i := 1; i < len(buses); i++ {
		for {
			time += increment
			if (time + buses[i][0]) % buses[i][1] == 0 {
				increment = utils.LCM(buses[i-1][1], buses[i][1], busTimes...)
				busTimes = append(busTimes, buses[i][1])
				break
			}
		}
	}
	return time
}
