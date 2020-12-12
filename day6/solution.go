package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main()  {
	data, _ := ioutil.ReadFile("day6/input.txt")
	arr := strings.Split(string(data), "\n\n")
	res1 := 0
	res2 := 0
	for _, str := range arr {
		kv := make(map[string]int)
		entries := strings.Split(str, "\n")
		for _, entry := range entries {
			entryArr := strings.Split(entry, "")
			for _, e := range entryArr {
				v, ok := kv[e]
				if !ok {
					kv[e] = 1
				} else {
					kv[e] = v + 1
				}
			}
		}
		res1 += len(kv)
		for _, v := range kv {
			if v == len(entries) {
				res2++
			}
		}
	}
	fmt.Println("Puzzle 1: ", res1)
	fmt.Println("Puzzle 2: ", res2)
}
