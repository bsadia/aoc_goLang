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
		{expected: 102, input: `2413432311323
		3215453535623
		3255245654254
		3446585845452
		4546657867536
		1438598798454
		4457876987766
		3637877979653
		4654967986887
		4564679986453
		1224686865563
		2546548887735
		4322674655533`},
		{expected: 94},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input))[0])
	assert.Equal(t, tests[1].expected, part_1([]byte(tests[0].input))[1])

}
