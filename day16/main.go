package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type State struct {
	Pos image.Point
	Dir image.Point
}

var (
	U = image.Point{0, -1}
	R = image.Point{1, 0}
	D = image.Point{0, 1}
	L = image.Point{-1, 0}
)

var directions = map[rune]map[image.Point][]image.Point{
	'.':  {U: {U}, R: {R}, D: {D}, L: {L}},
	'/':  {U: {R}, R: {U}, D: {L}, L: {D}},
	'\\': {U: {L}, R: {D}, D: {R}, L: {U}},
	'|':  {U: {U}, R: {U, D}, D: {D}, L: {U, D}},
	'-':  {U: {L, R}, R: {R}, D: {L, R}, L: {L}},
}

func main() {
	input, _ := os.ReadFile("day16/input.txt")
	fmt.Println(part_1(input))
	fmt.Println(part_2(input))

}
func part_1(file []byte) int {
	split := strings.Fields(string(file))

	grid := map[image.Point]rune{}
	for y, s := range split {

		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}
	answer := bfs(State{image.Point{0, 0}, R}, grid)
	return answer
}

func part_2(file []byte) int {
	split := strings.Fields(string(file))

	grid := map[image.Point]rune{}
	border := []State{}
	for y, s := range split {
		border = append(border, State{image.Point{0, y}, R}, State{image.Point{len(s) - 1, y}, L})
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}
	for x := range split[0] {
		border = append(border, State{image.Point{x, 0}, D}, State{image.Point{x, len(split) - 1}, U})
	}

	part2 := 0
	for _, s := range border {
		part2 = max(part2, bfs(s, grid))
	}

	return part2
}

func bfs(start State, grid map[image.Point]rune) int {
	energized := map[image.Point]struct{}{}
	queue := []State{start}
	seen := map[State]struct{}{start: {}}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		for _, d := range directions[grid[state.Pos]][state.Dir] {
			energized[state.Pos] = struct{}{}
			next := State{state.Pos.Add(d), d}

			if _, ok := seen[next]; !ok {
				seen[next] = struct{}{}
				queue = append(queue, next)

			}
		}
	}

	return len(energized)
}
