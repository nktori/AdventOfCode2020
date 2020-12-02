package utils

import (
	"fmt"
	"strconv"
)

func PrintDay(day int) {
	fmt.Println("=======================\nDay ",day,"\n-----------------------")
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprint("Could not convert [",s,"] to int"))
	}
	return i
}
