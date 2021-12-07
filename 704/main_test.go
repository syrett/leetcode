package main

import "testing"

func TestSearch(t *testing.T) {
	type Case struct {
		input  []int
		target int
		result int
	}

	caseList := []Case{
		Case{[]int{-1, 0, 5, 3, 9, 12}, 9, 4},
		Case{[]int{}, 3, -1},
	}

	for _, c := range caseList {
		result := search(c.input, c.target)
		if result != c.result {
			t.Errorf("FAIL:%v, target:%d  %d!=%d", c.input, c.target, c.result, result)
		}
	}
}
