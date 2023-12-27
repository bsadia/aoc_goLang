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
		{expected: 136, input: `O....#....
		O.OO#....#
		.....##...
		OO.#O....O
		.O.....O#.
		O.#..O.#.#
		..O..#O..O
		.......O..
		#....###..
		#OO..#....`},
		{expected: 64},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input)))
	assert.Equal(t, tests[1].expected, part_2([]byte(tests[0].input)))

}
