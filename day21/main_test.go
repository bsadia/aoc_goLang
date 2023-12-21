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
		{expected: 16, input: `...........
		.....###.#.
		.###.##..#.
		..#.#...#..
		....#.#....
		.##..S####.
		.##..#...#.
		.......##..
		.##.#.####.
		.##..##.##.
		...........`},
	}

	assert.Equal(t, tests[0].expected, solution([]byte(tests[0].input), 6)[0])

}
