package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day08/input.txt")

	fmt.Println(part_1(input))
	fmt.Println(part_2(input))
}
func part_1(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	network, direction := parse(lines)

	start := "AAA"
	end := "ZZZ"
	steps := 0
	for !(start == end) {
		for _, dir := range direction {
			if dir == 'L' {
				start = network[start][0]
			} else {
				start = network[start][1]
			}
			steps++
			if start == end {
				break
			}
		}
	}
	return steps

}
func parse(lines []string) (map[string][2]string, string) {

	direction := lines[0]
	path := lines[1]
	re := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

	network := make(map[string][2]string)
	for _, m := range re.FindAllStringSubmatch(path, -1) {
		network[strings.TrimSpace(m[1])] = [2]string{strings.TrimSpace(m[2]), strings.TrimSpace(m[3])}
	}
	return network, direction
}

func part_2(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	network, direction := parse(lines)
	start := "A"
	end := "Z"
	results := []int{}

	for node := range network {
		if !strings.HasSuffix(node, start) {
			continue
		}

		steps := 0
		for !strings.HasSuffix(node, end) {
			next_dir := direction[steps%len(direction)]
			if next_dir == 'L' {
				node = network[node][0]
			} else {
				node = network[node][1]
			}
			steps++

		}
		results = append(results, steps)

	}
	val := results[0]
	for i := 1; i < len(results); i++ {
		val = lcm(val, results[i])
	}

	return val
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
