package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	up    = vec2{0, +1}
	right = vec2{+1, 0}
	down  = vec2{0, -1}
	left  = vec2{-1, 0}
)

type vec2 struct {
	x, y int
}

func (v *vec2) dist() int {
	x := v.x
	if x < 0 {
		x = -x
	}

	y := v.y
	if y < 0 {
		y = -y
	}

	return x + y
}

func readInput(path string) [][]string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var wires [][]string

	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.Split(line, ",")
		wires = append(wires, splittedLine)
	}

	err = file.Close()

	if err != nil {
		log.Fatalf("failed closing file: %s", err)
	}

	return wires
}

func parseMovement(movement string) (vec2, int) {
	movementRunes := []rune(movement)
	direction := movementRunes[0]
	movementLength, _ := strconv.Atoi(string(movementRunes[1:]))

	var directionVec vec2
	switch direction {
	case 'U':
		directionVec = up
	case 'R':
		directionVec = right
	case 'D':
		directionVec = down
	case 'L':
		directionVec = left
	default:
		log.Fatalf("Unknown direction : %v\n", direction)
	}

	return directionVec, movementLength
}

func main() {
	wires := readInput("input.txt")

	// map of the coordinates where a wire pass
	wireAtPos := make(map[vec2]int)

	currentPos := vec2{0, 0}
	currentSteps := 0

	// loop over each movement of the first wire
	for _, movement := range wires[0] {

		directionVec, movementLength := parseMovement(movement)

		for i := 0; i < movementLength; i++ {
			currentSteps++
			currentPos = vec2{currentPos.x + directionVec.x, currentPos.y + directionVec.y}
			wireAtPos[currentPos] = currentSteps
		}

	}

	minDistance := 1<<63 - 1 // max int
	currentPos = vec2{0, 0}
	currentSteps = 0

	// loop over each movement of the second wire
	for _, movement := range wires[1] {

		directionVec, movementLength := parseMovement(movement)

		for i := 0; i < movementLength; i++ {
			currentPos = vec2{currentPos.x + directionVec.x, currentPos.y + directionVec.y}
			currentSteps++
			if firstWireStepsAt, ok := wireAtPos[currentPos]; ok {
				if minDistance > currentSteps+firstWireStepsAt {
					minDistance = currentSteps + firstWireStepsAt
				}
			}
		}

	}

	fmt.Println(minDistance)
}
