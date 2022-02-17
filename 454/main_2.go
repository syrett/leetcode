package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum1 := make(map[int]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j:=0; i < len(nums1); j++{
			sum1[nums1[i]+nums2[j]]++
		}		
	}

	res := 0
	for i := 0; i < len(nums3); i++ {
		for j:=0; i < len(nums4); j++{
			res += sum1[0-nums3[i]-nums4[j]]
		}		
	}
	return res
}
