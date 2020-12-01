package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	path, err := os.Getwd()
	if err != nil {
		panic("Could not get working directory")
	}
	contents, err := ioutil.ReadFile(fmt.Sprint(path,filename))
	if err != nil {
		panic(fmt.Sprint("Error reading file: ", filename))
	}
	return string(contents)
}

func ParseInput(inputFile string) []int {
	stringInput := strings.Split(readFile(inputFile), "\n")
	numbers := make([]int, len(stringInput))
	for i, s := range stringInput {
		num, _ := strconv.Atoi(s)
		numbers[i] = num
	}
	return numbers
}
