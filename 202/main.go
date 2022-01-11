package main

func isHappy(n int) bool {
	var sumMap = make(map[int]bool, 0)
	sum := calHappy(n)
	for ; sum != 1 && !sumMap[sum]; sum = calHappy(sum) {
		sumMap[sum] = true
	}

	if sum == 1 {
		return true
	}
	return false

}

func calHappy(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
