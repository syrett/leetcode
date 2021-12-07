package main

func search(nums []int, target int) int {
	numLen := 0
	cursor := 0
	tmpNums := nums
	middle := 0

	for {
		numLen = len(tmpNums)
		if numLen == 0 {
			return -1
		}
		if numLen == 1 {
			if tmpNums[0] == target {
				return cursor
			} else {
				return -1
			}
		}
		middle = numLen / 2
		if target < tmpNums[middle] {
			tmpNums = tmpNums[0:middle]
		} else {
			tmpNums = tmpNums[middle:numLen]
			cursor = cursor + middle
		}
	}
	return -1
}
