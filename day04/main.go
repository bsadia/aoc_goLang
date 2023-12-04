package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func part_1(file []byte) (int, map[int]int) {

	cards_copies := map[int]int{}

	lines := strings.Split(string(file), "\n")
	re := regexp.MustCompile(`\d+`)

	part1 := 0

	for i, s := range lines {

		game := strings.Split(s, ":")

		your_number, winning_number := re.FindAllString((strings.Split(game[1], "|")[0]), -1), re.FindAllString(strings.Split(game[1], "|")[1], -1)

		winners := 0
		for _, value := range your_number {
			for _, winning_value := range winning_number {
				if value == winning_value {
					winners++
				}
			}
		}

		cards_copies[i]++
		for j := 1; j <= winners; j++ {
			cards_copies[i+j] += cards_copies[i]
		}

		if winners > 0 {
			part1 += 1 << (winners - 1)
		}

	}

	return part1, cards_copies

}

func part_2(card_copies map[int]int) int {

	part2 := 0

	for _, value := range card_copies {
		part2 += value
	}

	return part2
}

func main() {
	file, err := os.ReadFile("day04/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	part_1, card_copies := part_1(file)
	part_2 := part_2(card_copies)

	fmt.Println(part_1)
	fmt.Println(part_2)
}
