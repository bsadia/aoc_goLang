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
		{expected: 5, input: `1,0,1~1,2,1
		0,0,2~2,0,2
		0,2,3~2,2,3
		0,0,4~0,2,4
		2,0,5~2,2,5
		0,1,6~2,1,6
		1,1,8~1,1,9`},
		{expected: 7},
	}

	assert.Equal(t, tests[0].expected, part_1(([]byte(tests[0].input)))[0])
	assert.Equal(t, tests[1].expected, part_1(([]byte(tests[0].input)))[1])

}
