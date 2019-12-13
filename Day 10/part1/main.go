package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type vec2 struct {
	x, y int
}

func (v vec2) sub(other vec2) vec2 {
	dx := v.x - other.x
	dy := v.y - other.y

	return vec2{dx, dy}
}

func (v vec2) div(s int) vec2 {
	dx := v.x / s
	dy := v.y / s

	return vec2{dx, dy}
}

func gcd(a, b int) int {
	var tmp int
	if b > a {
		tmp = a
		a = b
		b = tmp
	}

	for b != 0 {
		tmp = a % b
		a = b
		b = tmp
	}

	return a
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func readInput(path string) map[vec2]bool {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	asteroids := make(map[vec2]bool)

	y := 0
	for scanner.Scan() {
		inputRow := scanner.Text()
		for x, char := range []rune(inputRow) {
			if char == '.' {
				continue
			}
			pos := vec2{x, y}
			asteroids[pos] = false
		}
		y++
	}

	err = file.Close()

	if err != nil {
		log.Fatalf("failed closing file: %s", err)
	}

	return asteroids
}

func getVisibleAsteroids(asteroids map[vec2]bool, selectedAsteroid vec2) int {
	visibleAsteroids := make(map[vec2]bool)
	for asteroid := range asteroids {
		if asteroid == selectedAsteroid {
			continue
		}

		direction := asteroid.sub(selectedAsteroid)
		s := gcd(abs(direction.x), abs(direction.y))
		direction = direction.div(s)

		visibleAsteroids[direction] = false
	}

	return len(visibleAsteroids)
}

func main() {

	asteroids := readInput("input.txt")
	maxVisible := -1
	var maxVisibilityAsteroid vec2
	for asteroid := range asteroids {
		visibleAsteroids := getVisibleAsteroids(asteroids, asteroid)
		if visibleAsteroids > maxVisible {
			maxVisible = visibleAsteroids
			maxVisibilityAsteroid = asteroid
		}
	}

	fmt.Println(maxVisible, maxVisibilityAsteroid)
}
