package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("day5/input.txt")
	boardingPasses := strings.Split(string(data), "\n")
	seatIDs := make([]int, len(boardingPasses))
	highest := 0
	for i, boardingPass := range boardingPasses {
		rows := boardingPass[:7]
		rowNo := search(rows, 0, 127, "F", "B")
		cols := boardingPass[7:]
		colNo := search(cols, 0, 7, "L", "R")
		seatID := rowNo*8 + colNo
		seatIDs[i] = seatID
		if seatID > highest {
			highest = seatID
		}
	}
	fmt.Println("Highest Seat ID: ", highest)
	sort.Ints(seatIDs)
	low := seatIDs[0]
	high := seatIDs[len(seatIDs)-1]
	for i := low; i <= high; i++ {
		if i != seatIDs[i - low] {
			fmt.Println("My Seat ID: ", i)
			break
		}
	}
}

func search(chars string, start int, end int, down string, up string) int {
	res := start
	for _, c := range chars {
		half := (start + end) / 2
		if up == string(c) {
			start = half + 1
			res = end
		} else {
			end = half
			res = start
		}
	}
	return res
}
