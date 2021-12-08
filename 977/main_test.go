package main

import "testing"

type Case struct {
	nums   []int
	expect []int
}

var caseList = []Case{
	{[]int{-4,-1,0,3,10},[]int{0,1,9,16,100}},
	{[]int{-7,-3,2,3,11},[]int{4,9,9,49,121}},		
}


func TestSortedSquares(t *testing.T) {
	for _, item := range caseList {
		result := sortedSquares(item.nums)
		for i, v := range result {
			if v != item.expect[i] {
				t.Errorf("expect != result: %d != %d", item.expect, result)				
			}
		}
	}
}

func TestSortedSquares2(t *testing.T) {
	for _, item := range caseList {
		result := sortedSquares2(item.nums)
		for i, v := range result {
			if v != item.expect[i] {
				t.Errorf("expect != result: %d != %d", item.expect,result)				
			}
		}
	}
}
