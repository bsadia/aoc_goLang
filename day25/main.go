package main

import (
	"container/list"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Link struct {
	from, to string
}
type Graph map[string][]string

func main() {
	data, _ := os.ReadFile("day25/input.txt")
	fmt.Println("Part 1: ", part_1(data))
}
func part_1(data []byte) int {

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(data)), "\t", ""), "\n")

	graph := Graph{}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		node := parts[0]
		ends := strings.Split(parts[1], " ")
		for _, end := range ends {
			graph[node] = append(graph[node], end)
			graph[end] = append(graph[end], node)
		}
	}

	edges_to_remove := []Link{}
	for len(edges_to_remove) < 3 {
		edge := remove_max_link(count_edges(graph))
		edges_to_remove = append(edges_to_remove, edge)

		graph[edge.from] = remove(graph[edge.from], edge.to)
		graph[edge.to] = remove(graph[edge.to], edge.from)
	}
	//fmt.Println("edges_to_remove", edges_to_remove)

	visited := map[string]bool{}
	queue := []string{edges_to_remove[0].from}

	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]

		for _, to := range graph[from] {
			if _, ok := visited[to]; ok {
				continue
			}
			queue = append(queue, to)
			visited[to] = true
		}
	}
	count := len(visited)
	return count * (len(graph) - count)
}

func remove(slice []string, s string) []string {
	for i, val := range slice {
		if val == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

type kv struct {
	Key   Link
	Value int
}

func count_edges(nodes Graph) map[Link]int {
	found := map[Link]int{}

	for start := range nodes {
		visited := map[string]bool{}
		queue := list.New()
		queue.PushBack(start)

		for queue.Len() > 0 {
			element := queue.Front()
			from := element.Value.(string)
			queue.Remove(element)

			for _, to := range nodes[from] {
				if _, ok := visited[to]; ok {
					continue
				}
				queue.PushBack(to)
				visited[to] = true
				var edge Link
				if from < to {
					edge = Link{from, to}
				} else {
					edge = Link{to, from}
				}
				found[edge]++
			}
		}
	}
	return found
}
func remove_max_link(edges map[Link]int) Link {

	var ss []kv
	for k, v := range edges {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	max_edge := ss[0].Key
	delete(edges, max_edge)
	return max_edge
}
