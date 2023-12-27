package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day15/input.txt")

	fmt.Println(part_1(input))
	fmt.Println(part_2(input))
}

func part_1(file []byte) int {

	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")
	sequence := strings.Split(lines[0], ",")

	hashSum := 0
	for _, step := range sequence {
		result := 0
		for _, char := range step {
			result = (result + int(char)) * 17 % 256
		}
		hashSum += result

	}

	return hashSum

}
func part_2(file []byte) int {
	re := regexp.MustCompile(`(\w+)([-=])(\d*)`)

	part2 := 0
	boxes, focal := [256][]string{}, map[string]int{}
	for _, m := range re.FindAllStringSubmatch(string(file), -1) {
		h := 0
		for _, r := range m[1] {
			h = (h + int(r)) * 17 % 256
		}
		i := slices.Index(boxes[h], m[1])

		if m[2] == "-" && i != -1 {
			boxes[h] = slices.Delete(boxes[h], i, i+1)
		} else if m[2] == "=" {
			focal[m[1]] = int(m[3][0] - '0')
			if i == -1 {
				boxes[h] = append(boxes[h], m[1])
			}
		}
	}

	for i, b := range boxes {
		for j, l := range b {
			part2 += (i + 1) * (j + 1) * focal[l]
		}
	}
	return part2
}
