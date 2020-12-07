package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type ContainedBag struct {
	count int
	name  string
}

func arrayContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sumUp(s string, m map[string][]ContainedBag) int {
	res := 0
	for _, b := range m[s] {
		res += b.count
		res += sumUp(b.name, m) * b.count
	}
	return res
}

func recurseCount(search string, m map[string][]string, dups map[string]struct{}) (int, map[string]struct{}) {
	result := 0
	for _, i := range m[search] {
		if _, ok := dups[i]; !ok {

			dups[i] = struct{}{}
			result++

			r, newdups := recurseCount(i, m, dups)
			result += r
			for k, v := range newdups {
				dups[k] = v
			}
		}

	}
	return result, dups
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	validCount2 := 0

	reverse := make(map[string][]string)
	forward := make(map[string][]ContainedBag)

	for _, l := range lines {
		var s = regexp.MustCompile(" contain ").Split(l, 2)
		bag := regexp.MustCompile("(.*) bags").FindStringSubmatch(s[0])[1]

		for _, e := range regexp.MustCompile(" ?bags?[.,]+").Split(s[1], -1) {
			if len(e) > 0 && !strings.HasPrefix(e, "no other") {

				reduced := regexp.MustCompile(" ?([0-9]+) (.*)").FindStringSubmatch(e)
				num, _ := strconv.Atoi(reduced[1])
				contained := reduced[2]
				forward[bag] = append(forward[bag], ContainedBag{num, contained})

				if c, ok := reverse[contained]; ok {
					reverse[contained] = append(c, bag)
				} else {
					reverse[contained] = []string{bag}
				}
			}
		}
	}

	res, _ := recurseCount("shiny gold", reverse, make(map[string]struct{}))
	println(res)

	validCount2 = sumUp("shiny gold", forward)
	println(validCount2)

}
