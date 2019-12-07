package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type objectOrbit struct {
	orbiting  string
	orbitedBy []string
}

func readInput(path string) map[string]objectOrbit {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	orbitMap := make(map[string]objectOrbit)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		scanner.Text()

		line := scanner.Text()
		splittedLine := strings.Split(line, ")")
		object := splittedLine[0]
		orbitingObject := splittedLine[1]

		objectOrbit := orbitMap[object]
		objectOrbit.orbitedBy = append(objectOrbit.orbitedBy, orbitingObject)
		orbitMap[object] = objectOrbit

		orbitinObjectOrbit := orbitMap[orbitingObject]
		orbitinObjectOrbit.orbiting = object
		orbitMap[orbitingObject] = orbitinObjectOrbit
	}

	err = file.Close()

	if err != nil {
		log.Fatalf("failed closing file: %s", err)
	}

	return orbitMap
}

func exploreMap(orbitMap map[string]objectOrbit, fromObject, currentObject, stopObject string) int {
	var exploredDepth int

	if currentObject == stopObject {
		// we have reached our destination
		return 1
	}

	currentObjectOrbits := orbitMap[currentObject]

	if currentObjectOrbits.orbiting != fromObject {
		// verify the "parent", the one the current object is orbiting
		exploredDepth += exploreMap(orbitMap, currentObject, currentObjectOrbits.orbiting, stopObject)
		if exploredDepth > 0 {
			return exploredDepth + 1
		}
	}

	for _, orbitingObject := range currentObjectOrbits.orbitedBy {
		if orbitingObject == fromObject {
			continue
		}

		exploredDepth += exploreMap(orbitMap, currentObject, orbitingObject, stopObject)

		if exploredDepth > 0 {
			return exploredDepth + 1
		}
	}

	return 0
}

func main() {
	orbitMap := readInput("input.txt")

	startObject := orbitMap["YOU"].orbiting
	stopObject := orbitMap["SAN"].orbiting

	minPath := exploreMap(orbitMap, "YOU", startObject, stopObject) - 1

	fmt.Println(minPath)
}
