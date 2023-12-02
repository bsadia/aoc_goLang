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
		{expected: 142, input: `1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`},
		{expected: 281,
			input: `two1nine
		eightwothree
		abcone2threexyz
		xtwone3four
		4nineeightseven2
		zoneight234
		7pqrstsixteen`},
	}

	test_part1 := 0
	test_part2 := 0
	for _, input := range strings.Split(tests[0].input, "\n") {
		test_part1 += part_1(input)

	}
	for _, input := range strings.Split(tests[1].input, "\n") {
		test_part2 += part_2(input)

	}
	assert.Equal(t, tests[0].expected, test_part1)
	assert.Equal(t, tests[1].expected, test_part2)
}
