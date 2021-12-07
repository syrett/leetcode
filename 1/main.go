package main

// 28ms, 3.4M
func twoSum(nums []int, target int) []int {
	numLen := len(nums)
	pointList := make([]int, 0)
	for i, v := range nums {
		if i >= numLen-1 {
			break
		}
		for j, v2 := range nums[i:] {
			if v+v2 == target {
				pointList = append(pointList, i)
				pointList = append(pointList, j+i)
			}
		}
	}
	return pointList
}

// 8ms,6.2M
func twoSumHash(nums []int, target int) []int {
	numMap := make(map[int][]int, 0)
	for i, v := range nums {
		if list, ok := numMap[v]; ok {
			list = append(list, i)
			numMap[v] = list
		} else {
			numMap[v] = []int{i}
		}
	}
	for value, orginIndex := range numMap {
		expectValue := target - value
		if anotherIndex, ok := numMap[expectValue]; ok {
			if value == expectValue {
				if len(orginIndex) > 1 {
					return []int{orginIndex[0], orginIndex[1]}
				}
			} else {
				return []int{orginIndex[0], anotherIndex[0]}
			}
		}
	}
	return nil
}

func twoSumHashOne(nums []int, target int) []int {
	numMap := make(map[int]int, 0)
	for index, value := range nums {
		if expectIndex, ok := numMap[target-value]; ok {
			return []int{index, expectIndex}
		}
		numMap[value] = index
	}
	return nil
}
