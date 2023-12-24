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
		{expected: 94, input: `#.#####################
		#.......#########...###
		#######.#########.#.###
		###.....#.>.>.###.#.###
		###v#####.#v#.###.#.###
		###.>...#.#.#.....#...#
		###v###.#.#.#########.#
		###...#.#.#.......#...#
		#####.#.#.#######.#.###
		#.....#.#.#.......#...#
		#.#####.#.#.#########v#
		#.#...#...#...###...>.#
		#.#.#v#######v###.###v#
		#...#.>.#...>.>.#.###.#
		#####v#.#.###v#.#.###.#
		#.....#...#...#.#.#...#
		#.#########.###.#.#.###
		#...###...#...#...#.###
		###.###.#.###v#####v###
		#...#...#.#.>.>.#.>.###
		#.###.###.#.###.#.#v###
		#.....###...###...#...#
		#####################.#`},
		{expected: 154},
	}

	assert.Equal(t, tests[0].expected, Part1(([]byte(tests[0].input))))
	assert.Equal(t, tests[1].expected, Part2(([]byte(tests[0].input))))

}