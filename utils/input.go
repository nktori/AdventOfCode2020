package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readFile(filename string) string {
	path, err := os.Getwd()
	if err != nil {
		panic("Could not get working directory")
	}
	contents, err := ioutil.ReadFile(fmt.Sprint(path, filename))
	if err != nil {
		panic(fmt.Sprint("Error reading file:", filename))
	}
	return string(contents)
}

func ParseStrings(filename string) []string {
	return strings.Split(readFile(filename), "\n")
}

func ParseNumbers(filename string) []int {
	stringInput := ParseStrings(filename)
	numbers := make([]int, len(stringInput))
	for i, s := range stringInput {
		numbers[i] = StringToInt(s)
	}
	return numbers
}

func ParseGroups(filename string) [][]string {
	var groups [][]string
	for _, group := range strings.Split(strings.TrimSpace(readFile(filename)), "\n\n") {
		groups = append(groups, strings.Fields(group))
	}
	return groups
}

func ParseAsGroups(filename string) []string {
	return strings.Split(strings.TrimSpace(readFile(filename)), "\n\n")
}
