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
		{expected: 32000000, input: `broadcaster -> a, b, c
		%a -> b
		%b -> c
		%c -> inv
		&inv -> a`},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input)))

}
