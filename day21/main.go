package main

import (
	"container/list"
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {

	data, _ := os.ReadFile("day21/input.txt")
	answer := solution(data, 64)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", answer[0], answer[1])

}

func solution(data []byte, total_steps int) []int {

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	q := list.New()

	symbols := map[image.Point]string{}
	for y, line := range lines {
		l := strings.TrimSpace(line)
		//fmt.Println(len(s))
		for x, r := range l {

			symbols[image.Point{x, y}] = string(r)

			if r == 'S' {

				q.PushBack(image.Point{x, y})

			}
		}
	}

	steps := 0
	total := 26501365
	part_1 := 0
	part_2 := 0

	directions := []image.Point{
		{-1, 0}, {0, -1}, {0, 1}, {1, 0},
	}

	polynomial := make([]int, 0)
	for steps < total {
		new_Q := list.New()
		visited := make(map[image.Point]bool)

		for q.Len() > 0 {
			element := q.Front()

			val := element.Value.(image.Point)
			q.Remove(element)
			x, y := val.X, val.Y

			for _, d := range directions {
				newPos := image.Point{x + d.X, y + d.Y}

				ref := image.Point{((newPos.Y % len(lines)) + len(lines)) % len(lines), ((newPos.X % len(lines)) + len(lines)) % len(lines)}

				if symbols[ref] != "#" {
					if _, ok := visited[newPos]; !ok {
						visited[newPos] = true
						new_Q.PushBack(newPos)
					}
				}

			}

		}
		steps += 1
		q = new_Q //update queue
		if steps%(len(lines)) == total%len(lines) {
			polynomial = append(polynomial, len(visited))

			if len(polynomial) == 3 {
				p0 := polynomial[0]
				p1 := polynomial[1] - polynomial[0]
				p2 := polynomial[2] - polynomial[1]

				part_2 = p0 + (p1 * (total / len(lines))) + ((total/len(lines))*((total/len(lines))-1)/2)*(p2-p1)
				break
			}
		}
		if steps == total_steps {
			part_1 = len(visited)
		}

	}

	return []int{part_1, part_2}
}
