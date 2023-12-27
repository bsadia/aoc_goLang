package main

import (
	"container/heap"
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
)

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }

type State struct {
	Pos image.Point
	Dir image.Point
}

func main() {
	input, _ := os.ReadFile("day17/input.txt")
	answer := part_1(input)
	fmt.Println("Part 1: ", answer[0])
	fmt.Println("Part 2: ", answer[1])

}
func part_1(file []byte) []int {
	lines := strings.Fields(string(file))

	heatMap := map[image.Point]int{}
	for y, s := range lines {
		for x, r := range s {
			num, _ := strconv.Atoi(string(r))
			heatMap[image.Point{x, y}] = num
		}
	}

	goal := image.Point{len(lines[0]) - 1, len(lines) - 1}

	part_1 := find_shortest_path(1, 3, heatMap, goal)
	part_2 := find_shortest_path(4, 10, heatMap, goal)

	return []int{part_1, part_2}
}

func find_shortest_path(min, max int, heatMap map[image.Point]int, end image.Point) int {
	queue := make(PQ[State], 0)
	visited := make(map[State]bool)

	enqueue(&queue, State{image.Point{0, 0}, image.Point{1, 0}}, 0)
	enqueue(&queue, State{image.Point{0, 0}, image.Point{0, 1}}, 0)

	for len(queue) > 0 {
		state, heat := dequeue(&queue)

		if state.Pos == end {
			return heat
		}

		if _, ok := visited[state]; ok {
			continue
		}

		visited[state] = true

		for i := -max; i <= max; i++ {
			move(i, min, state, heatMap, &queue, heat)
		}
	}

	return -1
}

func enqueue(queue *PQ[State], state State, priority int) {
	heap.Push(queue, pqi[State]{state, priority})
}

func dequeue(queue *PQ[State]) (State, int) {
	item := heap.Pop(queue).(pqi[State])
	return item.v, item.p
}

func move(i, min int, state State, heatMap map[image.Point]int, queue *PQ[State], heat int) {
	newPos := state.Pos.Add(state.Dir.Mul(i))

	if _, ok := heatMap[newPos]; !ok || i > -min && i < min {
		return
	}

	step := 1
	if math.Signbit(float64(i)) {
		step = -1
	}
	newHeat := 0

	for j := step; j != i+step; j += step {
		newHeat += heatMap[state.Pos.Add(state.Dir.Mul(j))]
	}
	newHeat = heat + newHeat
	enqueue(queue, State{newPos, image.Point{state.Dir.Y, state.Dir.X}}, newHeat)
}
