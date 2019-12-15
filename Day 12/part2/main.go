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

func lcm(a, b int) int {
	return abs(a*b) / gcd(a, b)
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

func main() {
	moons := readInput("input.txt")

	var baseX, baseY, baseZ string
	for i := range moons {
		baseX += fmt.Sprintf("%d%d", moons[i].pos.x, moons[i].vel.x)
		baseY += fmt.Sprintf("%d%d", moons[i].pos.y, moons[i].vel.y)
		baseZ += fmt.Sprintf("%d%d", moons[i].pos.z, moons[i].vel.z)
	}

	seenStep := make(map[byte]int)
	step := 0
	for {
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
		step++

		var hashX, hashY, hashZ string
		for i := range moons {
			moons[i].pos.add(moons[i].vel)
			hashX += fmt.Sprintf("%d%d", moons[i].pos.x, moons[i].vel.x)
			hashY += fmt.Sprintf("%d%d", moons[i].pos.y, moons[i].vel.y)
			hashZ += fmt.Sprintf("%d%d", moons[i].pos.z, moons[i].vel.z)
		}

		if baseX == hashX && seenStep['x'] == 0 {
			seenStep['x'] = step
		}

		if baseY == hashY && seenStep['y'] == 0 {
			seenStep['y'] = step
		}

		if baseZ == hashZ && seenStep['z'] == 0 {
			seenStep['z'] = step
		}

		if len(seenStep) == 3 {
			break
		}
	}

	fmt.Println(lcm(lcm(seenStep['x'], seenStep['y']), seenStep['z']))

}
