package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	input, _ := os.ReadFile("day22/input.txt")
	answer := part_1(input)
	fmt.Println("Part 1: ", answer[0])
	fmt.Println("Part 2: ", answer[1])

}
func part_1(data []byte) []int {

	re := regexp.MustCompile(`\d+`)
	var stack [][]int
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(data)), "\t", ""), "\n")

	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		row := make([]int, len(matches))
		for m := range matches {
			row[m], _ = strconv.Atoi(matches[m])

		}

		stack = append(stack, row)

	}

	sort.Slice(stack, func(i, j int) bool {
		return stack[i][2] < stack[j][2]
	})
	//fmt.Println(stack)

	for i := range stack {
		stack[i][3] += 1
		stack[i][4] += 1
		stack[i][5] += 1
	}
	//fmt.Println(stack)
	peaks := func() [][]int {
		peaks := make([][]int, 12)
		for i := range peaks {
			peaks[i] = make([]int, 12)
		}
		return peaks
	}

	// first run without skipping any bricks
	brick_fall(stack, -1, peaks())

	part1 := 0
	part2 := 0

	for i := range stack {
		stackCopy := make([][]int, len(stack))
		copy(stackCopy, stack)
		ss, dropped := brick_fall(stackCopy, i, peaks())

		if ss {
			part1 += 1
		}
		part2 += dropped

	}

	return []int{part1, part2}

}

func brick_fall(stack [][]int, skip int, peaks [][]int) (bool, int) {

	falls := 0

	for i, arr := range stack {
		if i == skip {
			continue
		}

		x1, y1, z1, x2, y2, z2 := arr[0], arr[1], arr[2], arr[3], arr[4], arr[5]

		// get the peak value
		peak := 0
		for i := x1; i < x2; i++ {
			for j := y1; j < y2; j++ {
				if peaks[i][j] > peak {
					peak = peaks[i][j]
					//peaks[i][j] = peak + z2 - z1
				}
			}
		}

		for i := x1; i < x2; i++ {
			for j := y1; j < y2; j++ {
				peaks[i][j] = peak + z2 - z1
			}
		}
		stack[i] = []int{x1, y1, peak, x2, y2, peak + z2 - z1}

		if peak < z1 {
			falls++
		}
	}
	if falls == 0 {
		return true, falls
	} else {
		return false, falls
	}
}
