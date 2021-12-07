package main

import "testing"

func TestRemoveElement(t *testing.T) {
	type Case struct {
		nums   []int
		val    int
		expect int
	}

	caseList := []Case{
		{[]int{2, 3, 4, 2}, 2, 2},
		{[]int{2, 3, 4, 2}, 1, 4},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, item := range caseList {
		result := removeElement(item.nums, item.val)
		if result != item.expect {
			t.Errorf("expect != result: %d != %d", result, item.expect)
		}
	}
}
