package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day13/input.txt")

	answer := part_1(input)

	fmt.Println("Part 1: ", answer[0])
	fmt.Println("Part 2: ", answer[1])
}

func part_1(file []byte) []int {
	sum_p1 := 0
	sum_p2 := 0
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(file)), "\t", ""), "\n\n")
	for _, line := range lines {
		l := strings.Fields(line)
		row := []string{}
		col := make([]string, len(l[0]))
		for _, s := range l {
			row = append(row, s)
			for i, r := range s {
				col[i] += string(r)
			}
		}

		sum_p1 += compare(row, false)*100 + compare(col, false)
		sum_p2 += compare(row, true)*100 + compare(col, true)

	}
	return []int{sum_p1, sum_p2}
}

func compare(cols []string, smudge bool) int {
	for i := 1; i < len(cols); i++ {
		l := min(i, len(cols)-i)
		a, b := slices.Clone(cols[i-l:i]), cols[i:i+l]
		slices.Reverse(a)
		if smudge {
			diffs := 0
			for i := range a {
				for j := range a[i] {
					if a[i][j] != b[i][j] {
						diffs++
					}
				}
			}
			if diffs == 1 {
				return i
			}
		} else if slices.Equal(a, b) {
			return i
		}
	}
	return 0

}
