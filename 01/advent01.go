package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	var nums = make([]int, 0, len(lines))

	for _, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	for start, num1 := range nums {
		for _, num2 := range nums[(start + 1):] {
			if num1+num2 == 2020 {
				println(num1 * num2)
			}
		}
	}

	for start, num1 := range nums {
		for next, num2 := range nums[(start + 1):] {
			for _, num3 := range nums[(next + 1):] {
				if num1+num2+num3 == 2020 {
					println(num1 * num2 * num3)
				}
			}
		}
	}
}
