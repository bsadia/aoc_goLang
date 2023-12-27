package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {

	input, _ := os.ReadFile("day11/input.txt")

	fmt.Println("Part 1: ", part_1(input, 2))
	fmt.Println("Part 2: ", part_1(input, 1000000))

}

func part_1(input []byte, exp int) int {

	lines := strings.Fields(string(input))
	dist := 0

	galaxies := make([]image.Point, 0, len(lines))

	dy := 0
	for y, s := range lines {
		if !strings.Contains(s, "#") {
			dy += exp - 1
		}

		dx := 0
		for x, r := range s {
			containsHash := false
			for _, s := range lines {
				ch := s[x]
				if ch == '#' {
					containsHash = true
					break
				}
			}
			if !containsHash {
				dx += exp - 1
			}
			if r == '#' {

				for _, g := range galaxies {
					dist += abs(x+dx-g.X) + abs(y+dy-g.Y)
				}
				galaxies = append(galaxies, image.Point{x + dx, y + dy})
			}

		}
	}

	//fmt.Println("answer", dist)
	return dist

}
func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
