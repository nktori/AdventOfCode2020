package utils

import (
	"fmt"
	"strconv"
)

func PrintDay(day int, name string) {
	fmt.Println("===========================\nDay", day, "-", name, "\n---------------------------")
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprint("Could not convert [", s, "] to int"))
	}
	return i
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinMax(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		min = MinInt(min, v)
		max = MaxInt(max, v)
	}
	return min, max
}

func AbsInt(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}
