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
	data, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		fmt.Println("Error occured while reading file", err)
		return
	}
	arr := strings.Split(string(data), "\n")
	len := len(arr) - 1
	sumMap := make(map[int]int)
	for i := 0; i < len; i++ {
		iVal, _ := strconv.Atoi(arr[i])
		rem := 2020 - iVal
		if rem == sumMap[iVal] {
			fmt.Println(iVal * rem)
			return
		}
		sumMap[rem] = iVal
	}
}

func puzzle2() {
	data, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		fmt.Println("Error occured while reading file", err)
		return
	}
	arr := strings.Split(string(data), "\n")
	len := len(arr) - 1
	sumMap := make(map[int]int)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			iVal, _ := strconv.Atoi(arr[i])
			jVal, _ := strconv.Atoi(arr[j])
			sum := iVal + jVal
			rem := 2020 - sum
			if rem > 0 && rem == sumMap[jVal] {
				fmt.Println(iVal * jVal * rem)
				return
			}
			if rem > 0 {
				sumMap[rem] = jVal
			}
		}
	}
}
