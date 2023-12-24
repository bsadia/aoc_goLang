package main

import (
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

var (
	start         = image.Point{}
	end           = image.Point{}
	width, height int
	intersections = []image.Point{}
	neighbors     = make(map[image.Point][]image.Point)
	graph         = make(map[image.Point][][2]interface{})
	directions    = []image.Point{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
)

func main() {

	file, _ := os.ReadFile("day23/input.txt")
	fmt.Println("Part 1: ", Part1(file), "with total time: ", time.Since(time.Now()))
	start_time := time.Now()
	fmt.Println("Part 2: ", Part2(file), "with  total time: ", time.Since(start_time))

}
func get_dist(cur image.Point, dist int, seen map[image.Point]bool) (image.Point, int) {
	for _, p := range intersections {
		if p == cur {
			return cur, dist
		}
	}

	for _, nb := range neighbors[cur] {
		if !seen[nb] {
			seen[cur] = true
			return get_dist(nb, dist+1, seen)
		}
	}

	return image.Point{}, 0
}

func bfs(start, end image.Point, score int, seen map[image.Point]bool) []int {
	if start == end {
		return []int{score}
	}

	var scores []int

	for _, pair := range graph[start] {

		current := pair[0].(image.Point)
		dist := pair[1].(int)
		if !seen[current] {
			seen[current] = true
			scores = append(scores, bfs(current, end, score+dist, seen)...)
			delete(seen, current)
		}
	}

	return scores
}

func Part1(file []byte) int {

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(file)), "\t", ""), "\n")
	width = len(lines[0])
	height = len(lines)

	start = image.Point{0, 1}
	end = image.Point{height - 1, width - 2}

	intersections, neighbors = get_neighbors_intersections(lines, true)
	graph = get_graph(intersections, neighbors)

	return get_max(bfs(start, end, 0, map[image.Point]bool{start: true}))

}

func Part2(file []byte) int {

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(file)), "\t", ""), "\n")

	width = len(lines[0])
	height = len(lines)

	start = image.Point{0, 1}
	end = image.Point{height - 1, width - 2}

	intersections, neighbors = get_neighbors_intersections(lines, false)

	graph = get_graph(intersections, neighbors)

	return get_max(bfs(start, end, 0, map[image.Point]bool{start: true}))
}

func get_graph(intersections []image.Point, neighbors map[image.Point][]image.Point) map[image.Point][][2]interface{} {
	graph = make(map[image.Point][][2]interface{})

	for _, i := range intersections {
		for _, n := range neighbors[i] {
			t, d := get_dist(n, 1, map[image.Point]bool{i: true})
			graph[i] = append(graph[i], [2]interface{}{t, d})
		}
	}
	return graph
}

func get_neighbors_intersections(lines []string, part_1 bool) ([]image.Point, map[image.Point][]image.Point) {

	intersections = []image.Point{start, end}
	neighbors = make(map[image.Point][]image.Point)

	for i, line := range lines {
		for j, val := range line {
			if strings.ContainsAny(string(val), ".v^><") {

				count := 0

				var cord image.Point
				if part_1 {
					switch string(val) {

					case "v":
						mr, nc := i+directions[2].X, j+directions[2].Y
						cord = image.Point{mr, nc}

					case "^":

						mr, nc := i+directions[3].X, j+directions[3].Y
						cord = image.Point{mr, nc}
					case ">":
						mr, nc := i+directions[1].X, j+directions[1].Y
						cord = image.Point{mr, nc}

					case "<":

						mr, nc := i+directions[0].X, j+directions[0].Y
						cord = image.Point{mr, nc}

					default:

						for _, dir := range directions {
							m, n := i+dir.X, j+dir.Y
							if (m >= 0 && m < height && n >= 0 && n < width) && (string(lines[m][n]) != "#") {

								count++
								neighbors[image.Point{i, j}] = append(neighbors[image.Point{i, j}], image.Point{m, n})
							}
						}
					}
					if cord.X >= 0 && cord.X < height && cord.Y >= 0 && cord.Y < width {

						neighbors[image.Point{i, j}] = append(neighbors[image.Point{i, j}], cord)
					}
					if count > 2 {
						intersections = append(intersections, image.Point{i, j})
					}
				} else {
					count := 0
					for _, dir := range directions {
						m, n := i+dir.X, j+dir.Y
						if (m >= 0 && m < height && n >= 0 && n < width) && (string(lines[m][n]) != "#") {

							count++
							neighbors[image.Point{i, j}] = append(neighbors[image.Point{i, j}], image.Point{m, n})
						}
					}
					if count > 2 {
						intersections = append(intersections, image.Point{i, j})
					}

				}

			}
		}
	}
	return intersections, neighbors
}
func get_max(result []int) int {
	max_score := result[0]
	for _, score := range result {
		if score > max_score {
			max_score = score
		}
	}
	return max_score
}
