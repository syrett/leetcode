package main

import (
	"sort"
)

func main() {

}

func smallerNumbersThanCurrent(nums []int) []int {
	originNums := make([]int, len(nums))
	copy(originNums, nums)

	countMap := make(map[int]int, 0)
	sort.Ints(nums)
	for i, v := range nums {
		if _, ok := countMap[v]; !ok {
			countMap[v] = i
		}

	}

	resultNums := make([]int, len(nums))
	for _, v := range originNums {
	}
	return nums
}
