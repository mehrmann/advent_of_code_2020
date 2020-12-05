package main

import (
	"io/ioutil"
	"strings"
)

func calcRow(s string) int {
	var min int = 0
	var max int = 127
	//64,32,16,8,4,2,1
	step := []int{64, 32, 16, 8, 4, 2, 1}

	for i, p := range s[0:7] {
		if p == 'B' {
			min += step[i]
		} else {
			max -= step[i]
		}
		//println(i, p, min, max)
	}
	return min
}

func calcColumn(s string) int {
	var min int = 0
	var max int = 8
	//64,32,16,8,4,2,1
	step := []int{4, 2, 1}

	for i, p := range s[7:10] {
		if p == 'R' {
			min += step[i]
		} else {
			max -= step[i]
		}
		//println(i, p, min, max)
	}
	return min
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	used := make([]bool, 128*8, 128*8)
	lines := strings.Split(string(b), "\n")
	max := 0

	for _, l := range lines {
		row := calcRow(l)
		col := calcColumn(l)
		seat := (row * 8) + col
		if seat > max {
			max = seat
		}
		used[seat] = true
	}

	println(max)
	for i := 1; i < 128*8; i++ {
		if used[i-1] == true && used[i] == false && used[i+1] == true {
			println(i)
		}
	}
}
