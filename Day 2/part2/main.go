package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const input = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,19,6,23,2,13,23,27,1,27,13,31,1,9,31,35,1,35,9,39,1,39,5,43,2,6,43,47,1,47,6,51,2,51,9,55,2,55,13,59,1,59,6,63,1,10,63,67,2,67,9,71,2,6,71,75,1,75,5,79,2,79,10,83,1,5,83,87,2,9,87,91,1,5,91,95,2,13,95,99,1,99,10,103,1,103,2,107,1,107,6,0,99,2,14,0,0"

const expectedOutput = 19690720

var opcodes []int

func init() {
	splittedInput := strings.Split(input, ",")

	for _, code := range splittedInput {
		opcode, err := strconv.Atoi(code)
		if err != nil {
			log.Fatalf("Could not convert to int: %s\n", err)
		}
		opcodes = append(opcodes, opcode)
	}
}

func tryNounVerb(noun, verb int) int {

	codes := make([]int, len(opcodes))
	copy(codes, opcodes)

	codes[1] = noun
	codes[2] = verb

	index := 0
	for codes[index] != 99 {
		index1 := codes[index+1]
		index2 := codes[index+2]
		index3 := codes[index+3]

		switch codes[index] {
		case 1:
			codes[index3] = codes[index1] + codes[index2]
		case 2:
			codes[index3] = codes[index1] * codes[index2]
		default:
			log.Fatalf("Unknown opcode %d at %d\n", codes[index], index)
		}

		index += 4
	}

	return codes[0]
}

func main() {

	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			output := tryNounVerb(noun, verb)
			if output == expectedOutput {
				fmt.Println(noun*100 + verb)
			}
		}
	}
}
