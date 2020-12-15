package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func arrayIndex(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

type adapter struct {
	jolt     int
	adapters []adapter
}

func convertToIntArray(strings []string) []int {
	ints := make([]int, len(strings), len(strings))
	for i, s := range strings {
		number, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = number
	}
	return ints
}

var cache map[int]int

func countCombinations(group []int) int {
	if len(group) == 1 {
		return 1
	}

	target := group[len(group)-1]
	if v, ok := cache[target]; ok {
		return v
	}

	var combinations int

	for i := len(group) - 2; i >= 0 && target-group[i] <= 3; i-- {
		combinations += countCombinations(group[:i+1])
	}
	cache[target] = combinations
	return combinations

}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	ints := convertToIntArray(lines)
	sort.Ints(ints)
	ints = append(ints, ints[len(ints)-1]+3)
	cache = make(map[int]int, len(ints))

	var ones, threes int
	last := 0
	currentGroup := []int{0}
	result2 := 1
	for _, number := range ints {
		diff := number - last
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
			result2 *= countCombinations(currentGroup)
			currentGroup = nil
		}
		currentGroup = append(currentGroup, number)
		last = number
	}

	println(ones*threes, result2)
}
