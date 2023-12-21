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
		{expected: 114, input: `0 3 6 9 12 15
		1 3 6 10 15 21
		10 13 16 21 30 45`},
		{expected: 2},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input))[0])
	assert.Equal(t, tests[1].expected, part_1([]byte(tests[0].input))[1])

}
