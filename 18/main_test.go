package main

import "testing"

func TestFourSum(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
	}

	caseList := []testCase{
		{[]int{2, 2, 2, 2, 2}, 8},
	}

	for _, item := range caseList {
		result := fourSum(item.nums, item.target)
		t.Logf("result:%+v\n", result)
	}
}
