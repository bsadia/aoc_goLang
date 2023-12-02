package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Condition: only 12 red cubes, 13 green cubes, and 14 blue cubes?
func part_1(input string) bool {
	game_subset := strings.Split((strings.Split(input, ":")[1]), ";")

	for _, val := range game_subset {

		pattern := regexp.MustCompile(`(\d+)\s+([a-zA-Z]+)`)
		matches := pattern.FindAllStringSubmatch(val, -1)

		// Iterate over the matches and check conditions
		for _, match := range matches {
			number, _ := strconv.Atoi(match[1])
			color := match[2]

			// Check conditions for each color-value pair
			if (color == "red" && number > 12) ||
				(color == "green" && number > 13) ||
				(color == "blue" && number > 14) {
				return false
			}

		}
	}

	return true
}

func part_2(input string) int {

	game_subset := strings.Split((strings.Split(input, ":")[1]), ";")
	red := 0
	blue := 0
	green := 0
	for _, val := range game_subset {

		pattern := regexp.MustCompile(`(\d+)\s+([a-zA-Z]+)`)
		matches := pattern.FindAllStringSubmatch(val, -1)

		// Iterate over the matches and check conditions

		for _, match := range matches {
			number, _ := strconv.Atoi(match[1])
			color := match[2]
			if color == "red" && red < number {
				red = number
			}
			if color == "blue" && blue < number {
				blue = number
			}
			if color == "green" && green < number {
				green = number
			}

		}
	}

	return red * blue * green
}

func main() {

	// For this problem, we will read one line at a time from the input file and process it
	file, err := os.Open("day02/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum_part1 := 0
	power_part2 := 0

	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()

		if part_1(line) {
			sum_part1 += lineNumber
		}
		powerNumber := part_2(line)
		power_part2 += powerNumber
	}
	fmt.Printf("Part_1: %d\n", sum_part1)
	fmt.Printf("Part_2: %d\n", power_part2)

}
