package main

import (
	"io/ioutil"
	"strings"
)

func checkMovement(lines []string, xinc int, yinc int) int {
	stride := len(lines[0])
	trees := 0

	x := 0
	for y := yinc; y < len(lines); y = y + yinc {
		x = (x + xinc) % stride
		if lines[y][x] == '#' {
			trees++
		}
	}
	return trees
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	println(checkMovement(lines, 3, 1))
	println(checkMovement(lines, 1, 1) * checkMovement(lines, 3, 1) * checkMovement(lines, 5, 1) * checkMovement(lines, 7, 1) * checkMovement(lines, 1, 2))

}
