package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
)

func part_1(file []byte) int {
	dig_dir := map[string]image.Point{
		"U": {0, -1},
		"D": {0, 1},
		"L": {-1, 0},
		"R": {1, 0},
	}
	re := regexp.MustCompile(`(.) (.*?) \(#(.*?)(.)\)`)
	area := 0
	start := image.Point{0, 0}
	for _, m := range re.FindAllStringSubmatch(string(file), -1) {
		length, _ := strconv.Atoi(m[2])
		next := start.Add(dig_dir[m[1]].Mul(length))
		//fmt.Println(next)
		area += ((start.X*next.Y - start.Y*next.X) + length) // using shoelace formula
		start = next
	}

	return area/2 + 1
}

func part_2(file []byte) int {
	dig_dir := map[string]image.Point{
		"0": {1, 0},  // 0 = right
		"1": {0, 1},  // 1 = down
		"2": {-1, 0}, // 2 = left
		"3": {0, -1}, // 3 = up
	}
	re := regexp.MustCompile(`(.) (.*?) \(#(.*?)(.)\)`)
	area := 0
	start := image.Point{0, 0}
	for _, m := range re.FindAllStringSubmatch(string(file), -1) {
		length, _ := strconv.ParseInt(m[3], 16, strconv.IntSize)

		next := start.Add(dig_dir[m[4]].Mul(int(length)))
		//fmt.Println(next)
		area += ((start.X*next.Y - start.Y*next.X) + int(length)) // using shoelace formula
		start = next
	}

	return area/2 + 1

}

func main() {
	file, _ := os.ReadFile("day18/input.txt")
	fmt.Println("Part 1:", part_1(file))
	fmt.Println("Part 2:", part_2(file))
}
