package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	numArr := make([][]int, 0)

	for first := 0; first < n-3 && nums[first]+nums[first+1]+nums[first+2]+nums[first+3] <= target; first++ {

		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1; second < n-2 && nums[first]+nums[second]+nums[second+1]+nums[second+2] <= target; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			four := n - 1
			for third := second + 1; third < n-1; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}

				for third < four && nums[first]+nums[second]+nums[third]+nums[four] > target {
					four--
				}

				if third == four {
					break
				}

				if four < n-1 && nums[four] == nums[four+1] {
					continue
				}

				if nums[first]+nums[second]+nums[third]+nums[four] == target {
					numArr = append(numArr, []int{nums[first], nums[second], nums[third], nums[four]})
				}
			}
		}
	}

	return numArr
}
