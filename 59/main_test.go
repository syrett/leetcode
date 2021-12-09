package main

import "testing"

func TestGenerateMatrix(t *testing.T) {
	type Case struct {
		n int
		expect [][]int
	}

	var caseList = []Case{
		{3,[][]int{{1,2,3},{8,9,4},{7,6,5}}},
		{1,[][]int{{1}}},
		{4,[][]int{{1,2,3,4},{12,13,14,5},{11,16,15,6},{10,9,8,7}}},				
	}

	for _, item := range caseList {
		result := generateMatrix(item.n)
		for i:=0;i<item.n;i++{
			for j:=0;j<item.n;j++{
				if item.expect[i][j] != result[i][j]{
					t.Errorf("expect != result: %v != %v", item.expect, result)
				}
			}
		}
	}	
}

