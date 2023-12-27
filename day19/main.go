package main

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Cat  string
	Op   string
	Val  int
	Next string
}

func main() {

	input, _ := os.ReadFile("day19/input.txt")

	fmt.Println("Part 1: ", part_1(input))
	fmt.Println("Part 2: ", part_2(input))

}
func parse_input(file []byte) (map[string][]Rule, []map[string]int) {
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(file)), "\t", ""), "\n\n")

	wf := map[string][]Rule{}
	for _, m := range regexp.MustCompile(`(\w+){(.*),(\w+)}`).FindAllStringSubmatch(lines[0], -1) {
		w := m[1]

		for _, m := range regexp.MustCompile(`(\w+)(<|>)(\d+):(\w+)`).FindAllStringSubmatch(m[2], -1) {
			v, _ := strconv.Atoi(m[3])
			wf[w] = append(wf[w], Rule{m[1], m[2], v, m[4]})
		}
		wf[w] = append(wf[w], Rule{Next: m[3]})
	}

	all_ratings := []map[string]int{}

	for _, s := range strings.Fields(lines[1]) {
		re := regexp.MustCompile(`\{(.*)\}`)
		matches := re.FindStringSubmatch(s)

		if len(matches) > 1 {
			pairs := strings.Split(matches[1], ",")
			rating := make(map[string]int)

			for _, pair := range pairs {
				kv := strings.Split(pair, "=")
				if len(kv) == 2 {
					key := kv[0]
					value, _ := strconv.Atoi(kv[1])
					rating[key] = value
				}
			}
			all_ratings = append(all_ratings, rating)

		}
	}
	return wf, all_ratings
}

func part_1(file []byte) int {
	wf, all_ratings := parse_input(file)
	answer := 0

	for _, part := range all_ratings {

		current := "in"

		for current != "A" && current != "R" {
			currentWorkflow := wf[current]
			for _, rule := range currentWorkflow {
				if rule.Cat == "" || (rule.Op == "<" && part[rule.Cat] < rule.Val) ||
					(rule.Op == ">" && part[rule.Cat] > rule.Val) {
					current = rule.Next
					break
				}
			}
		}
		if current == "A" {
			result := 0
			for _, val := range part {
				result += val
			}
			answer += result
		}
	}
	fmt.Println(answer)
	return answer
}

func part_2(file []byte) int {
	wf, _ := parse_input(file)

	ranges := map[string][2]int{
		"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000},
	}

	return get_combinations(wf, "in", ranges)
}

func get_combinations(wf map[string][]Rule, current string, ranges map[string][2]int) (result int) {

	if current == "R" {
		return 0
	}
	if current == "A" {

		result := 1
		for _, r := range ranges {
			result *= r[1] - r[0] + 1
		}
		return result
	}
	currentWorkflow := wf[current]
	for _, rule := range currentWorkflow {

		next := maps.Clone(ranges)

		if rule.Op == "<" {
			next[rule.Cat] = [2]int{next[rule.Cat][0], min(next[rule.Cat][1], rule.Val-1)}
			ranges[rule.Cat] = [2]int{max(next[rule.Cat][0], rule.Val), ranges[rule.Cat][1]}
		} else if rule.Op == ">" {
			next[rule.Cat] = [2]int{max(next[rule.Cat][0], rule.Val+1), next[rule.Cat][1]}
			ranges[rule.Cat] = [2]int{ranges[rule.Cat][0], min(next[rule.Cat][1], rule.Val)}

		}
		result += get_combinations(wf, rule.Next, next)

	}
	return result
}
