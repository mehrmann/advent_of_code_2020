package main

import (
	"io/ioutil"
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

func findContigiousSum(ints []int, number int) (int, int) {
	for i, n := range ints {
		sum := n
		min := n
		max := n

		for x := i + 1; sum < number && x < len(ints); x++ {
			if ints[x] < min {
				min = ints[x]
			}
			if ints[x] > max {
				max = ints[x]
			}
			sum += ints[x]
			if sum == number {
				return min, max
			}
		}

	}
	return -1, -1
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	ints := convertToIntArray(lines)
	preamblesize := 25

	valid := make([]int, 0, 25)
	for i, number := range ints {

		if i < preamblesize {
			valid = append(valid, number)
		} else {
			found := false
			for j, n := range valid {
				//take number, subtract check if found in valid
				remainder := number - n
				remainderPos := arrayIndex(valid, remainder)
				if remainderPos != -1 && remainderPos != j {
					//seems to be working
					found = true
					break
				}
			}
			if !found {
				println(number)
				min, max := findContigiousSum(ints, number)
				println(min + max)
				break
			} else {
				valid = valid[1:]
				valid = append(valid, number)
			}
		}

	}

}
