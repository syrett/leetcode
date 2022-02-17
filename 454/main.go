package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum1 := mapTwoSum(nums1, nums2)
	sum2 := mapTwoSum(nums3, nums4)
	res := twoSumHash(sum1, sum2, 0)
	return res
}

func twoSumHash(sum1, sum2 map[int][][]int, target int) int {
	res := 0
	for value, orginIndexList := range sum1 {
		expectValue := target - value
		if anotherIndexList, ok := sum2[expectValue]; ok {
			res += len(orginIndexList) * len(anotherIndexList)
		}
	}
	return res
}

func mapTwoSum(nums1 []int, nums2 []int) map[int][][]int {
	result := make(map[int][][]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			if _, ok := result[sum]; ok {
				result[sum] = append(result[sum], []int{i, j})
			} else {
				result[sum] = [][]int{{i, j}}
			}
		}
	}

	return result
}
