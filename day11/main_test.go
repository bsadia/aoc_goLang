package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	tests := []struct {
		expected int
		input    string
	}{
		{expected: 374, input: `...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....`},
		{expected: 1030},
		{expected: 8410},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input), 2))
	assert.Equal(t, tests[1].expected, part_1([]byte(tests[0].input), 10))
	assert.Equal(t, tests[2].expected, part_1([]byte(tests[0].input), 100))

}
