package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nktori/AdventOfCode2020/utils"
)

type passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColour     string
	eyeColour      string
	id             string
}

func Solve() {
	utils.PrintDay(4, "Passport Processing")
	input := utils.ParseGroups("/inputs/day04/day04.txt")
	fmt.Println("Problem 1:", problem1(input))
	fmt.Println("Problem 2:", problem2(input))
}

func problem1(input [][]string) int {
	passports := parseInput(input)
	valid := 0
	for _, p := range passports {
		if hasRequiredFields(p) {
			valid++
		}
	}
	return valid
}

func problem2(input [][]string) int {
	passports := parseInput(input)
	valid := 0
	for _, p := range passports {
		if hasValidFields(p) {
			valid++
		}
	}
	return valid
}

func hasRequiredFields(pass passport) bool {
	return pass.birthYear != "" &&
		pass.issueYear != "" &&
		pass.expirationYear != "" &&
		pass.height != "" &&
		pass.hairColour != "" &&
		pass.eyeColour != "" &&
		pass.id != ""
}

func hasValidFields(pass passport) bool {
	if !hasRequiredFields(pass) {
		return false
	}
	return matchYear(pass.birthYear, 1920, 2002) &&
		matchYear(pass.issueYear, 2010, 2020) &&
		matchYear(pass.expirationYear, 2020, 2030) &&
		matchHeight(pass.height) &&
		matchHairColour(pass.hairColour) &&
		matchEyeColour(pass.eyeColour) &&
		matchPassportId(pass.id)
}

func matchYear(number string, min, max int) bool {
	i, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}

func matchHeight(height string) bool {
	match, _ := regexp.MatchString("(^[1]([5-8][0-9]|[9][0-3])cm$)|((^[59]|[6][0-9]|[7][0-6])in$)", height)
	return match
}

func matchHairColour(hairColour string) bool {
	match, _ := regexp.MatchString("^#[0-9a-f]{6}$", hairColour)
	return match
}

func matchEyeColour(eyeColour string) bool {
	match, _ := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", eyeColour)
	return match
}

func matchPassportId(id string) bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", id)
	return match
}

func parseInput(input [][]string) []passport {
	var passports []passport
	for _, group := range input {
		fields := map[string]string{}
		for _, field := range group {
			split := strings.Split(field, ":")
			fields[split[0]] = split[1]
		}
		passports = append(passports, passport{
			birthYear:      fields["byr"],
			issueYear:      fields["iyr"],
			expirationYear: fields["eyr"],
			height:         fields["hgt"],
			hairColour:     fields["hcl"],
			eyeColour:      fields["ecl"],
			id:             fields["pid"],
		})
	}
	return passports
}
