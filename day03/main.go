package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	X, Y int
}

func (p Point) add_point(d Point) Point {
	return Point{p.X + d.X, p.Y + d.Y}
}

func get_symbols(file []byte) map[Point]string {
	symbols := map[Point]string{}

	for y, s := range strings.Fields(string(file)) {
		for x, r := range s {
			if r != '.' && !unicode.IsDigit(r) {
				symbols[Point{x, y}] = string(r)
			}
		}
	}
	return symbols
}

func get_engine_parts(file []byte, symbols map[Point]string) map[Point][]int {
	engine_parts := map[Point][]int{}
	re := regexp.MustCompile(`\d+`)
	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for y, s := range strings.Fields(string(file)) {
		for _, match := range re.FindAllStringIndex(s, -1) {
			keys := map[Point]bool{}
			for x := match[0]; x < match[1]; x++ {
				for _, d := range directions {
					keys[Point{x, y}.add_point(d)] = true
				}
			}

			n, _ := strconv.Atoi(s[match[0]:match[1]])
			for p := range keys {
				if _, exists := symbols[p]; exists {
					engine_parts[p] = append(engine_parts[p], n)
				}
			}
		}
	}
	return engine_parts

}
func part_1(file []byte) int {
	symbols := get_symbols(file)
	engine_parts := get_engine_parts(file, symbols)

	part_numbers := 0
	for _, values := range engine_parts {
		for _, value := range values {
			part_numbers += value
		}
	}
	return part_numbers
}

func part_2(file []byte) int {
	symbols := get_symbols(file)
	engine_parts := get_engine_parts(file, symbols)
	gear_ratio := 0
	for index, neighbors := range engine_parts {
		if symbols[index] == "*" && len(neighbors) == 2 {
			gear_ratio += neighbors[0] * neighbors[1]
		}
	}
	return gear_ratio
}

func main() {
	file, err := os.ReadFile("day03/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	sum_part1 := part_1(file)
	prod_part2 := part_2(file)

	fmt.Println(sum_part1)
	fmt.Println(prod_part2)
}
