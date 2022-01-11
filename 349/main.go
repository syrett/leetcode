package main

func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int, 0)
	resSlice := make([]int, 0)

	for _, v := range nums1 {
		numMap[v] = 1
	}

	for _, v := range nums2 {
		if _, ok := numMap[v]; ok {
			numMap[v] = numMap[v] + 1
			if numMap[v] == 2 {
				resSlice = append(resSlice, v)
			}
		}
	}

	return resSlice
}
