package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func arrayContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func contains(s map[string]string, e string) bool {
	_, ok := s[e]
	return ok
}

func validateSimple(fields map[string]string) bool {
	return contains(fields, "byr") &&
		contains(fields, "iyr") &&
		contains(fields, "eyr") &&
		contains(fields, "hgt") &&
		contains(fields, "hcl") &&
		contains(fields, "ecl") &&
		contains(fields, "pid")

}
func validate(fields map[string]string) bool {
	valid := validateSimple(fields)

	if byr, ok := fields["byr"]; ok {
		byr, err := strconv.Atoi(byr)
		if err != nil {
			valid = false
		} else {
			if byr >= 1920 && byr <= 2002 {
				println("valid byr:", byr)
			} else {
				println("invalid byr:", byr)
				valid = false
			}

		}
	}

	if iyr, ok := fields["iyr"]; ok {
		iyr, err := strconv.Atoi(iyr)
		if err != nil {

			valid = false
		} else {
			if iyr >= 2010 && iyr <= 2020 {
				println("valid iyr:", iyr)
			} else {
				println("invalid iyr:", iyr)
				valid = false
			}
		}
	}

	if eyr, ok := fields["eyr"]; ok {
		eyr, err := strconv.Atoi(eyr)
		if err != nil {
			valid = false
		} else {
			if eyr >= 2020 && eyr <= 2030 {
				println("valid eyr:", eyr)
			} else {
				println("invalid eyr:", eyr)
				valid = false
			}

		}
	}

	if hgt, ok := fields["hgt"]; ok {
		var h int
		_, err := fmt.Sscanf(hgt, "%din", &h)
		if err != nil {
			_, err := fmt.Sscanf(hgt, "%dcm", &h)

			if err != nil {
				println("invalid hgt", hgt)
				valid = false
			} else {
				if h >= 150 && h <= 193 {
					println("valid hgt:", hgt)
				} else {
					println("invalid hgt", hgt)
					valid = false
				}

			}
		} else {
			if h >= 59 && h <= 76 {
				println("valid hgt:", hgt)
			} else {
				println("invalid hgt", hgt)
				valid = false
			}

		}
	}

	if hcl, ok := fields["hcl"]; ok {
		m := regexp.MustCompile("^#[0-9a-f]{6}$")
		if !m.MatchString(hcl) {
			println("invalid hcl", hcl)
			valid = false
		} else {
			println("valid hcl", hcl)
		}
	}

	if ecl, ok := fields["ecl"]; ok {
		validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		if !arrayContains(validEyeColors, ecl) {
			println("invalid ecl", ecl)
			valid = false
		} else {
			println("valid ecl", ecl)
		}
	}

	if pid, ok := fields["pid"]; ok {
		m := regexp.MustCompile("^[0-9]{9}$")
		if !m.MatchString(pid) {
			println("invalid pid", pid)
			valid = false
		} else {
			println("valid pid", pid)
		}
	}

	return valid
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	result := 0
	result_2 := 0

	fields := make(map[string]string)
	for _, l := range lines {
		if len(l) == 0 {
			//validate
			if validateSimple(fields) {
				result++
			}
			if validate(fields) {
				result_2++
			}
			fields = make(map[string]string)
		} else {
			entries := strings.Split(l, " ")
			for _, e := range entries {
				keyvalue := strings.Split(e, ":")
				fields[keyvalue[0]] = keyvalue[1]
			}
		}
	}

	//last entry scan...
	if validateSimple(fields) {
		result++
	}
	if validate(fields) {
		result_2++
	}

	println(result)
	println(result_2)

}
