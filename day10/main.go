package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strings"
)

var directions = map[string][]image.Point{
	"|": {{0, -1}, {0, 1}},  // north, south (vertical)
	"-": {{1, 0}, {-1, 0}},  // east, west (horizontal)
	"L": {{0, -1}, {1, 0}},  // north, east (90-degree bend)
	"J": {{0, -1}, {-1, 0}}, // north, west (90-degree bend)
	"7": {{0, 1}, {-1, 0}},  // south, west (90-degree bend)
	"F": {{0, 1}, {1, 0}},   // south, east (90-degree bend)
}

//var visitedMap map[image.Point]bool

func main() {

	data, _ := os.ReadFile("day10/input.txt")
	answer := part_1(data)
	fmt.Println("Part 1: ", answer[0])
	fmt.Println("Part 2: ", answer[1])
	//part_2(data) // only works for actual input. for test input it gives 1 more values always

}
func part_1(file []byte) []int {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	width := len(lines[0]) - 1
	height := len(lines) - 1

	start := image.Point{}

	symbols := map[image.Point]string{}
	for y, line := range lines {
		l := strings.TrimSpace(line)
		for x, r := range l {

			symbols[image.Point{x, y}] = string(r)

			if r == 'S' {
				start = image.Point{x, y}

			}
		}
	}

	fmt.Println(start)
	symbols[start] = find_start_symbol(symbols, start, width, height)

	fmt.Println(symbols[start])
	visited := []image.Point{}
	area := 0
	current := start
	next := start
	for {
		if current != start && next == start {
			break
		}
		visited = append(visited, current)
		//visitedMap[current] = true
		current, next = next, start

		for _, d := range directions[symbols[current]] {
			if !slices.Contains(visited, current.Add(d)) {
				next = current.Add(d)
			}
		}

		area += current.X*next.Y - current.Y*next.X
	}

	p1 := len(visited) / 2

	p2 := (abs(area)-len(visited))/2 + 1

	return []int{p1, p2}
}

func find_start_symbol(symbols map[image.Point]string, start image.Point, width, height int) string {

	mapping := map[[4]bool]string{
		{true, false, true, false}: "|",
		{false, true, false, true}: "-",
		{true, true, false, false}: "L",
		{true, false, false, true}: "J",
		{false, false, true, true}: "7",
		{false, true, true, false}: "F",
	}
	north := start.Add(image.Point{0, -1})
	east := start.Add(image.Point{1, 0})
	south := start.Add(image.Point{0, 1})
	west := start.Add(image.Point{-1, 0})

	symbols[start] = mapping[[4]bool{
		strings.Contains("7F|", symbols[north]) && north.Y >= 0,
		strings.Contains("-7J", symbols[east]) && east.X <= width,
		strings.Contains("JL|", symbols[south]) && south.Y <= height,
		strings.Contains("-FL", symbols[west]) && west.X >= 0,
	}]
	//fmt.Println(symbols[start])

	return symbols[start]
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func part_2(file []byte) { // gives correct result for input but 1 extra for test results
// 	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
// 	width := len(lines[0]) - 1
// 	height := len(lines) - 1

// 	start := image.Point{}

// 	symbols := map[image.Point]string{}
// 	for y, line := range lines {
// 		l := strings.TrimSpace(line)
// 		for x, r := range l {

// 			symbols[image.Point{x, y}] = string(r)

// 			if r == 'S' {
// 				start = image.Point{x, y}

// 			}
// 		}
// 	}

// 	symbols[start] = find_start_symbol(symbols, start, width, height)
// 	inside_map := make(map[image.Point]bool)

// 	for y := 0; y <= height; y++ {
// 		inside := false
// 		for x := 0; x <= width; x++ {

// 			if _, ok := visitedMap[image.Point{x, y}]; ok {
// 				if symbols[image.Point{x, y}] == "|" || symbols[image.Point{x, y}] == "L" || symbols[image.Point{x, y}] == "J" {
// 					inside = !inside

// 				}
// 			} else if inside {
// 				inside_map[image.Point{x, y}] = inside

// 			}

// 		}

// 	}

// 	fmt.Println("inside map", len(inside_map))

// }
