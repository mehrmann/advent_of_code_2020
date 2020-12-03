package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Entry struct {
	min  int
	max  int
	c    string
	pass string
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	validCount := 0
	validCount2 := 0

	for _, l := range lines {
		//var data = Entry{}
		//println(l)
		var s = regexp.MustCompile("[- :]").Split(l, 4)

		min, _ := strconv.Atoi(s[0])
		max, _ := strconv.Atoi(s[1])
		c := s[2]
		pass := s[3]

		letterCount := strings.Count(pass, c)
		if letterCount >= min && letterCount <= max {
			validCount++
		}

		a := pass[min] == c[0]
		b := pass[max] == c[0]
		if (a || b) && !(a && b) {
			validCount2++
		}

	}
	println(validCount)
	println(validCount2)

}
