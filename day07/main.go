package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
	hType int
}

var hand_types = map[string]int{
	"FIVE_OF_A_KIND":  7,
	"FOUR_OF_A_KIND":  6,
	"FULL_HOUSE":      5,
	"THREE_OF_A_KIND": 4,
	"TWO_PAIR":        3,
	"ONE_PAIR":        2,
	"HIGH_CARD":       1,
}

func main() {
	input, _ := os.ReadFile("day07/input.txt")

	fmt.Println("Part 1:", part_1(input, false))
	fmt.Println("Part 2:", part_1(input, true))
}
func part_1(input []byte, joker bool) int {

	hands := []Hand{}

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		fields := strings.Split(s, " ")
		cards := fields[0]
		bid, _ := strconv.Atoi(fields[1])
		hand_type := get_hand_type(cards, joker)
		hands = append(hands, Hand{cards, bid, hand_type})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compare_hands(hands[i], hands[j], joker)
	})
	//fmt.Println(hands)
	total := 0
	for index, hand := range hands {
		total += (index + 1) * hand.Bid
	}
	//fmt.Println(total)
	return total

}

func get_hand_type(card string, joker bool) int {

	list_1 := map[rune]int{}
	list_2 := []int{}
	jCounter := 0

	for _, s := range card {
		if joker {
			if s == 'J' {
				jCounter++
				continue
			}
		}

		list_1[s] += 1

	}

	for _, value := range list_1 {
		list_2 = append(list_2, value)
	}

	sort.Ints(list_2)
	highest := 0 + jCounter
	if len(list_2) > 0 {
		highest = (list_2[len(list_2)-1]) + jCounter
	}
	secondHighest := 0
	if len(list_2) > 1 {
		secondHighest = (list_2[len(list_2)-2])

	}

	htype := 0

	switch highest {
	case 5:
		htype = hand_types["FIVE_OF_A_KIND"]
	case 4:
		htype = hand_types["FOUR_OF_A_KIND"]
	case 3:
		htype = hand_types["THREE_OF_A_KIND"]
		if secondHighest == 2 {
			htype = hand_types["FULL_HOUSE"]
		}
	case 2:
		htype = hand_types["ONE_PAIR"]
		if secondHighest == 2 {
			htype = hand_types["TWO_PAIR"]
		}
	default:
		htype = hand_types["HIGH_CARD"]
	}
	return htype

}

func compare_hands(h1 Hand, h2 Hand, joker bool) bool {

	if h1.hType == h2.hType {
		for i, v1 := range h1.Cards {
			v2 := rune(h2.Cards[i])
			if v1 == v2 {
				continue
			}

			return findPos(v1, joker) < findPos(v2, joker)
		}

		return true
	} else {
		return h1.hType < h2.hType
	}
}
func findPos(c rune, joker bool) int {
	cardPositions := make(map[rune]int)
	order := "23456789TJQKA"
	if joker {
		order = "J23456789TQKA"
	}
	for i, v := range order {
		cardPositions[v] = i
	}
	return cardPositions[c]
}
