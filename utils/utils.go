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
	} else {
		return b
	}
}
