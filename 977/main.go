package main

func sortedSquares(nums []int) []int {
	if len(nums) < 1 {
		return []int{}
	}
	result := []int{nums[0] * nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			result = append([]int{nums[i] * nums[i]}, result...)
		} else {
			if nums[i]*nums[i] >= result[i-1] {
				result = append(result, nums[i]*nums[i])
			} else if nums[i]*nums[i] <= result[0] {
				result = append([]int{nums[i] * nums[i]}, result...)
			} else {
				lenResult := len(result)
				for j := 0; j < len(result); j++ {
					if result[j] >= nums[i]*nums[i] {
						result = append(result[0:j], nums[i]*nums[i])
						result = append(result, result[j+1:lenResult]...)
						continue
					}
				}
			}
		}
	}
	return result
}
