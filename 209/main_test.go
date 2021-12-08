package main

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	type Case struct {
		nums   []int
		target int
		expect int
	}

	var caseList = []Case{
		{[]int{2,3,1,2,4,3}, 7, 2},
		{[]int{1,4,4}, 4, 1},
		{[]int{1,1,1,1,1,1,1,1}, 11, 0},				
	}

	for _, item := range caseList {
		result := minSubArrayLen(item.target, item.nums)
		if item.expect != result {
			t.Errorf("expect != result: %d != %d", item.expect, result)				
		}
	}	
}
