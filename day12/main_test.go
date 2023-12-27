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
		{expected: 21, input: `???.### 1,1,3
		.??..??...?##. 1,1,3
		?#?#?#?#?#?#?#? 1,3,1,6
		????.#...#... 4,1,1
		????.######..#####. 1,6,5
		?###???????? 3,2,1`},
		{expected: 525152},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input)))
	assert.Equal(t, tests[1].expected, part_2([]byte(tests[0].input)))

}
