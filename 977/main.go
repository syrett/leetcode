package main

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	rNums := make([]int, len(nums))
	for i,v := range nums {
		rNums[i] = v*v
	}

	sort.Ints(rNums)
	return rNums
}

// 双指针
func sortedSquares2(nums []int) []int {
	rNums := make([]int, len(nums))
	i,j := 0, len(nums) -1
	for pos:=j;pos>=0;pos--{
		if nums[i]*nums[i] > nums[j]*nums[j] {
			rNums[pos] = nums[i]*nums[i]
			i++
		}else{
			rNums[pos] = nums[j]*nums[j]
			j--			
		}
	}

	return rNums
}
