package main

import "testing"

func TestIsHappy(t *testing.T) {
	type testCase struct {
		input   int
		isHappy bool
	}

	caseList := []testCase{
		{1, true},
		{4, false},
	}

	for _, item := range caseList {
		isHappyR := isHappy(item.input)
		if isHappyR != item.isHappy {
			t.Errorf("input:%d, output(%v) != expect(%v)", item.input, isHappyR, item.isHappy)
		}
	}
}
