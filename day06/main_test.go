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
		{expected: 288, input: `Time:      7  15   30
		Distance:  9  40  200`},
		{expected: 71503},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input)))
	assert.Equal(t, tests[1].expected, part_2([]byte(tests[0].input)))
}
