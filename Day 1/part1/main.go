package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	err = file.Close()

	if err != nil {
		log.Fatalf("failed closing file: %s", err)
	}

	return fileLines
}

func main() {
	fileLines := readFile("input.txt")
	var total int

	for _, line := range fileLines {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Could not convert to int: %s\n", err)
		}
		total += (i / 3) - 2
	}

	fmt.Println(total)
}
