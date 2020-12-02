package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	data, _ := ioutil.ReadFile("day2/input.text")
	arr := strings.Split(string(data), "\n")
	res := 0
	for _, str := range arr {
		list := strings.Split(str, " ")
		occurences := strings.Split(list[0], "-")
		minRange, _ := strconv.Atoi(occurences[0])
		maxRange, _ := strconv.Atoi(occurences[1])
		element := list[1][:len(list[1])-1]
		password := list[2]
		count := 0
		for _, s := range password {
			if element == string(s) {
				count++
			}
		}
		if count >= minRange && count <= maxRange {
			res++
		}
	}
	fmt.Println(res)
}

func puzzle2() {
	data, _ := ioutil.ReadFile("day2/input.text")
	arr := strings.Split(string(data), "\n")
	res := 0
	for _, str := range arr {
		list := strings.Split(str, " ")
		occurences := strings.Split(list[0], "-")
		minRange, _ := strconv.Atoi(occurences[0])
		maxRange, _ := strconv.Atoi(occurences[1])
		element := list[1][:len(list[1])-1]
		password := list[2]
		count := 0
		if element == string(password[minRange-1]) {
			count++
		}
		if element == string(password[maxRange-1]) {
			count++
		}
		if count == 1 {
			res++
		}
	}
	fmt.Println(res)
}
