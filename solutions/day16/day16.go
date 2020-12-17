package day16

import (
	"fmt"
	"github.com/nktori/AdventOfCode2020/utils"
	"strings"
)

type rule struct {
	name   string
	ranges [][]int
}

func Solve() {
	utils.PrintDay(16, "Ticket Translation")
	input := utils.ParseAsGroups("/inputs/day16/day16.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input []string) int {
	rules, _, nearbyTickets := parseInput(input)
	sum := 0
	for _, ticket := range nearbyTickets {
	ticket:
		for _, v := range ticket {
			for _, r := range rules {
				if isValidForRanges(v, r.ranges) {
					continue ticket
				}
			}
			sum += v
		}
	}
	return sum
}

func problem2(input []string) int {
	rules, yourTicket, nearbyTickets := parseInput(input)
	validTickets := [][]int{yourTicket}
	for _, ticket := range nearbyTickets {
		ticketInvalid := false
		for _, v := range ticket {
			ticketInvalid = ticketInvalid || !isValidForRules(v, rules)
		}
		if !ticketInvalid {
			validTickets = append(validTickets, ticket)
		}
	}
	possibleFields := make(map[int]map[string]bool)
	for i := range yourTicket {
		possibleFields[i] = make(map[string]bool)
		for _, k := range rules {
			possibleFields[i][k.name] = true
		}
	}
	for _, ticket := range validTickets {
		for i, v := range ticket {
			for _, r := range rules {
				if !isValidForRanges(v, r.ranges) {
					delete(possibleFields[i], r.name)
				}
			}
		}
	}
	finalPositions := make(map[int]string)
	for len(finalPositions) < len(possibleFields) {
	possibleFieldsLoop:
		for i, fields := range possibleFields {
			if len(fields) > 1 {
				continue
			}
			for fieldName := range fields {
				finalPositions[i] = fieldName
				for _, fields := range possibleFields {
					delete(fields, fieldName)
				}
				break possibleFieldsLoop
			}
		}
	}
	answer := 1
	for k, v := range finalPositions {
		if strings.HasPrefix(v, "departure") {
			answer *= yourTicket[k]
		}
	}
	return answer
}

func isValidForRanges(v int, ranges [][]int) bool {
	return (v >= ranges[0][0] && v <= ranges[0][1]) || (v >= ranges[1][0] && v <= ranges[1][1])
}

func isValidForRules(v int, rules []rule) bool {
	valid := false
	for _, r := range rules {
		valid = valid || isValidForRanges(v, r.ranges)
	}
	return valid
}

func parseInput(input []string) ([]rule, []int, [][]int) {
	rules := parseRules(input[0])
	yourTicket := parseTicket(strings.Split(input[1], "\n")[1])
	nearbyTicketsString := strings.Split(input[2], "\n")
	var nearbyTickets [][]int
	for i := 1; i < len(nearbyTicketsString); i++ {
		nearbyTickets = append(nearbyTickets, parseTicket(nearbyTicketsString[i]))
	}
	return rules, yourTicket, nearbyTickets
}

func parseRules(rulesString string) []rule {
	var rules []rule
	for _, line := range strings.Split(rulesString, "\n") {
		fields := strings.Fields(line)
		fieldsLen := len(fields)

		name := strings.Trim(strings.Join(fields[:fieldsLen-3], " "), ":")
		r1, r2 := getRange(fields[fieldsLen-3]), getRange(fields[fieldsLen-1])
		ranges := [][]int{{utils.StringToInt(r1[0]), utils.StringToInt(r1[1])}, {utils.StringToInt(r2[0]), utils.StringToInt(r2[1])}}

		rules = append(rules, rule{name, ranges})
	}
	return rules
}

func getRange(r string) []string {
	return strings.Fields(strings.ReplaceAll(r, "-", " "))
}

func parseTicket(t string) []int {
	var arr []int
	for _, v := range strings.Split(t, ",") {
		arr = append(arr, utils.StringToInt(v))
	}
	return arr
}
