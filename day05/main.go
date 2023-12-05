package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func get_input_map(lines []string) map[int][][]int {
	all_maps := make(map[int][][]int)

	for i := 1; i < len(lines); i++ {

		single_map := [][]int{}
		for _, line := range strings.Split(lines[i], "\n") {
			matches := regexp.MustCompile(`-?\d+`).FindAllString(line, -1)
			map_values := []int{}

			for _, match := range matches {
				num, err := strconv.Atoi(match)
				if err == nil {
					map_values = append(map_values, num)
				}
			}
			if len(map_values) > 0 {
				single_map = append(single_map, map_values)
			}
		}

		all_maps[i] = single_map

	}

	return all_maps
}

func process_input(file []byte) ([]int, map[int][][]int) {

	lines := strings.Split(string(file), "\n\n")
	seedsArray := strings.TrimSpace((strings.Split(lines[0], ":"))[1])

	seeds := []int{}
	matches := regexp.MustCompile(`-?\d+`).FindAllString(seedsArray, -1)
	for _, match := range matches {
		seed, _ := strconv.Atoi(string(match))
		seeds = append(seeds, seed)
	}

	map_values := get_input_map(lines)
	return seeds, map_values

}

func part_1(file []byte) int {

	seeds, map_values := process_input(file)
	results := []int{}

	for _, seed := range seeds {
		for i := 1; i <= len(map_values); i++ {
			seed = compute_seed(seed, map_values[i], 1)
		}
		results = append(results, seed)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	return results[0]

}

func part_2(file []byte) int {
	seeds, map_values := process_input(file)

	//fmt.Println(seeds)
	for n := 0; ; n++ {

		location := n

		for j := len(map_values); j > 0; j-- {
			location = compute_seed(location, map_values[j], 2)
		}

		for i := 0; i < len(seeds); i += 2 {
			x := seeds[i]
			y := seeds[i+1]

			if x <= location && location < x+y {
				return n

			}
		}

	}

}

func compute_seed(current_seed int, mapping [][]int, part int) int {
	for _, m := range mapping {
		destination := m[0]
		start := m[1]
		length := m[2]
		if part == 1 {

			if start <= current_seed && current_seed < (start+length) {
				new_seed := destination + (current_seed - start)
				return new_seed
			}
		} else {
			if destination <= current_seed && current_seed < (destination+length) {
				new_seed := start + (current_seed - destination)
				return new_seed
			}

		}
	}
	return current_seed
}

func main() {
	file, err := os.ReadFile("day05/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	part_1 := part_1([]byte(file))
	part_2 := part_2([]byte(file))

	fmt.Println(part_1)
	fmt.Println(part_2)

}
