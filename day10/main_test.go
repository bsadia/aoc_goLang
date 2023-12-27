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
		{expected: 8, input: `..F7.
		.FJ|.
		SJ.L7
		|F--J
		LJ...`},
		{expected: 4, input: `...........
		.S-------7.
		.|F-----7|.
		.||.....||.
		.||.....||.
		.|L-7.F-J|.
		.|..|.|..|.
		.L--J.L--J.
		...........`},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input))[0])
	assert.Equal(t, tests[1].expected, part_1([]byte(tests[1].input))[1])

}
