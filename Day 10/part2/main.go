package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func sortAsteroids(asteroids map[vec2]bool, selectedAsteroid vec2) [][]vec2 {
	// map[DIRECTION] => map[DISTANCE] => ASTEROID (direction and distance from the selected asteroid)
	dirDistPosMap := make(map[vec2]map[int]vec2)
	angleDirectionMap := make(map[float64]vec2)
	for asteroid := range asteroids {
		if asteroid == selectedAsteroid {
			continue
		}
		direction := asteroid.sub(selectedAsteroid)
		distance := gcd(abs(direction.x), abs(direction.y))
		direction = direction.div(distance)

		// calculate the angle of the direction for sorting later
		dirAngle := math.Atan2(float64(direction.x), -float64(direction.y))
		dirAngle = math.Mod(dirAngle, math.Pi)
		if dirAngle < 0 {
			dirAngle += 2 * math.Pi
		}
		angleDirectionMap[dirAngle] = direction

		distPosMap, ok := dirDistPosMap[direction]
		if !ok {
			distPosMap = make(map[int]vec2)
		}
		distPosMap[distance] = asteroid
		dirDistPosMap[direction] = distPosMap
	}

	var dirAngles []float64
	for angle := range angleDirectionMap {
		dirAngles = append(dirAngles, angle)
	}

	// sort positions by angle and distance
	sort.Float64s(dirAngles)
	var angleOrderedPositions [][]vec2
	for _, angle := range dirAngles {

		dir := angleDirectionMap[angle]
		distPositionsMap := dirDistPosMap[dir]

		var distances []int
		for dist := range distPositionsMap {
			distances = append(distances, dist)
		}

		sort.Ints(distances)
		var positions []vec2
		for _, dist := range distances {
			positions = append(positions, distPositionsMap[dist])
		}

		angleOrderedPositions = append(angleOrderedPositions, positions)
	}

	return angleOrderedPositions
}

func main() {

	asteroids := readInput("input.txt")

	maxVisibilityAsteroid := vec2{22, 28} // got it from part1
	angleOrderedPositions := sortAsteroids(asteroids, maxVisibilityAsteroid)

	destroyed := 0
	i := 0
	var lastPosition vec2
	for len(angleOrderedPositions) > 0 {
		lastPosition = angleOrderedPositions[i][0]
		angleOrderedPositions[i] = angleOrderedPositions[i][1:]
		destroyed++

		if len(angleOrderedPositions[i]) > 1 {
			i = (i + 1) % len(angleOrderedPositions)
		} else {
			angleOrderedPositions = append(angleOrderedPositions[:i], angleOrderedPositions[i+1:]...)
			if i >= len(angleOrderedPositions) {
				i = 0
			}
		}

		if destroyed == 199 {
			break
		}
	}

	fmt.Println(lastPosition.x*100 + lastPosition.y)

}
