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
		{expected: 405, input: `#.##..##.
		..#.##.#.
		##......#
		##......#
		..#.##.#.
		..##..##.
		#.#.##.#.
		
		#...##..#
		#....#..#
		..##..###
		#####.##.
		#####.##.
		..##..###
		#....#..#`},
		{expected: 400},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input))[0])
	assert.Equal(t, tests[1].expected, part_1([]byte(tests[0].input))[1])

}
