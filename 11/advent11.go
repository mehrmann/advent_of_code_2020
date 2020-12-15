package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type Space struct {
	isSeat     bool
	isOccupied bool
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func printSeatMap(spaces []Space, stride int) {
	for i, x := range spaces {
		if x.isSeat == false {
			fmt.Printf(".")
		} else if x.isOccupied == true {
			fmt.Printf("#")
		} else {
			fmt.Printf("L")
		}
		if (i+1)%stride == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func calcNextState(spaces []Space, x int, y int, maxX int, maxY int) Space {
	space := spaces[y*maxX+x]
	if space.isSeat == false {
		return Space{false, false}
	} else {
		//check 8 directions
		countOccupied := 0
		for cy := Max(y-1, 0); cy < Min(y+2, maxY); cy++ {
			for cx := Max(x-1, 0); cx < Min(x+2, maxX); cx++ {
				//println(cx, cy, spaces[cx+cy*maxX].isOccupied)
				if cx != x || cy != y {

					if spaces[cx+cy*maxX].isOccupied {
						countOccupied = countOccupied + 1
					}
				}
			}
		}
		//println(x, y, countOccupied)
		if space.isOccupied == true && countOccupied >= 4 {
			return Space{true, false}
		} else if space.isOccupied == false && countOccupied == 0 {
			return Space{true, true}
		}
	}
	return space
}

func calcNextState2(spaces []Space, x int, y int, maxX int, maxY int) Space {
	space := spaces[y*maxX+x]
	if space.isSeat == false {
		return Space{false, false}
	} else {
		//check 8 directions
		countOccupied := 0
		//look right
		for cx := x; cx < maxX; cx++ {
			if cx != x {
				seen := spaces[cx+y*maxX]
				if seen.isOccupied {
					//println(x, y, ">", cx, y)
					countOccupied = countOccupied + 1
					break
				} else if seen.isSeat {
					break
				}
			}
		}
		//left
		for cx := x; cx >= 0; cx-- {
			if cx != x {
				seen := spaces[cx+y*maxX]
				if seen.isOccupied {
					//println(x, y, "<", cx, y)
					countOccupied = countOccupied + 1
					break
				} else if seen.isSeat {
					break
				}
			}
		}
		//down
		for cy := y; cy < maxY; cy++ {
			if cy != y {
				seen := spaces[x+cy*maxX]
				if seen.isOccupied {
					//println(x, y, "d", x, cy)
					countOccupied = countOccupied + 1
					break
				} else if seen.isSeat {
					break
				}
			}
		}
		//up
		for cy := y; cy >= 0; cy-- {
			if cy != y {
				seen := spaces[x+cy*maxX]
				if seen.isOccupied {
					//println(x, y, "u", x, cy)
					countOccupied = countOccupied + 1
					break
				} else if seen.isSeat {
					break
				}
			}
		}
		for i := 0; i <= Max(maxX, maxY); i++ {
			cx := x + i
			cy := y + i
			if cx >= 0 && cx < maxX && cy >= 0 && cy < maxY && (cx != x || cy != y) {
				seen := spaces[cx+cy*maxX]
				if seen.isOccupied {
					//println(x, y, i, cx, cy)
					countOccupied = countOccupied + 1
					break
				} else if seen.isSeat {
					break
				}
			}
		}
		for i := 0; i <= Max(maxX, maxY); i++ {
			cx := x - i
			cy := y + i
			if cx >= 0 && cx < maxX && cy >= 0 && cy < maxY && (cx != x || cy != y) {
				seen := spaces[cx+cy*maxX]
				if seen.isOccupied {
					//println(x, y, i, cx, cy)
					countOccupied = countOccupied + 1
					break
				} else if seen.isSeat {
					break
				}
			}
		}
		for i := 0; i <= Max(maxX, maxY); i++ {
			cx := x + i
			cy := y - i
			if cx >= 0 && cx < maxX && cy >= 0 && cy < maxY && (cx != x || cy != y) {
				seen := spaces[cx+cy*maxX]
				if seen.isOccupied {
					//println(x, y, i, cx, cy)
					countOccupied = countOccupied + 1
					break
				} else if seen.isSeat {
					break
				}
			}
		}
		for i := 0; i <= Max(maxX, maxY); i++ {
			cx := x - i
			cy := y - i
			if cx >= 0 && cx < maxX && cy >= 0 && cy < maxY && (cx != x || cy != y) {

				seen := spaces[cx+cy*maxX]
				if seen.isOccupied {
					//println(x, y, i, cx, cy)
					countOccupied = countOccupied + 1
					break
				} else if seen.isSeat {
					break
				}
			}
		}

		//println("-->", x, y, countOccupied)
		if space.isOccupied == true && countOccupied >= 5 {
			return Space{true, false}
		} else if space.isOccupied == false && countOccupied == 0 {
			return Space{true, true}
		}
	}
	return space
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	stride := len(lines[0])
	spaces := make([]Space, len(lines)*stride, len(lines)*stride)
	spaces2 := make([]Space, len(lines)*stride, len(lines)*stride)

	for y, l := range lines {
		for x, s := range l {
			seat := s == 'L'
			spaces[y*stride+x] = Space{seat, false}
			spaces2[y*stride+x] = Space{seat, false}
		}
	}

	//printSeatMap(spaces, stride)

	for {
		nextState := make([]Space, len(lines)*stride, len(lines)*stride)
		for y := 0; y < len(lines); y++ {
			for x := 0; x < stride; x++ {
				//fmt.Printf("checking %d, %d\n", x, y)
				nextState[y*stride+x] = calcNextState(spaces, x, y, stride, len(lines))
			}
		}
		//printSeatMap(nextState, stride)
		if reflect.DeepEqual(nextState, spaces) {
			break
		}
		spaces = nextState
	}

	part1 := 0
	for _, x := range spaces {
		if x.isOccupied == true {
			part1++
		}
	}

	println(part1)

	for {
		nextState := make([]Space, len(lines)*stride, len(lines)*stride)
		for y := 0; y < len(lines); y++ {
			for x := 0; x < stride; x++ {
				//fmt.Printf("checking %d, %d\n", x, y)
				nextState[y*stride+x] = calcNextState2(spaces2, x, y, stride, len(lines))
			}
		}
		//printSeatMap(nextState, stride)
		if reflect.DeepEqual(nextState, spaces2) {
			break
		}
		spaces2 = nextState
	}

	part2 := 0
	for _, x := range spaces2 {
		if x.isOccupied == true {
			part2++
		}
	}

	println(part2)

}
