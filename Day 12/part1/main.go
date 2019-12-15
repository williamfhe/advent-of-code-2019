package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type moon struct {
	pos vec3
	vel vec3
}

func (m moon) energy() int {
	pot := m.pos.sum()
	kin := m.vel.sum()
	return pot * kin
}

type vec3 struct {
	x, y, z int
}

func (v *vec3) add(other vec3) {
	v.x += other.x
	v.y += other.y
	v.z += other.z
}

func (v vec3) sum() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func abs(s int) int {
	if s < 0 {
		s = -s
	}
	return s
}

func readInput(path string) []moon {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`(?m)<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
	var moons []moon
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalf("Invalid line in input : %s\n", line)
		}
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		z, _ := strconv.Atoi(matches[3])

		m := moon{pos: vec3{x, y, z}}
		moons = append(moons, m)
	}

	err = file.Close()

	if err != nil {
		log.Fatalf("failed closing file: %s", err)
	}

	return moons
}

func calculateEnergy(moons []moon) int {
	totalEnergy := 0
	for _, m := range moons {
		totalEnergy += m.energy()
	}
	return totalEnergy
}

func main() {
	moons := readInput("input.txt")
	n := 1000
	for step := 0; step < n; step++ {
		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				if moons[i].pos.x < moons[j].pos.x {
					moons[i].vel.x++
					moons[j].vel.x--
				}
				if moons[i].pos.x > moons[j].pos.x {
					moons[i].vel.x--
					moons[j].vel.x++
				}
				if moons[i].pos.y < moons[j].pos.y {
					moons[i].vel.y++
					moons[j].vel.y--
				}
				if moons[i].pos.y > moons[j].pos.y {
					moons[i].vel.y--
					moons[j].vel.y++
				}
				if moons[i].pos.z < moons[j].pos.z {
					moons[i].vel.z++
					moons[j].vel.z--
				}
				if moons[i].pos.z > moons[j].pos.z {
					moons[i].vel.z--
					moons[j].vel.z++
				}
			}
		}

		for i := range moons {
			moons[i].pos.add(moons[i].vel)
		}
	}

	totalEnergy := calculateEnergy(moons)
	fmt.Println(totalEnergy)
}
