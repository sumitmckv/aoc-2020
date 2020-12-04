package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	data, _ := ioutil.ReadFile("day3/input.txt")
	arr := strings.Split(string(data), "\n")
	fmt.Println(noOfTrees(arr, 3, 1))
}

func puzzle2() {
	data, _ := ioutil.ReadFile("day3/input.txt")
	arr := strings.Split(string(data), "\n")
	rightArr := [5]int{1, 3, 5, 7, 1}
	down := 1
	res := 1
	for i, right := range rightArr {
		if i == 4 {
			down = 2
		}
		res *= noOfTrees(arr, right, down)
	}
	fmt.Println(res)
}

func noOfTrees(arr []string, right int, down int) int {
	l1 := len(arr)
	l2 := len(arr[0])
	res := 0
	j := right
	for i := down; i < l1; i += down {
		str := arr[i]
		if string(str[j]) == "#" {
			res++
		}
		j += right
		j = j % l2
	}
	return res
}
