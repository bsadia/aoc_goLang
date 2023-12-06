package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func process_input(file string) []int {

	line := strings.TrimSpace((strings.Split(file, ":"))[1])

	output := []int{}
	matches := regexp.MustCompile(`-?\d+`).FindAllString(line, -1)
	for _, match := range matches {
		t, _ := strconv.Atoi(string(match))
		output = append(output, t)
	}

	return output
}

func part_1(file []byte) int {

	lines := strings.Split(string(file), "\n")

	time := process_input(lines[0])
	distance := process_input(lines[1])

	winning_arr := make([]int, len(time))
	prod := 1
	for ind, t := range time {
		for i := 1; i <= t; i++ {
			time_remain := t - i
			speed := i
			disp := speed * time_remain
			if disp > distance[ind] {
				winning_arr[ind]++
			}
		}
		prod *= winning_arr[ind]

	}

	return prod

}

func part_2(file []byte) int {

	lines := strings.Split(string(file), "\n")
	time, _ := strconv.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))

	winning_counter := 0
	for j := 1; j <= time; j++ {
		time_remain := time - j
		speed := j
		if (speed * time_remain) > distance {
			winning_counter++
		}
	}

	return winning_counter

}

func main() {
	file, err := os.ReadFile("day06/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	part_1 := part_1([]byte(file))
	part_2 := part_2([]byte(file))

	fmt.Println(part_1)
	fmt.Println(part_2)

}
