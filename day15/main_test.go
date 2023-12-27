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
		{expected: 1320, input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`},
		{expected: 145},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input)))
	assert.Equal(t, tests[1].expected, part_2([]byte(tests[0].input)))

}
