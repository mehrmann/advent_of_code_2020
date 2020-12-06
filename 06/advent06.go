package main

import (
	"io/ioutil"
	"strings"
)

func calcSame(fields map[rune]int, people int) int {
	count := 0
	for _, v := range fields {
		if v == people {
			count++
		}
	}
	return count
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	result := 0
	result2 := 0
	people := 0

	fields := make(map[rune]int)
	for _, l := range lines {
		if len(l) == 0 {
			result += len(fields)
			result2 += calcSame(fields, people)
			fields = make(map[rune]int)
			people = 0
		} else {
			people++
			for _, e := range l {
				if val, ok := fields[e]; ok {
					fields[e] = val + 1
				} else {
					fields[e] = 1
				}
			}
		}
	}

	result += len(fields)
	result2 += calcSame(fields, people)

	println(result)
	println(result2)

}
