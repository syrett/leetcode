package main

import "testing"

func TestReverse(t *testing.T) {
	type Case struct {
		origin int
		result int
	}

	caseList := []Case{
		Case{123, 321},
		Case{0, 0},
		Case{-123, -321},
		Case{1534236469, 0},
	}

	for _, c := range caseList {
		result := reverse(c.origin)
		if result != c.result {
			t.Errorf("FAIL: origin:%d, %d!=%d", c.origin, result, c.result)
		}
	}
}
