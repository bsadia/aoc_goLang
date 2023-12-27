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
		{expected: 54, input: `jqt: rhn xhk nvd
		rsh: frs pzl lsr
		xhk: hfx
		cmg: qnr nvd lhk bvb
		rhn: xhk bvb hfx
		bvb: xhk hfx
		pzl: lsr hfx nvd
		qnr: nvd
		ntq: jqt hfx bvb xhk
		nvd: lhk
		lsr: lhk
		rzs: qnr cmg lsr rsh
		frs: qnr lhk lsr`},
	}

	assert.Equal(t, tests[0].expected, part_1([]byte(tests[0].input)))

}
