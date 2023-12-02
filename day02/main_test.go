package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	tests := []struct {
		expected int
		input    string
	}{
		{expected: 8, input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`},
		{expected: 2286},
	}

	test_part1 := 0
	test_part2 := 0
	for index, input := range strings.Split(tests[0].input, "\n") {
		if part_1(input) {
			test_part1 += index + 1
		}

		temp := part_2(input)
		test_part2 += temp

	}
	assert.Equal(t, tests[0].expected, test_part1)
	assert.Equal(t, tests[1].expected, test_part2)
}
