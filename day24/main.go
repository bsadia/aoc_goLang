package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, _ := os.ReadFile("day24/input.txt")

	fmt.Println("Part 1: ", part_1(input, 200000000000000.0, 400000000000000.0))

	data := readData(input)

	y := solve(get_col(0, 1, 3, 4, data))
	//fmt.Println(y)

	z := solve(get_col(1, 2, 4, 5, data))
	//fmt.Println(z)

	result := y[0] + y[1] + z[0]
	fmt.Println("Part 2: ", int(math.Round(result)))

}

func part_1(file []byte, Le float64, Me float64) int {

	data := readData(file)

	X := 0.0
	Y := 0.0
	count := 0

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {

			stone1 := data[i]

			stone2 := data[j]

			px, py, vx, vy := float64(stone1[0]), float64(stone1[1]), float64(stone1[3]), float64(stone1[4])
			px2, py2, vx2, vy2 := float64(stone2[0]), float64(stone2[1]), float64(stone2[3]), float64(stone2[4])

			m1 := vy / vx
			b1 := py - (px*vy)/vx

			m2 := vy2 / vx2
			b2 := py2 - (px2*vy2)/vx2
			if m1 != m2 {
				X = (b2 - b1) / (m1 - m2)

				Y = (m1 * X) + b1
			}

			t1 := (X - px) / vx
			t2 := (X - px2) / vx2

			if t1 >= 0 && t2 >= 0 && Le <= X && X <= Me && Le <= Y && Y <= Me {
				count++
			}

		}

	}

	return count

}

func readData(file []byte) [][]int {

	var data [][]int
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(strings.ReplaceAll(line, "@", ","), ",")
		var values []int
		for _, part := range parts {
			val, _ := strconv.Atoi(strings.TrimSpace(part))

			values = append(values, val)
		}

		data = append(data, values)
	}
	return data
}

func get_col(a, b, c, d int, data [][]int) ([][]int, []int) {
	A := make([][]int, len(data))
	B := make([]int, len(data))

	for i, r := range data {
		A[i] = []int{(r[c]), (-(r[d])), (r[a]), (r[b])}
		B[i] = ((r[b])*(r[c]) - (r[a])*(r[d]))
	}

	return A, B
}

func solve(a [][]int, b []int) []float64 {

	m2 := make([][]int, len(a))

	for i := range m2 {
		m2[i] = append((a[i]), b[i])
	}
	mNew := make([][]float64, 4)
	for i := range m2[:4] {
		row := make([]float64, len(m2[i]))
		for j := range m2[i] {
			row[j] = float64(m2[i][j]) - float64(m2[4][j])
		}
		mNew[i] = row
	}
	m := mNew

	for i := range m {
		tem := []float64{}
		for k := range m[i] {
			tt := float64(m[i][k]) / float64(m[i][i])
			if !math.IsNaN(tt) {
				tem = append(tem, (tt))
			} else {
				tem = append(tem, m[i][k])
			}

		}

		m[i] = tem

		for j := i + 1; j < len(m); j++ {
			tem2 := []float64{}
			for k := range m[i] {
				t2 := m[j][k] - m[i][k]*m[j][i]
				tem2 = append(tem2, t2)
			}
			m[j] = tem2
		}
	}
	for i := len(m) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			temp := []float64{}
			for k := range m[i] {
				tt := m[j][k] - m[i][k]*m[j][i]
				temp = append(temp, tt)
			}
			m[j] = temp
		}
	}

	result := make([]float64, len(m))
	for i, r := range m {
		result[i] = (r[len(r)-1])
	}
	//fmt.Println(result)

	return result
}
