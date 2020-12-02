package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	path, err := os.Getwd()
	if err != nil {
		panic("Could not get working directory")
	}
	contents, err := ioutil.ReadFile(fmt.Sprint(path,filename))
	if err != nil {
		panic(fmt.Sprint("Error reading file: ", filename))
	}
	return strings.Split(string(contents), "\n")
}

func ParseNumbers(filename string) []int {
	stringInput := ReadFile(filename)
	numbers := make([]int, len(stringInput))
	for i, s := range stringInput {
		numbers[i] = StringToInt(s)
	}
	return numbers
}
