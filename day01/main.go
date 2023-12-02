package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var numberMap = map[string]string{
	"zero":  "zero0zero",
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func part_1(line string) (combinedDigits int) {
	foundFirst := false
	var first, last int

	for _, char := range line {
		if unicode.IsDigit(char) {
			if !foundFirst {
				first = int(char - '0')
				foundFirst = true
			}
			last = int(char - '0')
		}
	}

	return first*10 + last
}

func part_2(line string) (combinedDigits int) {

	for word, digit := range numberMap {

		if strings.Contains(line, word) {
			line = strings.ReplaceAll(line, word, digit)
		}
	}
	combinedDigits = part_1(line)

	return combinedDigits
}

func main() {

	// For this problem, we will read one line at a time from the input file and process it
	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	sum_part1 := 0
	sum_part2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		combinedDigits := part_1(line)
		sum_part1 += combinedDigits

		combinedDigits = part_2(line)
		sum_part2 += combinedDigits

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Sum of Part_1: %d\n", sum_part1)
	fmt.Printf("Sum of part_2: %d\n", sum_part2)
}
