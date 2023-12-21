package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, _ := os.ReadFile("day20/input.txt")

	fmt.Println("Part 1:", part_1(data))
	fmt.Println("Part 2:", part_2(data))

}

func part_1(data []byte) int {
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(data)), "\t", ""), "\n")
	graph := make(map[string][]string)
	flipflops := make(map[string]bool)
	conjunctions := make(map[string]map[string]bool)

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		destinations := strings.Split(parts[1], ", ")

		if parts[0] == "broadcaster" {
			graph[parts[0]] = destinations
		} else {
			graph[parts[0][1:]] = destinations
		}

		switch parts[0][0] {
		case '%':
			flipflops[parts[0][1:]] = false
		case '&':
			conjunctions[parts[0][1:]] = make(map[string]bool)

		}

	}
	//fmt.Println(graph)

	for src, all_dest := range graph {
		for _, dest := range all_dest {
			if _, ok := conjunctions[dest]; ok {
				conjunctions[dest][src] = false
			}
		}
	}

	//fmt.Println(conjunctions)
	//fmt.Println(flipflops)
	ans := bfs(flipflops, conjunctions, graph)

	//fmt.Println(ans)

	return ans
}

func bfs(flipflop map[string]bool, conjunction map[string]map[string]bool, graph map[string][]string) int {
	low_pulses := 0
	high_pulses := 0

	for i := 1; i <= 1000; i++ {
		q := list.New()
		q.PushBack([]interface{}{"node", "broadcaster", false})

		for q.Len() > 0 {
			element := q.Front()
			sender, node, state := element.Value.([]interface{})[0].(string), element.Value.([]interface{})[1].(string), element.Value.([]interface{})[2].(bool)

			q.Remove(element)

			if state {
				high_pulses += 1
			} else {
				low_pulses += 1
			}

			if _, ok := flipflop[node]; ok {
				if !state {
					flipflop[node] = !flipflop[node]
					new_state := flipflop[node]
					for _, nv := range graph[node] {
						q.PushBack([]interface{}{node, nv, new_state})
					}
				}
			} else if _, ok := conjunction[node]; ok {
				conjunction[node][sender] = state
				//fmt.Println(conjunction[v])
				new_state := !get_all(conjunction[node])
				for _, nv := range graph[node] {
					q.PushBack([]interface{}{node, nv, new_state})
				}
			} else if node == "broadcaster" {
				for _, nv := range graph[node] {
					q.PushBack([]interface{}{node, nv, state})
				}
			}
		}
	}

	return low_pulses * high_pulses

}

func get_all(m map[string]bool) bool {
	for _, v := range m {
		if !v {
			return false
		}
	}
	return true
}

func part_2(data []byte) int {

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	graph := make(map[string][]string)
	var res []int
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		//fmt.Println(parts)
		graph[parts[0]] = strings.Split(parts[1], ", ")
	}
	//fmt.Println(graph)

	for _, m := range graph["broadcaster"] {

		bin := ""
		for {

			g := graph["%"+m]
			if len(g) == 2 {
				bin = "1" + bin
			} else {
				if _, ok := graph["%"+g[0]]; !ok {
					bin = "1" + bin
				} else {
					bin = "0" + bin
				}
			}
			next_v := make([]string, 0)
			for _, n := range graph["%"+m] {
				if _, ok := graph["%"+n]; ok {
					next_v = append(next_v, n)
				}
			}
			if len(next_v) == 0 {
				break
			}
			m = next_v[0]

		}

		val, _ := strconv.ParseInt(bin, 2, 64)

		res = append(res, int(val))
	}

	lcmValue := res[0]
	for _, v := range res[1:] {
		lcmValue = lcm(lcmValue, v)
	}

	return lcmValue
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
