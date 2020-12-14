package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nktori/AdventOfCode2020/utils"
)

func Solve() {
	utils.PrintDay(14, "Docking Data")
	input := utils.ParseStrings("/inputs/day14/day14.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int64 {
	mask := strings.Repeat("X", 36)
	memory := make(map[int]int64)
	for _, line := range input {
		lineSplit := strings.Split(line, " = ")
		if lineSplit[0] == "mask" {
			mask = lineSplit[1]
		} else {
			reg, _ := regexp.Compile("[^0-9]+")
			memSlot := utils.StringToInt(reg.ReplaceAllString(lineSplit[0], ""))
			memory[memSlot] = bitToInt(decode(utils.StringToInt(lineSplit[1]), mask, false))
		}
	}
	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	return sum
}

func problem2(input []string) int {
	mask := strings.Repeat("X", 36)
	memory := make(map[int64]int)
	for _, line := range input {
		lineSplit := strings.Split(line, " = ")
		if lineSplit[0] == "mask" {
			mask = lineSplit[1]
		} else {
			reg, _ := regexp.Compile("[^0-9]+")
			memSlot := utils.StringToInt(reg.ReplaceAllString(lineSplit[0], ""))
			value := utils.StringToInt(lineSplit[1])
			memAddresses := permute("", decode(memSlot, mask, true))
			for _, m := range memAddresses {
				memory[bitToInt(m)] = value
			}
		}
	}
	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}

func permute(address string, maskedAddress string) []string {
	if len(maskedAddress) == 0 {
		return []string{address}
	}
	if maskedAddress[0] == 'X' {
		return append(permute(address+"0", maskedAddress[1:]), permute(address+"1", maskedAddress[1:])...)
	} else {
		return permute(address+string(maskedAddress[0]), maskedAddress[1:])
	}
}

func decode(number int, mask string, memAddDecoder bool) string {
	zero := strings.Repeat("0", 36)
	bitNum := strconv.FormatInt(int64(number), 2)
	bit36Num := fmt.Sprintf("%s%s", zero[0:len(zero)-len(bitNum)], bitNum)
	maskedBitNum := ""
	for i, c := range bit36Num {
		if !memAddDecoder {
			if mask[i] == 'X' {
				maskedBitNum += string(c)
			} else {
				maskedBitNum += string(mask[i])
			}
		} else {
			switch mask[i] {
			case 'X':
				maskedBitNum += "X"
			case '0':
				maskedBitNum += string(c)
			case '1':
				maskedBitNum += "1"
			}
		}
	}
	return maskedBitNum
}

func bitToInt(bitString string) int64 {
	if i, err := strconv.ParseInt(bitString, 2, 64); err != nil {
		panic("Could not convert [" + bitString + "] to integer")
	} else {
		return i
	}
}
