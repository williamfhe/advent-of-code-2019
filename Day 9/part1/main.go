package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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
	prog := newProgram(instructions, []int{1})
	output := prog.run()
	fmt.Println(output[0])
}
