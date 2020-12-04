package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("day4/input.txt")
	arr := strings.Split(string(data), "\n\n")
	validPassportPart1 := 0
	validPassportPart2 := 0
	for _, v := range arr {
		fields := strings.Split(strings.ReplaceAll(v, "\n", " "), " ")
		fieldMap := make(map[string]string)
		for _, field := range fields {
			entry := strings.Split(field, ":")
			fieldMap[entry[0]] = entry[1]
		}
		l := len(fieldMap)
		_, ok := fieldMap["cid"]
		if l == 8 || (l == 7 && !ok) {
			validPassportPart1++
			if !isValidYear(fieldMap["byr"], 1920, 2002) {
				continue
			}
			if !isValidYear(fieldMap["iyr"], 2010, 2020) {
				continue
			}
			if !isValidYear(fieldMap["eyr"], 2020, 2030) {
				continue
			}
			if !isValidHeight(fieldMap["hgt"]) {
				continue
			}
			if !isValidHairColor(fieldMap["hcl"]) {
				continue
			}
			if !isValidEyeColor(fieldMap["ecl"]) {
				continue
			}
			if !isValidPassportID(fieldMap["pid"]) {
				continue
			}
			validPassportPart2++
		}
	}
	fmt.Println("Puzzle 1: ", validPassportPart1)
	fmt.Println("Puzzle 2: ", validPassportPart2)
}

func isValidYear(yearStr string, start int, end int) bool {
	if len(yearStr) < 4 {
		return false
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}
	return year >= start && year <= end
}

func isValidHeight(hgt string) bool {
	l := len(hgt)
	if l < 3 {
		return false
	}

	val, err := strconv.Atoi(hgt[:l-2])
	if err != nil {
		return false
	}
	unit := hgt[l-2:]
	if unit == "cm" && val >= 150 && val <= 193 {
		return true
	}
	return unit == "in" && val >= 59 && val <= 76
}

func isValidHairColor(hcl string) bool {
	m, _ := regexp.MatchString(`^#[a-f0-9]{6}$`, hcl)
	return m
}

func isValidEyeColor(ecl string) bool {
	possibleValue := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, v := range possibleValue {
		if v == ecl {
			return true
		}
	}
	return false
}

func isValidPassportID(pid string) bool {
	m, _ := regexp.MatchString(`^[0-9]{9}$`, pid)
	return m
}
