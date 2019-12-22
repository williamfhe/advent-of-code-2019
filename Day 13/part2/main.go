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
	instructions[0] = 2

	prog, inputChan, outputChan := newProgram(instructions)
	go prog.run()

	var ballPos, paddlePos vec2
	segmentPos := vec2{-1, 0}

	score := -1

	var outputs []int
	for msg := range outputChan {

		if msg.msgType == needInputMessage {
			inputChan <- ballPos.x - paddlePos.x
			continue
		}

		outputs = append(outputs, msg.data)
		if len(outputs) == 3 {
			pos := vec2{outputs[0], outputs[1]}
			if pos == segmentPos {
				score = outputs[2]
			}
			if outputs[2] == 3 {
				paddlePos = pos
			}

			if outputs[2] == 4 {
				ballPos = pos
			}

			outputs = nil
		}
	}

	fmt.Println(score)

}
