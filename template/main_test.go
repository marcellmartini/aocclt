package day01_test

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			"",
			[]string{},
			0,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := 1
			if got != tC.expect {
				t.Errorf("part1(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}
