package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day12/input.txt")
	fmt.Println(part_1(input))
	fmt.Println(part_2(input))
}
func part_2(file []byte) int {
	var all_keys = map[string]int{}
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(file)), "\t", ""), "\n")

	arr := 0
	for _, line := range lines {
		spring := strings.Fields(line)
		list := []int{}
		t := strings.Repeat(spring[1]+",", 5)
		//fmt.Println(t)
		for _, c := range strings.Split(t[:len(t)-1], ",") {
			num, _ := strconv.Atoi(c)
			list = append(list, num)
		}
		s := strings.Repeat(spring[0]+"?", 5)
		arr += get_groups(s, list, all_keys)
	}
	fmt.Println(arr)
	return arr

}

func part_1(file []byte) int {
	var all_keys = map[string]int{}
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(file)), "\t", ""), "\n")
	arr := 0
	for _, line := range lines {
		spring := strings.Fields(line)
		list := []int{}

		for _, c := range strings.Split(spring[1], ",") {
			num, _ := strconv.Atoi(c)
			list = append(list, num)
		}

		arr += get_groups(spring[0]+"?", list, all_keys)
	}
	fmt.Println(arr)
	return arr

}

func get_groups(spring string, lis []int, all_keys map[string]int) (r int) {

	var buffer bytes.Buffer
	buffer.WriteString(spring)
	buffer.WriteString(fmt.Sprint(lis))
	key := buffer.String()

	if r, ok := all_keys[key]; ok {
		return r
	}

	if spring == "" {
		if len(lis) == 0 {
			r = 1
		} else {
			r = 0
		}
		all_keys[key] = r
		return r
	}

	if spring[0] == '.' || spring[0] == '?' {
		r += get_groups(spring[1:], lis, all_keys)
	}
	if spring[0] == '#' || spring[0] == '?' {
		if len(lis) > 0 && lis[0] >= 0 && lis[0] < len(spring) && len(spring) > 0 {
			if !strings.Contains(spring[:lis[0]], ".") &&
				spring[lis[0]] != '#' {
				r += get_groups(spring[lis[0]+1:], lis[1:], all_keys)
			}
		}
	}
	all_keys[key] = r
	return r

}
