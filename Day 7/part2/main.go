package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(path string) []int {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file :%s\n", err)
	}

	var input []int

	splittedData := strings.Split(string(data), ",")

	for _, valueStr := range splittedData {
		valueInt, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Fatalf("Could not convert to int: %s\n", err)
		}
		input = append(input, valueInt)
	}

	return input
}

func copySlice(a []int) []int {
	cp := make([]int, len(a))
	copy(cp, a)
	return cp
}

func swap(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func permutations(a []int) <-chan []int {
	if len(a) < 2 {
		return nil
	}

	permChan := make(chan []int)

	go func(a []int) {
		permChan <- copySlice(a)

		c := make([]int, len(a))
		i := 0
		for i < len(a) {
			if c[i] < i {
				if i%2 == 0 {
					swap(a, 0, i)
				} else {
					swap(a, c[i], i)
				}
				permChan <- copySlice(a)
				c[i]++
				i = 0
			} else {
				c[i] = 0
				i++
			}
		}
		close(permChan)
	}(a)
	return permChan
}

func amplifierExecution(instructions []int, phase []int) int {
	outputStack := []int{0}

	var amplifierPrograms []*program
	for i := 0; i < len(phase); i++ {
		prog := newProgram(copySlice(instructions), []int{phase[i]})
		amplifierPrograms = append(amplifierPrograms, prog)
	}

	i := 0
	for {
		prog := amplifierPrograms[i%len(phase)]
		if prog.stopped {
			break
		}
		prog.addInputs(outputStack)
		outputStack = prog.run()
		i++
	}

	return outputStack[0]
}

func main() {
	instructions := readInput("input.txt")

	maxSignal := -1 << 63
	permChan := permutations([]int{5, 6, 7, 8, 9})
	for perm := range permChan {
		signal := amplifierExecution(instructions, perm)
		if signal > maxSignal {
			maxSignal = signal
		}
	}
	fmt.Println(maxSignal)
}
