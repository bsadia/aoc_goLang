package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day09/input.txt")

	answer := part_1(input)
	fmt.Println("Part 1: ", answer[0])
	fmt.Println("Part 2: ", answer[1])
}
func part_1(input []byte) []int {

	part1 := 0
	part2 := 0

	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		seq := []int{}
		for _, l := range strings.Split(strings.TrimSpace(line), " ") {
			s, _ := strconv.Atoi(l)
			seq = append(seq, s)
		}
		//fmt.Println(seq)
		part1 += get_next(seq)
		slices.Reverse(seq)
		part2 += get_next(seq)
	}

	return []int{part1, part2}
}
func get_next(v []int) int {
	val := []int{}
	//fmt.Println(s)
	for i := 0; i < len(v)-1; i++ {
		val = append(val, v[i+1]-v[i])
	}
	//fmt.Println(val)
	if !slices.ContainsFunc(val, func(x int) bool { return x != 0 }) {
		return v[len(v)-1]
	}
	return v[len(v)-1] + get_next(val)
}
