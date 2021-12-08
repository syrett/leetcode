package main

import ()

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n<2{
		if nums[0] >= target {
			return 1
		}else{
			return 0
		}
	}
	

	minStep := n+1
	sum := 0
	for start,end := 0,0; end<n;{
		sum += nums[end]
		for sum >= target {
			if end-start+1 < minStep {
				minStep = end-start+1
			}
			sum -= nums[start]
			start++
		}
		end++
	}

	if minStep > n {
		return 0
	}else{
		return minStep
	}
	
}
