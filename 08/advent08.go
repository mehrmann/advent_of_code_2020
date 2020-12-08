package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type OpCode string

const (
	NOP  = "nop"
	JMP  = "jmp"
	ACC  = "acc"
	TERM = "term"
)

type Instruction struct {
	opcode OpCode
	value  int
}

func runWithMods(listing []Instruction, pcToMod int, modInstruction Instruction) (bool, int) {
	acc := 0
	pc := 0
	visited := make(map[int]bool)
	for {
		if pc > len(listing) {
			return false, acc
		}
		if _, ok := visited[pc]; ok {
			return false, acc
		}
		visited[pc] = true

		instruction := listing[pc]
		if pc == pcToMod {
			instruction = modInstruction
		}
		//fmt.Printf("%s %d\n", instruction.opcode, instruction.value)
		if instruction.opcode == NOP {
			pc++
		} else if instruction.opcode == JMP {
			pc = pc + instruction.value
		} else if instruction.opcode == ACC {
			acc = acc + instruction.value
			pc++
		} else if instruction.opcode == TERM {
			return true, acc
		}
	}
}

func makeInstruction(l string) Instruction {
	s := strings.Split(l, " ")

	var opcode = OpCode(s[0])

	value, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}
	return Instruction{opcode, value}
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	listing := make([]Instruction, 0)
	for _, l := range lines {

		i := makeInstruction(l)
		listing = append(listing, i)
	}
	listing = append(listing, Instruction{OpCode("term"), 0})

	acc := 0
	pc := 0

	visited := make(map[int]bool)
	for {
		if _, ok := visited[pc]; ok {
			break
		}
		visited[pc] = true
		instruction := listing[pc]
		//fmt.Printf("%s %d\n", instruction.opcode, instruction.value)
		if instruction.opcode == NOP {
			pc++
		} else if instruction.opcode == JMP {
			pc = pc + instruction.value
		} else if instruction.opcode == ACC {
			acc = acc + instruction.value
			pc++
		}
	}
	println(acc)

	for i := 0; i < len(listing); i++ {
		instruction := listing[i]

		if instruction.opcode == JMP {
			didTerm, acc := runWithMods(listing, i, Instruction{NOP, instruction.value})
			if didTerm {
				println(acc)
			}
		} else if instruction.opcode == NOP {
			didTerm, acc := runWithMods(listing, i, Instruction{JMP, instruction.value})
			if didTerm {
				println(acc)
			}
		}
	}

}
