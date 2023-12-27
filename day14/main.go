package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day14/input.txt")
	fmt.Println("Part 1: ", part_1(input))
	fmt.Println("Part 2: ", part_2(input))
}

func part_1(file []byte) int {
	dish := strings.Fields(string(file))

	grid := make([][]byte, 0, len(dish))
	for _, line := range dish {
		grid = append(grid, []byte(line))
	}

	for row := 0; row < len(grid); row++ {
		for col, val := range grid[row] {
			if val == 'O' {
				for i := row - 1; i >= 0 && grid[i][col] == '.'; i-- {
					grid[i+1][col], grid[i][col] = grid[i][col], grid[i+1][col]
				}
			}
		}
	}

	result := 0
	for row := len(grid) - 1; row >= 0; row-- {
		for _, val := range grid[row] {
			if val == 'O' {
				result += len(grid) - row
			}
		}
	}
	return result

}

func part_2(file []byte) int {

	dish := strings.Fields(string(file))

	rot := make([]string, len(dish[0]))
	for row := range dish {
		for col := range dish[row] {
			rot[col] = fmt.Sprintf("%s%c", rot[col], dish[len(dish)-row-1][col])

		}
	}

	cycles := 1000000000
	seen := map[string]int{}
	for i := 0; i < cycles; i++ {
		if s, ok := seen[fmt.Sprint(rot)]; ok {
			i = cycles - (cycles-i)%(i-s)
		}
		seen[fmt.Sprint(rot)] = i
		for i := 0; i < 4; i++ {

			dish_ := slices.Clone(rot)
			for i := range dish_ {
				for strings.Contains(dish_[i], "O.") {
					dish_[i] = strings.ReplaceAll(dish_[i], "O.", ".O")
				}
			}

			temp2 := make([]string, len(dish_[0]))
			for row := range dish_ {
				for col := range dish_[row] {
					temp2[col] = fmt.Sprintf("%s%c", temp2[col], dish_[len(dish_)-row-1][col])
				}
			}
			rot = temp2

		}
	}

	temp := make([]string, len(rot[0]))
	for row := range rot {
		for col := range rot[row] {
			temp[col] = fmt.Sprintf("%s%c", temp[col], rot[len(rot)-row-1][col])
		}
	}
	l := 0
	for i, s := range temp {
		l += strings.Count(s, "O") * (i + 1)
	}

	return l
}
