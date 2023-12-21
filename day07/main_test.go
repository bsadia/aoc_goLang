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
		{expected: 6440, input: `32T3K 765
		T55J5 684
		KK677 28
		KTJJT 220
		QQQJA 483`},
		{expected: 5905},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input), false))
	assert.Equal(t, tests[1].expected, part_1([]byte(tests[0].input), true))

}
