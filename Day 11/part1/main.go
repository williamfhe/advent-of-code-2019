package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type vec2 struct {
	x, y int
}

func (v *vec2) add(other vec2) {
	v.x += other.x
	v.y += other.y
}

var directions = []vec2{vec2{0, 1}, vec2{1, 0}, vec2{0, -1}, vec2{-1, 0}}

func readInput(path string) map[int]int {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file :%s\n", err)
	}

	input := make(map[int]int)
	splittedData := strings.Split(string(data), ",")

	for index, valueStr := range splittedData {
		valueInt, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Fatalf("Could not convert to int: %s\n", err)
		}
		input[index] = valueInt
	}

	return input
}

func main() {
	instructions := readInput("input.txt")

	directionIndex := 0
	programInput := 0
	robotPosition := vec2{0, 0}
	paintedPanels := make(map[vec2]int)
	prog := newProgram(instructions, nil)
	for !prog.stopped {
		prog.addInputs([]int{programInput})
		output := prog.run()

		if len(output) > 2 || output[0] < 0 || output[0] > 1 || output[1] < 0 || output[1] > 1 {
			log.Fatalf("Invalid output from program: %v\n", output)
		}

		paintedPanels[robotPosition] = output[0]
		if output[1] == 0 {
			// turn left
			directionIndex = (directionIndex - 1 + 4) % 4
		} else {
			// turn right
			directionIndex = (directionIndex + 1 + 4) % 4
		}

		robotPosition.add(directions[directionIndex])
		programInput = paintedPanels[robotPosition]
	}

	fmt.Println(len(paintedPanels))
}
