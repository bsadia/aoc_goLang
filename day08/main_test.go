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
		{expected: 2, input: `RL

		AAA = (BBB, CCC)
		BBB = (DDD, EEE)
		CCC = (ZZZ, GGG)
		DDD = (DDD, DDD)
		EEE = (EEE, EEE)
		GGG = (GGG, GGG)
		ZZZ = (ZZZ, ZZZ)`},
		{expected: 6, input: `LLR

		AAA = (BBB, BBB)
		BBB = (AAA, ZZZ)
		ZZZ = (ZZZ, ZZZ)`},
		{expected: 6, input: `LR

		11A = (11B, XXX)
		11B = (XXX, 11Z)
		11Z = (11B, XXX)
		22A = (22B, XXX)
		22B = (22C, 22C)
		22C = (22Z, 22Z)
		22Z = (22B, 22B)
		XXX = (XXX, XXX)`},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input)))
	assert.Equal(t, tests[1].expected, part_1([]byte(tests[1].input)))
	assert.Equal(t, tests[2].expected, part_2([]byte(tests[2].input)))

}
